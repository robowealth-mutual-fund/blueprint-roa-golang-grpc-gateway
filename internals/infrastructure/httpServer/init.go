package httpServer

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/config"
	controllerProduct "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/controller/product"
	apiV1 "github.com/robowealth-mutual-fund/blueprint-roa-golang/pkg/api/v1"
	"google.golang.org/grpc"
	"net/http"
	"strconv"
)

type Server struct {
	Config      config.Configuration
	Server      *runtime.ServeMux
	HttpMux     *http.ServeMux
	ProductCtrl *controllerProduct.Controller
}

func (s *Server) Configure(ctx context.Context, opts []grpc.DialOption) {
	apiV1.RegisterProductServiceHandlerFromEndpoint(ctx, s.Server, "0.0.0.0:"+strconv.Itoa(s.Config.Port), opts)
}

func NewServer(config config.Configuration, rmux *runtime.ServeMux, httpMux *http.ServeMux, productCtrl *controllerProduct.Controller) *Server {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	s := &Server{
		Config:      config,
		Server:      rmux,
		HttpMux:     httpMux,
		ProductCtrl: productCtrl,
	}
	s.Configure(context.Background(), opts)
	return s
}
