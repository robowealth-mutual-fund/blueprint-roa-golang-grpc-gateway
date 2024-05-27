package cartLiveing

import (
	"context"

	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/entity"
	model "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/model/cart"
)

func (s *CartService) Create(ctx context.Context, request *model.Request) (string, error) {
	input := &entity.Cart{
		Name:   request.Name,
		Detail: request.Detail,
		Brand:  request.Brand,
		Price:  request.Price,
	}

	err := s.repository.Create(input)

	return input.Name, err
}
