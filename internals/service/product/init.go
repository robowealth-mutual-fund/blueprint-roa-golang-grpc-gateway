package wrapper

import (
	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/repository/postgres"
)

type ProductService struct {
	repository postgres.Repository
}

func NewService(r postgres.Repository) (service Service) {
	return &ProductService{
		repository: r,
	}
}
