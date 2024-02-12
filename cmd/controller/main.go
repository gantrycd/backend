package main

import (
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"

	"github.com/gantrycd/backend/internal/driver/k8s"
	"github.com/gantrycd/backend/internal/handler/core/controller"
	v1 "github.com/gantrycd/backend/proto/k8s-controller"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
}

func run() error {
	path := ".kube/config"
	client, err := k8s.New(&path).ConnectTyped("")
	if err != nil {
		return err
	}

	l := slog.New(slog.NewTextHandler(os.Stderr, nil)).WithGroup("server")

	grpcServer := grpc.NewServer()
	defer grpcServer.Stop()

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT")))
	if err != nil {
		return err
	}

	// TODO: Implement the server
	v1.RegisterK8SCustomControllerServer(grpcServer, controller.NewController(client))

	reflection.Register(grpcServer)

	go func() {
		l.Info("server is running at :8080")
		if err := grpcServer.Serve(listener); err != nil {
			l.Error("failed to serve", "error", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit
	l.Info("stopping gRPC server...")
	grpcServer.GracefulStop()

	return nil
}
