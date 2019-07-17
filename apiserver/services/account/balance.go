
package account

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
)

type BalanceHandler struct {
	container component.IContainer
}

func NewBalanceHandler(container component.IContainer) (u *BalanceHandler) {
	return &BalanceHandler{container: container}
}

//Handle .
func (u *BalanceHandler) Handle(ctx *context.Context) (r interface{}) {
	return "success"
}
