package account

import (
	"github.com/micro-plat/beanpay/beanpay/account"
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
)

type RecordHandler struct {
	container component.IContainer
}

func NewRecordHandler(container component.IContainer) (u *RecordHandler) {
	return &RecordHandler{container: container}
}

//QueryHandle 帐户资金变动查询
func (u *RecordHandler) QueryHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("---------------帐户资金变动查询--------------------")
	ctx.Log.Info("1. 参数校验")
	if err := ctx.Request.Check("uaid", "start_time"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2. 查询数据")
	data, err := account.Query(u.container.GetRegularDB(),
		ctx.Request.GetString("uaid"),
		ctx.Request.GetString("start_time"),
		ctx.Request.GetInt("pi", 0),
		ctx.Request.GetInt("ps", 10))
	if err != nil {
		return err
	}

	ctx.Log.Info("3. 处理返回结果")
	return data
}
