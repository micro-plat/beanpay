package account

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
)

type AccountHandler struct {
	container component.IContainer
}

func NewAccountHandler(container component.IContainer) (u *AccountHandler) {
	return &AccountHandler{container: container}
}

//Handle .
func (u *AccountHandler) Handle(ctx *context.Context) (r interface{}) {
	return "success"
}
