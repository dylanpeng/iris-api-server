package router

import (
	"google.golang.org/grpc"
	health "google.golang.org/grpc/health/grpc_health_v1"
	"juggernaut/gateway/logic/control"
	"juggernaut/lib/proto/juggernaut/service/gateway"
)

func (r *router) RegGrpcService(s *grpc.Server) {
	gateway.RegisterGatewayServiceServer(s, control.Gateway)
	health.RegisterHealthServer(s, control.Grpc)
}
