package account

import (
	"encoding/json"

	"github.com/micro-plat/beanpay/beanpay"
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
)

// RecordHandler .
type RecordHandler struct {
	container component.IContainer
}

// NewRecordHandler ．
func NewRecordHandler(container component.IContainer) (u *RecordHandler) {
	return &RecordHandler{container: container}
}

//QueryHandle 帐户资金变动查询
func (u *RecordHandler) QueryHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("---------------帐户资金变动查询--------------------")
	ctx.Log.Info("1. 参数校验")
	if err := ctx.Request.Check("ident"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2. 查询数据")
	bp := beanpay.GetAccount(ctx.Request.GetString("ident"), ctx.Request.GetString("group"))
	data, err := bp.QueryAccountRecords(ctx,
		ctx.Request.GetString("account_type"),
		ctx.Request.GetString("account_id"),
		ctx.Request.GetString("group"),
		ctx.Request.GetString("change_type"),
		ctx.Request.GetString("trade_type"),
		ctx.Request.GetString("eid"),
		ctx.Request.GetString("start_time"),
		ctx.Request.GetString("end_time"),
		ctx.Request.GetInt("pi", 1),
		ctx.Request.GetInt("ps", 10))
	if err != nil {
		return err
	}
	temp, err := json.Marshal(data)
	if err != nil {
		return err
	}
	ctx.Log.Info("3. 处理返回结果")
	return string(temp)
}
