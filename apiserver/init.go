package main

import (
	"github.com/micro-plat/beanpay/apiserver/services/account"
	"github.com/micro-plat/beanpay/apiserver/services/pkg"
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/hydra"
	"github.com/urfave/cli"
)

//init 检查应用程序配置文件，并根据配置初始化服务
func (r *apiserver) init() {

	r.Cli.Append(hydra.ModeRun, cli.BoolFlag{
		Name:  "pkg,p",
		Usage: "注册package服务",
	})

	r.Initializing(func(c component.IContainer) error {

		if _, err := c.GetDB(); err != nil {
			return err
		}

		r.Micro("/account", account.NewAccountHandler)
		r.Micro("/account/balance", account.NewBalanceHandler)
		r.Micro("/account/record", account.NewRecordHandler)

		if r.Cli.Context().Bool("pkg") {
			r.Micro("/package", pkg.NewPackageHandler)
			r.Micro("/package/capacity", pkg.NewCapacityHandler)
			r.Micro("/package/record", pkg.NewRecordHandler)
		}

		return nil
	})
}
