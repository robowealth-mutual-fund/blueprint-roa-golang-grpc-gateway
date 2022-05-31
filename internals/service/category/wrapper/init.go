package wrapper

import (
	"go.uber.org/dig"

	service "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/service/category"
)

type Wrapper struct {
	dig.In  `name:"wrapperCategory"`
	Service service.Service
}

func _(service service.Service) service.Service {
	return &Wrapper{
		Service: service,
	}
}
