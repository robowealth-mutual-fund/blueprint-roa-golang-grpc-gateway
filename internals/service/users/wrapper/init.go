package wrapper

import (
	"go.uber.org/dig"

	service "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/service/users"
)

type Wrapper struct {
	dig.In  `name:"wrapperUsers"`
	Service service.Service
}
