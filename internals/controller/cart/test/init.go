package test

import (
	"context"
	"net/http"

	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/service/cart/wrapper"
	"go.uber.org/dig"

	"github.com/stretchr/testify/suite"

	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/config"
	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/controller/cart"
	mockCart "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/service/cart/mocks"
)

type PackageTestSuite struct {
	suite.Suite
	router      *http.Server
	ctx         context.Context
	conf        config.Configuration
	ctrl        *cart.Controller
	wrapper     wrapper.Wrapper
	cartService *mockCart.Service
}

func (suite *PackageTestSuite) SetupSuite() {
	suite.conf = config.NewConfiguration()
	suite.ctx = context.Background()
	suite.cartService = &mockCart.Service{}
	suite.wrapper = wrapper.Wrapper{
		In:      dig.In{},
		Service: suite.cartService,
	}
	suite.ctrl = cart.NewController(suite.wrapper)
}
