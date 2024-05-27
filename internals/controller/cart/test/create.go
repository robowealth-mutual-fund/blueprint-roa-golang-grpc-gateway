package test

import (
	cartModel "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/model/cart"
	api_v1 "github.com/robowealth-mutual-fund/blueprint-roa-golang/pkg/api/v1"
	"github.com/stretchr/testify/mock"
)

func (suite *PackageTestSuite) TestCreate() {

	name := "aaa"
	mockCart := &cartModel.Request{
		Name:   "aaa",
		Detail: "aaa",
		Brand:  "aaa",
		Price:  "aaaa",
	}
	suite.cartService.On("Create", mock.Anything, mockCart).Once().Return(name, nil)
	data, err := suite.ctrl.Create(suite.ctx, &api_v1.CreateCartRequest{
		Name:   "aaa",
		Detail: "aaa",
		Brand:  "aaa",
		Price:  "aaaa",
	})
	suite.Equal("aaa", data.Name)
	suite.NoError(err)
}
