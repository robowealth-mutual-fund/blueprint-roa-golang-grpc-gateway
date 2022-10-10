package users

import (
	"context"

	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/entity"
	model "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/model/users"
)

func (s *UsersService) Create(ctx context.Context, request *model.Request) (int, error) {

	input := &entity.Users{
		FullName:    request.FirstName + " " + request.LastName,
		Address:     request.Address,
		PhoneNumber: request.PhoneNumber,
		Gender:      request.Gender,
	}

	err := s.repository.Create(input)

	//sp.LogKV("Repository result  :", err)

	return input.ID, err
}
