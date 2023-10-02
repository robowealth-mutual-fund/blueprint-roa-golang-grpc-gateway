package wrapper

import (
	"context"
	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/entity"
	model "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/model/product"
)

func (s *ProductService) Create(ctx context.Context, request *model.Request) (int, error) {

	//sp, ctx := opentracing.StartSpanFromContext(ctx, "service.users.Create")
	//defer sp.Finish()
	//
	//sp.LogKV("Service :", request)
	//sp.LogKV("Brand", request.Brand)

	input := &entity.Product{
		Name:   request.Name,
		Detail: request.Detail,
		Brand:  request.Brand,
		Price:  request.Price,
	}

	err := s.repository.Create(input)

	//sp.LogKV("Repository result  :", err)

	return input.ID, err
}
