package wrapper

import (
	service "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/service/product"
	"go.uber.org/dig"
)

type Wrapper struct {
	dig.In  `name:"wrapperProduct"`
	Service service.Service
}

func WrapProduct(service service.Service) service.Service {
	return &Wrapper{
		Service: service,
	}
}
