package account

import (
	"github.com/micro-plat/beanpay/beanpay"
	"github.com/micro-plat/hydra"
)

// AccountHandler AccountHandler
type AccountHandler struct {
}

// NewAccountHandler NewAccountHandler
func NewAccountHandler() (u *AccountHandler) {
	return &AccountHandler{}
}

//CreateHandle 创建资金帐户
func (u *AccountHandler) CreateHandle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("---------------创建资金帐户--------------------")
	ctx.Log().Info("1. 参数校验")
	if err := ctx.Request().Check("ident", "group", "name", "eid"); err != nil {
		return err
	}

	ctx.Log().Info("2. 创建帐户信息")
	bp := beanpay.GetAccount(ctx.Request().GetString("ident"), ctx.Request().GetString("group"))
	account, err := bp.CreateAccount(ctx,
		ctx.Request().GetString("eid"),
		ctx.Request().GetString("name"))
	if err != nil {
		return err
	}

	ctx.Log().Info("3. 处理返回结果")
	return account
}

//UpdateHandle 创建资金帐户
func (u *AccountHandler) UpdateHandle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("---------------修改资金帐户--------------------")
	ctx.Log().Info("1. 参数校验")
	if err := ctx.Request().Check("ident", "group", "name", "eid"); err != nil {
		return err
	}

	ctx.Log().Info("2. 修改资金帐户")
	bp := beanpay.GetAccount(ctx.Request().GetString("ident"), ctx.Request().GetString("group"))
	account, err := bp.SetAccountName(ctx,
		ctx.Request().GetString("eid"),
		ctx.Request().GetString("name"))
	if err != nil {
		return err
	}

	ctx.Log().Info("3. 处理返回结果")
	return account
}

//QueryHandle 查询资金帐户
func (u *AccountHandler) QueryHandle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("---------------查询资金帐户--------------------")
	ctx.Log().Info("1. 参数校验")
	if err := ctx.Request().Check("ident", "group", "eid"); err != nil {
		return err
	}

	ctx.Log().Info("2. 查询帐户信息")
	bp := beanpay.GetAccount(ctx.Request().GetString("ident"), ctx.Request().GetString("group"))
	account, err := bp.GetAccount(ctx, ctx.Request().GetString("eid"))
	if err != nil {
		return err
	}

	ctx.Log().Info("3. 处理返回结果")
	return account
}

//QueryListHandle 查询资金帐户列表
func (u *AccountHandler) QueryListHandle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("---------------查询资金帐户列表--------------------")
	ctx.Log().Info("1. 参数校验")
	if err := ctx.Request().Check("ident"); err != nil {
		return err
	}

	ctx.Log().Info("2. 查询帐户信息")
	bp := beanpay.GetAccount(ctx.Request().GetString("ident"), ctx.Request().GetString("group"))
	account, err := bp.QueryAccount(ctx,
		ctx.Request().GetString("eid"),
		ctx.Request().GetString("account_type"),
		ctx.Request().GetString("name"),
		ctx.Request().GetString("status"),
		ctx.Request().GetInt("pi", 1),
		ctx.Request().GetInt("ps", 10))
	if err != nil {
		return err
	}

	ctx.Log().Info("3. 处理返回结果")
	return account
}
