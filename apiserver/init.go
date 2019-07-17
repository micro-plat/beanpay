package main

import (
	"github.com/micro-plat/beanpay/apiserver/services/account"
	"github.com/micro-plat/beanpay/apiserver/services/pkg"
	"github.com/micro-plat/hydra/component"
)

//init 检查应用程序配置文件，并根据配置初始化服务
func (r *apiserver) init() {
	r.Initializing(func(c component.IContainer) error {
		//appconf.func#//
		//#appconf.func//

		//db.init#//
		//#db.init//

		//cache.init#//
		//#cache.init//

		//queue.init#//
		//#queue.init//

		//login.router#//
		//#login.router//

		//service.router#//
		r.Micro("/account/create", account.NewAccountHandler)
		r.Micro("/account/balance", account.NewBalanceHandler)
		r.Micro("/account/record", account.NewRecordHandler)

		r.Micro("/package/create", pkg.NewPackageHandler)
		r.Micro("/package/capacity", pkg.NewCapacityHandler)
		r.Micro("/package/record", pkg.NewRecordHandler)

		//#service.router//

		return nil
	})
}
