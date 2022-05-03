package service

import (
	"context"

	grpcHealth "github.com/harish908/Golang_Micro_B/proto/gen/health_check"
	"google.golang.org/protobuf/types/known/emptypb"
)

type HealthCheckServer struct{}

func (c *HealthCheckServer) CheckGRPCConnection(ctx context.Context, empty *emptypb.Empty) (*grpcHealth.HealthCheckResponse, error) {
	response := &grpcHealth.HealthCheckResponse{
		Connection: "Service B Up and running",
	}
	return response, nil
}
