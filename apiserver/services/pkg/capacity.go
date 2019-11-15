package pkg

import (
	"github.com/micro-plat/beanpay/beanpay"
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
)

// CapacityHandler ./
type CapacityHandler struct {
	container component.IContainer
}

// NewCapacityHandler 构建CapacityHandler
func NewCapacityHandler(container component.IContainer) (u *CapacityHandler) {
	return &CapacityHandler{container: container}
}

//AddHandle 添加服务包数量
func (u *CapacityHandler) AddHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("---------------添加服务包数量--------------------")
	ctx.Log.Info("1. 参数校验")
	if err := ctx.Request.Check("ident", "group", "eid", "spid", "trade_no", "num"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2. 添加服务包数量")
	bp := beanpay.GetPackage(ctx.Request.GetString("ident"), ctx.Request.GetString("group"))
	pkg, err := bp.AddCapacity(ctx,
		ctx.Request.GetString("eid"),
		ctx.Request.GetString("spid"),
		ctx.Request.GetString("trade_no"),
		ctx.Request.GetInt("num"))
	if err != nil {
		return err
	}

	ctx.Log.Info("3. 处理返回结果")
	return pkg
}

//DrawingHandle 提取服务包数量
func (u *CapacityHandler) DrawingHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("---------------提取服务包数量--------------------")
	ctx.Log.Info("1. 参数校验")
	if err := ctx.Request.Check("ident", "group", "eid", "spid", "trade_no", "num"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2. 提取服务包数量")
	bp := beanpay.GetPackage(ctx.Request.GetString("ident"), ctx.Request.GetString("group"))
	pkg, err := bp.DrawingCapacity(ctx,
		ctx.Request.GetString("eid"),
		ctx.Request.GetString("spid"),
		ctx.Request.GetString("trade_no"),
		ctx.Request.GetInt("num"))
	if err != nil {
		return err
	}

	ctx.Log.Info("3. 处理返回结果")
	return pkg
}

//DeductHandle 扣减服务包数量
func (u *CapacityHandler) DeductHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("---------------扣减服务包数量--------------------")
	ctx.Log.Info("1. 参数校验")
	if err := ctx.Request.Check("ident", "group", "eid", "spid", "trade_no", "num"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2. 扣减服务包数量")
	bp := beanpay.GetPackage(ctx.Request.GetString("ident"), ctx.Request.GetString("group"))
	pkg, err := bp.DeductCapacity(ctx,
		ctx.Request.GetString("eid"),
		ctx.Request.GetString("spid"),
		ctx.Request.GetString("trade_no"),
		ctx.Request.GetInt("num"))
	if err != nil {
		return err
	}

	ctx.Log.Info("3. 处理返回结果")
	return pkg
}

//RefundHandle 退回服务包数量
func (u *CapacityHandler) RefundHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("---------------退回服务包数量--------------------")
	ctx.Log.Info("1. 参数校验")
	if err := ctx.Request.Check("ident", "group", "eid", "spid", "trade_no", "num"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2. 退回服务包数量")
	bp := beanpay.GetPackage(ctx.Request.GetString("ident"), ctx.Request.GetString("group"))
	pkg, err := bp.RefundCapacity(ctx,
		ctx.Request.GetString("eid"),
		ctx.Request.GetString("spid"),
		ctx.Request.GetString("trade_no"),
		ctx.Request.GetInt("num"))
	if err != nil {
		return err
	}

	ctx.Log.Info("3. 处理返回结果")
	return pkg
}
