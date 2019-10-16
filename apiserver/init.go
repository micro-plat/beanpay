package main

import (
	"github.com/micro-plat/beanpay/apiserver/services/account"
	"github.com/micro-plat/beanpay/apiserver/services/pkg"
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/hydra/hydra"
	"github.com/urfave/cli"
)

//init 检查应用程序配置文件，并根据配置初始化服务
func init() {

	app.Handling(func(ctx *context.Context) (rt interface{}) {
		return nil
	})
	app.Cli.Append(hydra.ModeRun, cli.BoolFlag{
		Name:  "pkg,p",
		Usage: "注册package服务",
	})

	app.Initializing(func(c component.IContainer) error {

		if _, err := c.GetDB(); err != nil {
			return err
		}

		app.Micro("/account", account.NewAccountHandler)
		app.Micro("/account/balance", account.NewBalanceHandler)
		app.Micro("/account/record", account.NewRecordHandler)

		if app.Cli.Context().Bool("pkg") {
			app.Micro("/package", pkg.NewPackageHandler)
			app.Micro("/package/capacity", pkg.NewCapacityHandler)
			app.Micro("/package/record", pkg.NewRecordHandler)
		}

		return nil
	})
}
