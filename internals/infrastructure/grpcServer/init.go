package grpcserver

import (
	grpcErrors "git.robodev.co/imp/shared-utility/grpc_errors"
	validatorUtils "git.robodev.co/imp/shared-utility/validator"
	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/config"
	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/controller"
	controllerCategory "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/controller/category"
	controllerProduct "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/controller/product"
	controllerWarehouse "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/controller/warehouse"
	apiV1 "github.com/robowealth-mutual-fund/blueprint-roa-golang/pkg/api/v1"
	grpcHealthV1 "github.com/robowealth-mutual-fund/blueprint-roa-golang/pkg/grpc/health/v1"
	"google.golang.org/grpc"
)

type Server struct {
	Config        config.Configuration
	Server        *grpc.Server
	HealthCtrl    *controller.HealthZController
	PingPongCtrl  *controller.PingPongController
	ProductCtrl   *controllerProduct.Controller
	CategoryCtrl  *controllerCategory.Controller
	WarehouseCtrl *controllerWarehouse.Controller
}

// Configure ...
func (s *Server) Configure() {
	grpcHealthV1.RegisterHealthServer(s.Server, s.HealthCtrl)
	apiV1.RegisterPingPongServiceServer(s.Server, s.PingPongCtrl)
	apiV1.RegisterProductServiceServer(s.Server, s.ProductCtrl)
	apiV1.RegisterCategoryServiceServer(s.Server, s.CategoryCtrl)
	apiV1.RegisterWarehouseServiceServer(s.Server, s.WarehouseCtrl)
}

func NewServer(
	config config.Configuration,
	healthCtrl *controller.HealthZController,
	pingPongCtrl *controller.PingPongController,
	productCtrl *controllerProduct.Controller,
	categoryCtrl *controllerCategory.Controller,
	warehouseCtrl *controllerWarehouse.Controller,
	validator *validatorUtils.CustomValidator,
) *Server {
	option := grpc.ChainUnaryInterceptor(
		grpcErrors.UnaryServerInterceptor(),
		validatorUtils.UnaryServerInterceptor(validator),
	)

	s := &Server{
		Server:        grpc.NewServer(option, grpc.MaxRecvMsgSize(10*10e6), grpc.MaxSendMsgSize(10*10e6)),
		Config:        config,
		HealthCtrl:    healthCtrl,
		PingPongCtrl:  pingPongCtrl,
		ProductCtrl:   productCtrl,
		CategoryCtrl:  categoryCtrl,
		WarehouseCtrl: warehouseCtrl,
	}

	s.Configure()

	return s
}
