package pkg

import (
	"github.com/micro-plat/beanpay/beanpay"
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
)

type PackageHandler struct {
	container component.IContainer
}

func NewPackageHandler(container component.IContainer) (u *PackageHandler) {
	return &PackageHandler{container: container}
}

//CreateHandle 创建服务包
func (u *PackageHandler) CreateHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("---------------创建服务包--------------------")
	ctx.Log.Info("1. 参数校验")
	if err := ctx.Request.Check("eid", "sid", "name", "total"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2. 创建服务包")
	pkg, err := beanpay.CreatePackage(ctx,
		ctx.Request.GetString("eid"),
		ctx.Request.GetString("sid"),
		ctx.Request.GetString("name"),
		ctx.Request.GetInt("total"),
		ctx.Request.GetInt("daily"),
		ctx.Request.GetString("expires", "20991231"))
	if err != nil {
		return err
	}

	ctx.Log.Info("3. 处理返回结果")
	return pkg
}

//QueryHandle 查询服务包
func (u *PackageHandler) QueryHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("---------------查询服务包--------------------")
	ctx.Log.Info("1. 参数校验")
	if err := ctx.Request.Check("eid", "sid"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2. 查询服务包")
	pkg, err := beanpay.GetPackage(ctx,
		ctx.Request.GetString("eid"),
		ctx.Request.GetString("sid"))
	if err != nil {
		return err
	}

	ctx.Log.Info("3. 处理返回结果")
	return pkg
}
