package product

import (
	"context"
	"github.com/opentracing/opentracing-go"
	model "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/model/product"
	apiV1 "github.com/robowealth-mutual-fund/blueprint-roa-golang/pkg/api/v1"
)

func (c *Controller) Create(ctx context.Context, request *apiV1.CreateRequest) (*apiV1.CreateResponse, error) {
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
		span.LogKV("Handler ERROR :", err)
		return nil, err
	}
	return &apiV1.CreateResponse{Id: int32(id)}, nil
}
