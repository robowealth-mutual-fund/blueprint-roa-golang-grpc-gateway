package users

import (
	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/repository/postgres"
)

type UsersService struct {
	repository postgres.Repository
}

func NewService(r postgres.Repository) (service Service) {
	return &UsersService{
		repository: r,
	}
}
