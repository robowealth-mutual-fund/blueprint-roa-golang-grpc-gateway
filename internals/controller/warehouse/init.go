package warehouse

import (
	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/service/warehouse/wrapper"
)

type Controller struct {
	service wrapper.Wrapper
}

func NewController(s wrapper.Wrapper) *Controller {
	return &Controller{
		service: s,
	}
}
