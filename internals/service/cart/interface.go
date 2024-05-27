package cartLiveing

import (
	"context"

	model "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/model/cart"
)

//go:generate mockery --name=Service
type Service interface {
	Create(ctx context.Context, input *model.Request) (name string, err error)
}
