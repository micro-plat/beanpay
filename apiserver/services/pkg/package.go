package pkg

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
)

type PackageHandler struct {
	container component.IContainer
}

func NewPackageHandler(container component.IContainer) (u *PackageHandler) {
	return &PackageHandler{container: container}
}

//Handle .
func (u *PackageHandler) Handle(ctx *context.Context) (r interface{}) {
	return "success"
}
