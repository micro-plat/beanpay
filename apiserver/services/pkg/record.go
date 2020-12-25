package pkg

import (
	"github.com/micro-plat/beanpay/beanpay"
	"github.com/micro-plat/hydra"
)

// RecordHandler 结构体
type RecordHandler struct {
}

// NewRecordHandler 构建RecordHandler结构体
func NewRecordHandler() (u *RecordHandler) {
	return &RecordHandler{}
}

//QueryHandle 服务包数量变动查询
func (u *RecordHandler) QueryHandle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("---------------服务包数量变动查询--------------------")
	ctx.Log().Info("1. 参数校验")
	if err := ctx.Request().Check("ident", "group", "eid", "start_time", "end_time"); err != nil {
		return err
	}

	ctx.Log().Info("2. 服务包数量变动查询")
	bp := beanpay.GetPackage(ctx.Request().GetString("ident"), ctx.Request().GetString("group"))
	data, err := bp.QueryPackageRecords(ctx,
		ctx.Request().GetString("eid"),
		ctx.Request().GetString("spid"),
		ctx.Request().GetString("start_time"),
		ctx.Request().GetString("end_time"),
		ctx.Request().GetInt("pi", 0),
		ctx.Request().GetInt("ps", 10))
	if err != nil {
		return err
	}

	ctx.Log().Info("3. 处理返回结果")
	return data
}
