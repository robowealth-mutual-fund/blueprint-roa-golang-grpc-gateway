package test

import (
	"context"

	"github.com/stretchr/testify/suite"

	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/entity"
	cartModel "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/model/cart"
	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/repository/postgres/mocks"
	service "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/service/cart"
)

type PackageTestSuite struct {
	suite.Suite
	ctx     context.Context
	repo    *mocks.Repository
	service service.Service
}

func (suite *PackageTestSuite) SetupTest() {
	suite.ctx = context.Background()
	suite.repo = &mocks.Repository{}
	suite.service = service.NewService(suite.repo)
}

func (suite *PackageTestSuite) makeTestCart() (verbose *entity.Cart) {
	return &entity.Cart{
		Name:   "aaa",
		Detail: "aaa",
		Brand:  "bbbb",
		Price:  "123132",
	}
}

func (suite *PackageTestSuite) makeTestCreateInput() (input *cartModel.Request) {
	return &cartModel.Request{
		Name:   "aaa",
		Detail: "aaa",
		Brand:  "bbbb",
		Price:  "123132",
	}
}
