package product

import (
	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/service/product/wrapper"
)

type Controller struct {
	service wrapper.Wrapper
}

func NewController(s wrapper.Wrapper) *Controller {
	return &Controller{
		service: s,
	}
}
