package account

import (
	"github.com/micro-plat/beanpay/beanpay"
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
)

type AccountHandler struct {
	container component.IContainer
}

func NewAccountHandler(container component.IContainer) (u *AccountHandler) {
	return &AccountHandler{container: container}
}

//CreateHandle 创建资金帐户
func (u *AccountHandler) CreateHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("---------------创建资金帐户--------------------")
	ctx.Log.Info("1. 参数校验")
	if err := ctx.Request.Check("sid", "eid", "name"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2. 创建帐户信息")
	bp := beanpay.NewBeanpay(ctx.Request.GetString("sid"), ctx.Request.GetString("tp"))
	account, err := bp.CreateAccount(ctx,
		ctx.Request.GetString("eid"),
		ctx.Request.GetString("name"))
	if err != nil {
		return err
	}

	ctx.Log.Info("3. 处理返回结果")
	return account
}

//QueryHandle 查询资金帐户
func (u *AccountHandler) QueryHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("---------------查询资金帐户--------------------")
	ctx.Log.Info("1. 参数校验")
	if err := ctx.Request.Check("sid", "eid"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2. 查询帐户信息")
	bp := beanpay.NewBeanpay(ctx.Request.GetString("sid"), ctx.Request.GetString("tp"))
	account, err := bp.GetAccount(ctx, ctx.Request.GetString("eid"))
	if err != nil {
		return err
	}

	ctx.Log.Info("3. 处理返回结果")
	return account
}
