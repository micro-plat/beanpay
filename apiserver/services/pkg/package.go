package pkg

import (
	"github.com/micro-plat/beanpay/beanpay/pkg"
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
)

type PackageHandler struct {
	container component.IContainer
}

func NewPackageHandler(container component.IContainer) (u *PackageHandler) {
	return &PackageHandler{container: container}
}

//Handle 创建资金帐户
func (u *PackageHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("---------------创建服务包--------------------")
	ctx.Log.Info("1. 参数校验")
	if err := ctx.Request.Check("uaid", "spkg_id", "name", "total"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2. 创建服务包")
	id, err := pkg.Create(u.container.GetRegularDB(),
		ctx.Request.GetString("uaid"),
		ctx.Request.GetString("spkg_id"),
		ctx.Request.GetString("name"),
		ctx.Request.GetInt("total"),
		ctx.Request.GetInt("daily"),
		ctx.Request.GetString("expires", "20991231"))
	if err != nil {
		return err
	}

	ctx.Log.Info("3. 处理返回结果")
	return map[string]interface{}{
		"pkg_id":  id,
		"spkg_id": ctx.Request.GetString("pkg_id"),
		"name":    ctx.Request.GetString("name"),
	}
}
