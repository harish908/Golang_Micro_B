package main

import (
	"log"
	"net"
	"net/http"

	grpcHealth "github.com/harish908/Golang_Micro_B/proto/gen/health_check"
	"github.com/harish908/Golang_Micro_B/service"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

func main() {
	runGRPCServer()
	runHTTPServer()
}

func runHTTPServer() {
	e := echo.New()
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "Service B up")
	})

	e.Logger.Fatal(e.Start(":4001"))
}

func runGRPCServer() {
	lis, err := net.Listen("tcp", "127.0.0.1:5001")
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}
	s := grpc.NewServer()
	grpcHealth.RegisterHealthCheckServiceServer(s, &service.HealthCheckServer{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
