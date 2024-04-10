package category

import (
	"context"
	"github.com/opentracing/opentracing-go"
	model "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/model/category"
	apiV1 "github.com/robowealth-mutual-fund/blueprint-roa-golang/pkg/api/v1"
)

func (c *Controller) Create(ctx context.Context, request *apiV1.CategoryCreateRequest) (*apiV1.CategoryCreateResponse, error) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		ctx,
		opentracing.GlobalTracer(),
		"handler.category.Create",
	)
	defer span.Finish()
	span.LogKV("Input Handler", request)
	id, err := c.service.Create(ctx, &model.Request{
		Name:   request.GetName(),
		Detail: request.GetDetail(),
	})

	if err != nil {
		return nil, err
	}
	return &apiV1.CategoryCreateResponse{Id: int32(id)}, nil
}
