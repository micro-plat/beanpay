package main

import (
	"github.com/micro-plat/beanpay/beanpay/const/sql/creator"
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/conf"
)

//bindConf 绑定启动配置， 启动时检查注册中心配置是否存在，不存在则引导用户输入配置参数并自动创建到注册中心
func init() {
	app.Conf.API.SetMain(conf.NewAPIServerConf(":9090"))
	app.Conf.RPC.SetMain(conf.NewRPCServerConf(":9091"))

	app.Conf.API.Installer(func(c component.IContainer) error {
		if !app.Conf.Confirm("创建数据库表结构?") {
			return nil
		}
		//创建数据库
		db, err := c.GetDB()
		if err != nil {
			return err
		}
		err = creator.CreateDB(db)
		if err != nil {
			return err
		}
		c.GetLogger().Info("\t\t数据表创建成功")
		return nil
	})

}
