package warehouse

import (
	"context"
	model "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/model/warehouse"
)

//go:generate mockery --name=Service
type Service interface {
	Create(ctx context.Context, input *model.Request) (ID int, err error)
}
