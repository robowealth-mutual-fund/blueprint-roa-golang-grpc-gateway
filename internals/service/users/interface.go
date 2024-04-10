package users

import (
	"context"

	model "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/model/users"
)

//go:generate mockery --name=Service
type Service interface {
	Create(ctx context.Context, input *model.Request) (ID int, err error)
}
