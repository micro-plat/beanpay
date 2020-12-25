package pkg

import (
	"github.com/micro-plat/beanpay/beanpay"
	"github.com/micro-plat/hydra"
)

// PackageHandler 结构体
type PackageHandler struct {
}

// NewPackageHandler 构建PackageHandler
func NewPackageHandler() (u *PackageHandler) {
	return &PackageHandler{}
}

//CreateHandle 创建服务包
func (u *PackageHandler) CreateHandle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("---------------创建服务包--------------------")
	ctx.Log().Info("1. 参数校验")
	if err := ctx.Request().Check("ident", "group", "eid", "spid", "name", "total"); err != nil {
		return err
	}

	ctx.Log().Info("2. 创建服务包")
	bp := beanpay.GetPackage(ctx.Request().GetString("ident"), ctx.Request().GetString("group"))
	pkg, err := bp.CreatePackage(ctx,
		ctx.Request().GetString("eid"),
		ctx.Request().GetString("spid"),
		ctx.Request().GetString("name"),
		ctx.Request().GetInt("total"),
		ctx.Request().GetInt("daily"),
		ctx.Request().GetString("expires", "20991231"))
	if err != nil {
		return err
	}

	ctx.Log().Info("3. 处理返回结果")
	return pkg
}

//QueryHandle 查询服务包
func (u *PackageHandler) QueryHandle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("---------------查询服务包--------------------")
	ctx.Log().Info("1. 参数校验")
	if err := ctx.Request().Check("ident", "group", "eid", "spid"); err != nil {
		return err
	}

	ctx.Log().Info("2. 查询服务包")
	bp := beanpay.GetPackage(ctx.Request().GetString("ident"), ctx.Request().GetString("group"))
	pkg, err := bp.GetPackage(ctx,
		ctx.Request().GetString("eid"),
		ctx.Request().GetString("spid"))
	if err != nil {
		return err
	}

	ctx.Log().Info("3. 处理返回结果")
	return pkg
}
