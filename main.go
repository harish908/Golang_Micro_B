package main

import (
	"context"
	"flag"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	grpcHealth "github.com/harish908/Golang_Micro_B/proto/gen/health_check"
	"github.com/harish908/Golang_Micro_B/service"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:9090", "gRPC server endpoint")
)

func main() {
	// Add 1 because if one server gets failed we need to exit
	// wg := &sync.WaitGroup{}
	// wg.Add(1)
	// go runGRPCServer(wg)
	// go runHTTPServer(wg)
	// wg.Wait()
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go runGRPCServer(wg)
	go runGrpcHttpServer(wg)
	wg.Wait()
}

func runHTTPServer(wg *sync.WaitGroup) {
	e := echo.New()
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "Service B up")
	})

	e.Logger.Fatal(e.Start("0.0.0.0:80"))
	wg.Done()
}

func runGRPCServer(wg *sync.WaitGroup) {
	lis, err := net.Listen("tcp", *grpcServerEndpoint)
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}
	s := grpc.NewServer()
	grpcHealth.RegisterHealthCheckServiceServer(s, &service.HealthCheckServer{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}
	wg.Done()
}

func runGrpcHttpServer(wg *sync.WaitGroup) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := grpcHealth.RegisterHealthCheckServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		log.Fatalf("Grpc service error: %v", err)
	}

	err = http.ListenAndServe("0.0.0.0:8000", mux)
	if err != nil {
		log.Fatalf("Error while starting http server : %v", err)
	}
	wg.Done()
}
