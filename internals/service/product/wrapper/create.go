package wrapper

import (
	"context"
	model "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/model/product"

	"github.com/opentracing/opentracing-go"
)

func (wrp *Wrapper) Create(ctx context.Context, input *model.Request) (int, error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "Service.Product.Create")
	defer sp.Finish()

	sp.LogKV("Brand", input.Brand)

	id, err := wrp.Service.Create(ctx, input)

	sp.LogKV("ID", id)
	sp.LogKV("err", err)

	return id, err
}
