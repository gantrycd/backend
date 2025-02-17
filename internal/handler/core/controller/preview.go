package controller

import (
	"context"
	"errors"
	"fmt"
	"regexp"

	"github.com/aura-cd/backend/cmd/config"
	coreErr "github.com/aura-cd/backend/internal/error"
	"github.com/aura-cd/backend/internal/usecases/core/k8sclient"
	"github.com/aura-cd/backend/internal/utils"
	"github.com/aura-cd/backend/internal/utils/branch"
	v1 "github.com/aura-cd/backend/proto"
	"github.com/cloudflare/cloudflare-go"
	"google.golang.org/protobuf/types/known/emptypb"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *controller) CreatePreview(ctx context.Context, in *v1.CreatePreviewRequest) (*v1.CreatePreviewReply, error) {
	branchName := branch.Transpile1123(in.Branch)
	dep, err := c.control.GetDeployment(ctx,
		k8sclient.GetDeploymentParams{
			Namespace:     in.Organization,
			Repository:    in.Repository,
			PullRequestID: in.PullRequestId,
			Branch:        branchName,
		})
	if err != nil && !errors.Is(err, coreErr.ErrDeploymentsNotFound) {
		return nil, err
	}

	if dep != nil {
		return &v1.CreatePreviewReply{
			Name:      dep.Name,
			Namespace: dep.Namespace,
			Version:   dep.ResourceVersion,
		}, nil
	}

	return c.createDeployment(ctx, in)
}

func (c *controller) createDeployment(ctx context.Context, in *v1.CreatePreviewRequest) (*v1.CreatePreviewReply, error) {
	branchName := branch.Transpile1123(in.Branch)

	if in.Replicas == 0 {
		in.Replicas = 1
	}

	deps, err := c.control.CreateDeployment(ctx,
		k8sclient.CreateDeploymentParams{
			Namespace: in.Organization,
			AppName:   in.Repository,
			Image:     in.Image,
			Config:    in.Configs,
			Replicas:  in.Replicas,
		},
		k8sclient.WithRepositoryLabel(in.Repository),
		k8sclient.WithPrIDLabel(in.PullRequestId),
		k8sclient.WithEnvirionmentLabel(k8sclient.EnvPreview),
		k8sclient.WithBaseBranchLabel(branchName),
		k8sclient.WithCreatedByLabel(k8sclient.AppIdentifier),
	)
	if err != nil {
		return nil, err
	}

	// NodePortの指定がなかったら終了
	if in.ExposePorts == nil {
		return &v1.CreatePreviewReply{
			Name:      deps.Name,
			Namespace: deps.Namespace,
			Version:   deps.ResourceVersion,
		}, nil
	}

	service, err := c.control.CreateLoadBalancerService(ctx,
		k8sclient.CreateServiceNodePortParams{
			Namespace:   in.Organization,
			ServiceName: deps.Name,
			TargetPort:  in.ExposePorts,
		},
		k8sclient.WithAppLabel(in.Repository),
		k8sclient.WithRepositoryLabel(in.Repository),
		k8sclient.WithPrIDLabel(in.PullRequestId),
		k8sclient.WithEnvirionmentLabel(k8sclient.EnvPreview),
		k8sclient.WithBaseBranchLabel(branchName),
		k8sclient.WithCreatedByLabel(k8sclient.AppIdentifier),
	)
	if err != nil {
		return nil, err
	}

	domains, err := createCloudflared(c, ctx, in, service.Name)
	if err != nil {
		return nil, err
	}

	return &v1.CreatePreviewReply{
		Name:      deps.Name,
		Namespace: deps.Namespace,
		Version:   deps.ResourceVersion,
		External:  domains,
	}, nil
}

