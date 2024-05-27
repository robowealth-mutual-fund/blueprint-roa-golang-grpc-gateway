package wrapper

import (
	"context"

	model "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/model/cart"

	"github.com/opentracing/opentracing-go"
)

func (wrp *Wrapper) Create(ctx context.Context, input *model.Request) (string, error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "Service.Product.Create")
	defer sp.Finish()

	sp.LogKV("Brand", input.Brand)

	name, err := wrp.Service.Create(ctx, input)

	sp.LogKV("ID", name)
	sp.LogKV("err", err)

	return name, err
}
