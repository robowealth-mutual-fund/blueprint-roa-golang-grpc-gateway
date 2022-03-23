package wrapper

import (
	service "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/service/category"
	"go.uber.org/dig"
)

type Wrapper struct {
	dig.In  `name:"wrapperCategory"`
	Service service.Service
}

func WrapCategory(service service.Service) service.Service {
	return &Wrapper{
		Service: service,
	}
}