func createCloudflared(c *controller, ctx context.Context, in *v1.CreatePreviewRequest, serviceName string) ([]string, error) {
	baseDomain := fmt.Sprintf("%s-%s-%s", in.Organization, in.Repository, in.PullRequestId)
	cloudflaredConfigYaml, domains, subDomains := buildCloudflaredConfig(in.Organization, serviceName, baseDomain, in.ExposePorts)

	configMapName := fmt.Sprintf("%s-config-map", baseDomain)
	cloudflaredPodName := fmt.Sprintf("%s-cloudflared", baseDomain)

	if err := c.control.CreateConfigMap(ctx, corev1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      configMapName,
			Namespace: in.Organization,
		},
		Data: map[string]string{
			"config.yaml": cloudflaredConfigYaml,
		},
	}, metav1.CreateOptions{}); err != nil {
		return nil, err
	}

	if _, err := c.control.CreatePod(ctx, &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cloudflaredPodName,
			Namespace: in.Organization,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  cloudflaredPodName,
					Image: "cloudflare/cloudflared:latest",
					VolumeMounts: []corev1.VolumeMount{
						{
							Name:      "config",
							MountPath: "/etc/cloudflared/config/",
						},
						{
							Name:      "credentials",
							MountPath: "/etc/cloudflared/credentials/",
						},
					},
					Args: []string{"tunnel", "--config", "/etc/cloudflared/config/config.yaml", "run"},
				},
			},
			Volumes: []corev1.Volume{
				{
					Name: "config",
					VolumeSource: corev1.VolumeSource{
						ConfigMap: &corev1.ConfigMapVolumeSource{
							LocalObjectReference: corev1.LocalObjectReference{
								Name: configMapName,
							},
						},
					},
				},
				{
					Name: "credentials",
					VolumeSource: corev1.VolumeSource{
						Secret: &corev1.SecretVolumeSource{
							SecretName: "cloudflared-credentials",
						},
					},
				},
			},
		},
	}, metav1.CreateOptions{}); err != nil {
		return nil, err
	}
	if err := createCloudflareDNSRecord(ctx, subDomains); err != nil {
		return nil, err
	}
	return domains, nil
}

func buildCloudflaredConfig(namespace string, serviceName string, baseDomain string, ports []int32) (string, []string, []string) {
	var ingress = ""
	var domains []string
	var subDomains []string
	for _, port := range ports {
		subDomain := fmt.Sprintf("%s-%d", baseDomain, port)
		domain := fmt.Sprintf("%s.%s", subDomain, config.Config.Application.ExternalDomain)
		domains = append(domains, domain)
		subDomains = append(subDomains, subDomain)
		ingress += fmt.Sprintf(`  - hostname: %s
    service: http://%s.%s.svc.cluster.local:%d
`, domain, serviceName, namespace, port)
	}
	ingress += `  - service: http_status:404
`
	return fmt.Sprintf(`tunnel: %s
credentials-file: /etc/cloudflared/credentials/credentials.json
no-autoupdate: true

ingress:
%s`, config.Config.Application.CloudflaredTunnelId, ingress), domains, subDomains
}

