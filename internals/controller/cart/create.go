package cart

import (
	"context"

	model "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/model/cart"
	apiV1 "github.com/robowealth-mutual-fund/blueprint-roa-golang/pkg/api/v1"
)

func (c *Controller) Create(ctx context.Context, request *apiV1.CreateCartRequest) (*apiV1.CreateCartResponse, error) {

	name, err := c.service.Create(ctx, &model.Request{
		Name:   request.GetName(),
		Brand:  request.GetBrand(),
		Detail: request.GetDetail(),
		Price:  request.GetPrice(),
	})

	if err != nil {
		return nil, err
	}
	return &apiV1.CreateCartResponse{Name: name}, nil
}
