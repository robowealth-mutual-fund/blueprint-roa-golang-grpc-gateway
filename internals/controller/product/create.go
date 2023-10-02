package product

import (
	"context"
	"github.com/opentracing/opentracing-go"
	model "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/model/product"
	api_v1 "github.com/robowealth-mutual-fund/blueprint-roa-golang/pkg/api/v1"
)

func (c *Controller) Create(ctx context.Context, request *api_v1.CreateRequest) (*api_v1.CreateResponse, error) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		ctx,
		opentracing.GlobalTracer(),
		"handler.product.Create",
	)
	defer span.Finish()
	span.LogKV("Input Handler", request)
	id, err := c.service.Create(ctx, &model.Request{
		Name:   request.GetName(),
		Brand:  request.GetBrand(),
		Detail: request.GetDetail(),
		Price:  request.GetPrice(),
	})

	if err != nil {
		return nil, err
	}

	return &api_v1.CreateResponse{Id: int32(id)}, nil
}
