package warehouse

import (
	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/repository/postgres"
)

type WarehouseService struct {
	repository postgres.Repository
}

func NewService(r postgres.Repository) (service Service) {
	return &WarehouseService{
		repository: r,
	}
}
