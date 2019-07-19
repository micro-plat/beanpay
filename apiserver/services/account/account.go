package account

import (
	"github.com/micro-plat/beanpay/beanpay/account"
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
)

type AccountHandler struct {
	container component.IContainer
}

func NewAccountHandler(container component.IContainer) (u *AccountHandler) {
	return &AccountHandler{container: container}
}

//Handle 创建资金帐户
func (u *AccountHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("---------------创建资金帐户--------------------")
	ctx.Log.Info("1. 参数校验")
	if err := ctx.Request.Check("uaid", "name"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2. 创建帐户信息")
	id, err := account.Create(u.container.GetRegularDB(),
		ctx.Request.GetString("uaid"),
		ctx.Request.GetString("name"))
	if err != nil {
		return err
	}

	ctx.Log.Info("3. 处理返回结果")
	return map[string]interface{}{
		"account_id": id,
		"name":       ctx.Request.GetString("name"),
	}
}
