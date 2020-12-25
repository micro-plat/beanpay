package main

import (
	"github.com/micro-plat/beanpay/apiserver/services/account"
	"github.com/micro-plat/beanpay/apiserver/services/pkg"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithPlatName("beanpay"),
	hydra.WithSystemName("apiserver"),
	hydra.WithRunFlag("pkg", "是否注册package服务"),
	hydra.WithServerTypes(http.API),
)

func main() {

	//注册账户服务
	app.Micro("/account", account.NewAccountHandler)
	app.Micro("/account/balance", account.NewBalanceHandler)
	app.Micro("/account/record", account.NewRecordHandler)

	//注册服务包
	hydra.RunCli.OnStarting(func(c hydra.ICli) error {
		if c.IsSet("pkg") {
			app.Micro("/package", pkg.NewPackageHandler)
			app.Micro("/package/capacity", pkg.NewCapacityHandler)
			app.Micro("/package/record", pkg.NewRecordHandler)
		}
		return nil
	})

	//启动时检查配置
	app.OnStarting(func(conf hydra.IAPPConf) error {
		if _, err := hydra.C.DB().GetDB(); err != nil {
			return err
		}
		return nil
	}, http.API)
	app.Start()
}
