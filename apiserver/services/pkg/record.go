package pkg

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
)

type RecordHandler struct {
	container component.IContainer
}

func NewRecordHandler(container component.IContainer) (u *RecordHandler) {
	return &RecordHandler{container: container}
}

//Handle .
func (u *RecordHandler) Handle(ctx *context.Context) (r interface{}) {
	return "success"
}
