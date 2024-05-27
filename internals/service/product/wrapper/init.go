package wrapper

import (
	"go.uber.org/dig"

	service "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/service/product"
)

type Wrapper struct {
	dig.In  `name:"wrapperProduct"`
	Service service.Service
}
