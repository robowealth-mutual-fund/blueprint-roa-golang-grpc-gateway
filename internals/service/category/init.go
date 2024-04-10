package category

import (
	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/repository/postgres"
)

type CategoryService struct {
	repository postgres.Repository
}

func NewService(r postgres.Repository) (service Service) {
	return &CategoryService{
		repository: r,
	}
}
