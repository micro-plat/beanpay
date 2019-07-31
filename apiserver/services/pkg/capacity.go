package pkg

import (
	"github.com/micro-plat/beanpay/beanpay"
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
)

type CapacityHandler struct {
	container component.IContainer
}

func NewCapacityHandler(container component.IContainer) (u *CapacityHandler) {
	return &CapacityHandler{container: container}
}

//AddHandle 添加服务包数量
func (u *CapacityHandler) AddHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("---------------添加服务包数量--------------------")
	ctx.Log.Info("1. 参数校验")
	if err := ctx.Request.Check("uid", "sid", "trade_no", "capacity"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2. 添加服务包数量")
	err := beanpay.AddCapacity(ctx,
		ctx.Request.GetString("uid"),
		ctx.Request.GetString("sid"),
		ctx.Request.GetString("trade_no"),
		ctx.Request.GetInt("capacity"))
	if err != nil {
		return err
	}

	ctx.Log.Info("3. 处理返回结果")
	return "success"
}

//DeductHandle 扣减服务包数量
func (u *CapacityHandler) DeductHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("---------------扣减服务包数量--------------------")
	ctx.Log.Info("1. 参数校验")
	if err := ctx.Request.Check("uid", "sid", "trade_no", "capacity"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2. 扣减服务包数量")
	err := beanpay.DeductCapacity(ctx,
		ctx.Request.GetString("uid"),
		ctx.Request.GetString("sid"),
		ctx.Request.GetString("trade_no"),
		ctx.Request.GetInt("capacity"))
	if err != nil {
		return err
	}

	ctx.Log.Info("3. 处理返回结果")
	return "success"
}

//RefundHandle 退回服务包数量
func (u *CapacityHandler) RefundHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("---------------退回服务包数量--------------------")
	ctx.Log.Info("1. 参数校验")
	if err := ctx.Request.Check("uid", "sid", "trade_no", "capacity"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2. 退回服务包数量")
	err := beanpay.RefundCapacity(ctx,
		ctx.Request.GetString("uid"),
		ctx.Request.GetString("sid"),
		ctx.Request.GetString("trade_no"),
		ctx.Request.GetInt("capacity"))
	if err != nil {
		return err
	}

	ctx.Log.Info("3. 处理返回结果")
	return "success"
}
