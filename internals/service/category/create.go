package category

import (
	"context"

	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/entity"
	model "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/model/category"
)

func (s *CategoryService) Create(ctx context.Context, request *model.Request) (int, error) {

	input := &entity.Category{
		Name:   request.Name,
		Detail: request.Detail,
	}

	err := s.repository.Create(input)

	//sp.LogKV("Repository result  :", err)

	return input.ID, err
}
