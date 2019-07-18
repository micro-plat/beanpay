package pkg

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
)

type CapacityHandler struct {
	container component.IContainer
}

func NewCapacityHandler(container component.IContainer) (u *CapacityHandler) {
	return &CapacityHandler{container: container}
}

//Handle .
func (u *CapacityHandler) Handle(ctx *context.Context) (r interface{}) {
	return "success"
}
