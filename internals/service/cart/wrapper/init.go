package wrapper

import (
	"go.uber.org/dig"

	service "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/service/cart"
)

type Wrapper struct {
	dig.In  `name:"wrapperCart"`
	Service service.Service
}
