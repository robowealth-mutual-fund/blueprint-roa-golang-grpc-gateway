package grpcserver

import (
	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/config"
	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/controller"
	controllerProduct "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/controller/product"
	apiV1 "github.com/robowealth-mutual-fund/blueprint-roa-golang/pkg/api/v1"
	grpcHealthV1 "github.com/robowealth-mutual-fund/blueprint-roa-golang/pkg/grpc/health/v1"
	"google.golang.org/grpc"
)

type Server struct {
	Config       config.Configuration
	Server       *grpc.Server
	HealthCtrl   *controller.HealthZController
	PingPongCtrl *controller.PingPongController
	ProductCtrl  *controllerProduct.Controller
}

// Configure ...
func (s *Server) Configure() {
	grpcHealthV1.RegisterHealthServer(s.Server, s.HealthCtrl)
	apiV1.RegisterPingPongServiceServer(s.Server, s.PingPongCtrl)
	apiV1.RegisterProductServiceServer(s.Server, s.ProductCtrl)
}

func NewServer(
	config config.Configuration,
	healthCtrl *controller.HealthZController,
	pingPongCtrl *controller.PingPongController,
	productCtrl *controllerProduct.Controller,
// v *validator.CustomValidator,
) *Server {
	option := grpc.ChainUnaryInterceptor(
		// grpc_errors.UnaryServerInterceptor(),
		// validator.UnaryServerInterceptor(v),
	)

	s := &Server{
		Server:       grpc.NewServer(option, grpc.MaxRecvMsgSize(10*10e6), grpc.MaxSendMsgSize(10*10e6)),
		Config:       config,
		HealthCtrl:   healthCtrl,
		PingPongCtrl: pingPongCtrl,
		ProductCtrl:  productCtrl,
	}

	s.Configure()

	return s
}
