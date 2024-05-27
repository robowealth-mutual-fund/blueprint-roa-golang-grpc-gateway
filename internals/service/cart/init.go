package cartLiveing

import (
	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/repository/postgres"
)

type CartService struct {
	repository postgres.Repository
}

func NewService(r postgres.Repository) (service Service) {
	return &CartService{
		repository: r,
	}
}
