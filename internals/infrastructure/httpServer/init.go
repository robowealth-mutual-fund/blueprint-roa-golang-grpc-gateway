package httpServer

import (
	"context"
	"net/http"
	"strconv"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/config"
	controllerCategory "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/controller/category"
	controllerProduct "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/controller/product"
	controllerUsers "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/controller/users"
	controllerWarehouse "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/controller/warehouse"
	apiV1 "github.com/robowealth-mutual-fund/blueprint-roa-golang/pkg/api/v1"
	"google.golang.org/grpc"
)

type Server struct {
	Config        config.Configuration
	Server        *runtime.ServeMux
	HttpMux       *http.ServeMux
	ProductCtrl   *controllerProduct.Controller
	CategoryCtrl  *controllerCategory.Controller
	WarehouseCtrl *controllerWarehouse.Controller
	UsersCtrl     *controllerUsers.Controller
}

func (s *Server) Configure(ctx context.Context, opts []grpc.DialOption) {
	apiV1.RegisterProductServiceHandlerFromEndpoint(ctx, s.Server, "0.0.0.0:"+strconv.Itoa(s.Config.Port), opts)
	apiV1.RegisterCategoryServiceHandlerFromEndpoint(ctx, s.Server, "0.0.0.0:"+strconv.Itoa(s.Config.Port), opts)
	apiV1.RegisterWarehouseServiceHandlerFromEndpoint(ctx, s.Server, "0.0.0.0:"+strconv.Itoa(s.Config.Port), opts)
	apiV1.RegisterUsersServiceHandlerFromEndpoint(ctx, s.Server, "0.0.0.0:"+strconv.Itoa(s.Config.Port), opts)

}

func NewServer(config config.Configuration, rmux *runtime.ServeMux, httpMux *http.ServeMux,
	productCtrl *controllerProduct.Controller,
	categoryCtrl *controllerCategory.Controller,
	warehouseCtrl *controllerWarehouse.Controller,
	usersCtrl *controllerUsers.Controller,
) *Server {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	s := &Server{
		Config:        config,
		Server:        rmux,
		HttpMux:       httpMux,
		ProductCtrl:   productCtrl,
		CategoryCtrl:  categoryCtrl,
		WarehouseCtrl: warehouseCtrl,
		UsersCtrl:     usersCtrl,
	}
	s.Configure(context.Background(), opts)
	return s
}