func createCloudflareDNSRecord(ctx context.Context, subDomains []string) error {
	api, err := cloudflare.NewWithAPIToken(config.Config.Application.CloudflareAPIToken)
	if err != nil {
		return err
	}
	zoneId, err := getCloudflareZoneId(ctx, api, config.Config.Application.ExternalDomain)
	zone := cloudflare.ZoneIdentifier(zoneId)
	records, _, err := api.ListDNSRecords(ctx, zone, cloudflare.ListDNSRecordsParams{})
	if err != nil {
		return err
	}
	cnameDomain := fmt.Sprintf("%s.cfargotunnel.com", config.Config.Application.CloudflaredTunnelId)
	for _, subDomain := range subDomains {
		record := getRecordByDomain(records, subDomain)
		if record == nil {
			_, err := api.CreateDNSRecord(ctx, zone, cloudflare.CreateDNSRecordParams{
				Type:    "CNAME",
				Name:    subDomain,
				Content: cnameDomain,
				Proxied: cloudflare.BoolPtr(true),
			})
			if err != nil {
				return err
			}
		} else {
			if record.Content != cnameDomain {
				record.Content = cnameDomain
				_, err := api.UpdateDNSRecord(ctx, zone, cloudflare.UpdateDNSRecordParams{
					Type:    "CNAME",
					Name:    subDomain,
					Content: cnameDomain,
					Proxied: cloudflare.BoolPtr(true),
				})
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func deleteCloudflareDNSRecord(ctx context.Context, baseDomain string) error {
	api, err := cloudflare.NewWithAPIToken(config.Config.Application.CloudflareAPIToken)
	if err != nil {
		return err
	}
	zoneId, err := getCloudflareZoneId(ctx, api, config.Config.Application.ExternalDomain)
	zone := cloudflare.ZoneIdentifier(zoneId)
	records, _, err := api.ListDNSRecords(ctx, zone, cloudflare.ListDNSRecordsParams{})
	if err != nil {
		return err
	}
	re := regexp.MustCompile(fmt.Sprintf(`^%s-\d+\.%s$`, baseDomain, config.Config.Application.ExternalDomain))
	for _, record := range records {
		if re.MatchString(record.Name) {
			err := api.DeleteDNSRecord(ctx, zone, record.ID)
			if err != nil {
				return err
			}
		}
	}
	return nil

}

func getRecordByDomain(records []cloudflare.DNSRecord, domain string) *cloudflare.DNSRecord {
	for _, record := range records {
		if record.Name == domain {
			return &record
		}
	}
	return nil
}

func getCloudflareZoneId(ctx context.Context, api *cloudflare.API, domain string) (string, error) {
	zones, err := api.ListZones(ctx)
	if err != nil {
		return "", err
	}
	for _, zone := range zones {
		if zone.Name == domain {
			return zone.ID, nil
		}
	}
	return "", fmt.Errorf("zone not found: %s", domain)
}

func (c *controller) UpdatePreview(ctx context.Context, in *v1.CreatePreviewRequest) (*v1.CreatePreviewReply, error) {
	branchName := branch.Transpile1123(in.Branch)

	dep, err := c.control.GetDeployment(ctx,
		k8sclient.GetDeploymentParams{
			Namespace:     in.Organization,
			Repository:    in.Repository,
			PullRequestID: in.PullRequestId,
			Branch:        branchName,
		})
	if err != nil && !errors.Is(err, coreErr.ErrDeploymentsNotFound) {
		return nil, err
	}

	if dep == nil {
		return c.createDeployment(ctx, in)
	}

	dep, err = c.control.UpdateDeployment(ctx, dep, k8sclient.UpdateDeploymentParams{
		Namespace:     in.Organization,
		Repository:    in.Repository,
		PullRequestID: in.PullRequestId,
		Branch:        branch.Transpile1123(in.Branch),
		AppName:       in.Repository,
		Image:         utils.ToPtr(in.Image),
	})
	if err != nil {
		return nil, err
	}

	return &v1.CreatePreviewReply{
		Name:      dep.Name,
		Namespace: dep.Namespace,
		Version:   dep.ResourceVersion,
	}, nil
}

func (c *controller) DeletePreview(ctx context.Context, in *v1.DeletePreviewRequest) (*emptypb.Empty, error) {
	branchName := branch.Transpile1123(in.Branch)
	if err := c.control.DeleteDeployment(ctx,
		in.Organization,
		k8sclient.WithAppLabel(in.Repository),
		k8sclient.WithRepositoryLabel(in.Repository),
		k8sclient.WithBaseBranchLabel(branchName),
		k8sclient.WithPrIDLabel(in.PullRequestId),
		k8sclient.WithCreatedByLabel(k8sclient.AppIdentifier),
		k8sclient.WithEnvirionmentLabel(k8sclient.EnvPreview),
	); err != nil {
		return nil, err
	}

	if err := c.control.DeleteService(ctx,
		in.Organization,
		k8sclient.WithAppLabel(in.Repository),
		k8sclient.WithRepositoryLabel(in.Repository),
		k8sclient.WithBaseBranchLabel(branchName),
		k8sclient.WithPrIDLabel(in.PullRequestId),
		k8sclient.WithCreatedByLabel(k8sclient.AppIdentifier),
		k8sclient.WithEnvirionmentLabel(k8sclient.EnvPreview),
	); err != nil {
		return nil, err
	}

	if err := deleteCloudflared(c, ctx, in); err != nil {
		return nil, err
	}

	return new(emptypb.Empty), nil
}

func deleteCloudflared(c *controller, ctx context.Context, in *v1.DeletePreviewRequest) error {
	baseDomain := fmt.Sprintf("%s-%s-%s", in.Organization, in.Repository, in.PullRequestId)
	configMapName := fmt.Sprintf("%s-config-map", baseDomain)
	cloudflaredPodName := fmt.Sprintf("%s-cloudflared", baseDomain)

	if err := c.control.DeleteConfigMap(ctx, in.Organization, configMapName, metav1.DeleteOptions{}); err != nil {
		return err
	}
	if err := c.control.DeletePod(ctx, in.Organization, cloudflaredPodName, metav1.DeleteOptions{}); err != nil {
		return err
	}

	err := deleteCloudflareDNSRecord(ctx, baseDomain)
	if err != nil {
		return err
	}
	return nil
}
