package account

import (
	"github.com/micro-plat/beanpay/beanpay"
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
)

// BalanceHandler .
type BalanceHandler struct {
	container component.IContainer
}

// NewBalanceHandler .
func NewBalanceHandler(container component.IContainer) (u *BalanceHandler) {
	return &BalanceHandler{container: container}
}

//AddHandle 帐户加款
func (u *BalanceHandler) AddHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("---------------帐户加款--------------------")
	ctx.Log.Info("1. 参数校验")
	if err := ctx.Request.Check("ident", "group", "eid", "trade_no", "amount"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2. 帐户加款")
	bp := beanpay.GetAccount(ctx.Request.GetString("ident"), ctx.Request.GetString("group"))
	record, err := bp.AddAmount(ctx,
		ctx.Request.GetString("eid"),
		ctx.Request.GetString("trade_no"),
		ctx.Request.GetInt("amount"))
	if err != nil {
		return err
	}

	ctx.Log.Info("3. 处理返回结果")
	return record
}

//DrawingHandle 帐户提款
func (u *BalanceHandler) DrawingHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("---------------帐户提款--------------------")
	ctx.Log.Info("1. 参数校验")
	if err := ctx.Request.Check("ident", "group", "eid", "trade_no", "amount"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2. 帐户提款")
	bp := beanpay.GetAccount(ctx.Request.GetString("ident"), ctx.Request.GetString("group"))
	record, err := bp.DrawingAmount(ctx,
		ctx.Request.GetString("eid"),
		ctx.Request.GetString("trade_no"),
		ctx.Request.GetInt("amount"))
	if err != nil {
		return err
	}

	ctx.Log.Info("3. 处理返回结果")
	return record
}

//DeductHandle 帐户扣款
func (u *BalanceHandler) DeductHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("---------------帐户扣款--------------------")
	ctx.Log.Info("1. 参数校验")
	if err := ctx.Request.Check("ident", "group", "eid", "trade_no", "amount", "trade_type"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2. 帐户扣款")
	bp := beanpay.GetAccount(ctx.Request.GetString("ident"), ctx.Request.GetString("group"))
	record, err := bp.DeductAmount(ctx,
		ctx.Request.GetString("eid"),
		ctx.Request.GetString("trade_no"),
		ctx.Request.GetInt("trade_type"),
		ctx.Request.GetInt("amount"))
	if err != nil {
		return err
	}

	ctx.Log.Info("3. 处理返回结果")
	return record
}

//RefundHandle 帐户退款
func (u *BalanceHandler) RefundHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("---------------帐户退款--------------------")
	ctx.Log.Info("1. 参数校验")
	if err := ctx.Request.Check("ident", "group", "eid", "trade_no", "ext_no", "trade_type", "amount"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2. 帐户退款")
	bp := beanpay.GetAccount(ctx.Request.GetString("ident"), ctx.Request.GetString("group"))
	record, err := bp.RefundAmount(ctx,
		ctx.Request.GetString("eid"),
		ctx.Request.GetString("trade_no"),
		ctx.Request.GetString("ext_no"),
		ctx.Request.GetInt("trade_type"),
		ctx.Request.GetInt("amount"))
	if err != nil {
		return err
	}

	ctx.Log.Info("3. 处理返回结果")
	return record
}
