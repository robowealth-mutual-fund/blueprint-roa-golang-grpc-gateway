package controller

import (
	"context"

	grpc_v1 "github.com/robowealth-mutual-fund/blueprint-roa-golang/pkg/grpc/health/v1"
)

// HealthZController ...
type HealthZController struct{}

// Check ...
func (*HealthZController) Check(ctx context.Context, req *grpc_v1.HealthCheckRequest) (*grpc_v1.HealthCheckResponse, error) {
	return &grpc_v1.HealthCheckResponse{
		Status: grpc_v1.HealthCheckResponse_SERVING,
	}, nil
}

// NewHealthZController ...
func NewHealthZController() *HealthZController {
	return &HealthZController{}
}
