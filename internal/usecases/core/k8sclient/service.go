package k8sclient

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

type GetServicesParams struct {
	Namespace     string
	Repository    string
	PullRequestID string
	Branch        string
}

func (k *k8sClient) GetServices(ctx context.Context, param GetServicesParams, opts ...Option) ([]*corev1.Service, error) {
	o := newOption()
	for _, opt := range opts {
		opt(o)
	}

	services, err := k.client.CoreV1().Services(param.Namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var result []*corev1.Service

	for _, service := range services.Items {
		// laebls match
		if service.Labels[RepositoryLabel] == param.Repository {
			if len(param.PullRequestID) != 0 {
				if service.Labels[PullRequestID] == param.PullRequestID {
					result = append(result, &service)
				}
			}

			if len(param.Branch) != 0 {
				if service.Labels[BaseBranchLabel] == param.Branch {
					result = append(result, &service)
				}
			}
		}
	}

	return result, nil
}

type CreateServiceNodePortParams struct {
	Namespace   string
	ServiceName string
	TargetPort  []int32
}

func (k *k8sClient) CreateNodePortService(ctx context.Context, param CreateServiceNodePortParams, opts ...Option) (*corev1.Service, error) {
	o := newOption()
	for _, opt := range opts {
		opt(o)
	}

	var expose []corev1.ServicePort
	for _, port := range param.TargetPort {
		if port == 0 {
			continue
		}
		// rand, _ := random.RandomString(10)
		expose = append(expose, corev1.ServicePort{
			Port: port,
		})

	}

	return k.client.CoreV1().Services(param.Namespace).Create(ctx, &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: fmt.Sprintf("%s-", param.ServiceName),
			Labels:       o.labelSelector,
		},
		Spec: corev1.ServiceSpec{
			Ports:    expose,
			Selector: o.labelSelector,
			Type:     corev1.ServiceTypeNodePort,
		},
	}, metav1.CreateOptions{})
}

func (k *k8sClient) DeleteService(ctx context.Context, namespace string, opts ...Option) error {
	o := newOption()

	for _, opt := range opts {
		opt(o)
	}

	services, err := k.client.CoreV1().Services(namespace).List(ctx, metav1.ListOptions{
		LabelSelector: labels.Set(o.labelSelector).String(),
	})

	if err != nil {
		return err
	}

	for _, service := range services.Items {
		if err := k.client.CoreV1().Services(namespace).Delete(ctx, service.Name, metav1.DeleteOptions{}); err != nil {
			return err
		}
	}

	return nil
}
