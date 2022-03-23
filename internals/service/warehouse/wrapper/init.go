package wrapper

import (
	service "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/service/warehouse"
	"go.uber.org/dig"
)

type Wrapper struct {
	dig.In  `name:"wrapperWarehouse"`
	Service service.Service
}

func WrapWarehouse(service service.Service) service.Service {
	return &Wrapper{
		Service: service,
	}
}
