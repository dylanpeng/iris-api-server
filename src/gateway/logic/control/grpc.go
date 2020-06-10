package control

import (
	"context"
	"google.golang.org/grpc/codes"
	health "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
)

var Grpc = &grpcCtrl{}

type grpcCtrl struct{}

func (c *grpcCtrl) Check(ctx context.Context, req *health.HealthCheckRequest) (*health.HealthCheckResponse, error) {
	return &health.HealthCheckResponse{
		Status: health.HealthCheckResponse_SERVING,
	}, nil
}

func (c *grpcCtrl) Watch(in *health.HealthCheckRequest, stream health.Health_WatchServer) error {
	for {
		if err := stream.Send(&health.HealthCheckResponse{Status: health.HealthCheckResponse_SERVING}); err != nil {
			return status.Error(codes.Canceled, err.Error())
		}
	}
}
