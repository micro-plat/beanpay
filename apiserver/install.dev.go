// +build !prod

package main

import (
	"github.com/micro-plat/beanpay/beanpay/const/sql"
	"github.com/micro-plat/hydra/component"
)

//bindConf 绑定启动配置， 启动时检查注册中心配置是否存在，不存在则引导用户输入配置参数并自动创建到注册中心
func (s *apiserver) install() {
	s.IsDebug = true
	s.Conf.API.SetMainConf(`{"address":":9090"}`)
	// s.Conf.Plat.SetVarConf("db", "db", `{
	// 		"provider":"mysql",
	// 		"connString":"hydra:123456@tcp(192.168.0.36)/hydra?charset=utf8",
	// 		"maxOpen":20,
	// 		"maxIdle":10,
	// 		"lifeTime":600
	// }`)

	s.Conf.Plat.SetVarConf("db", "db", `{
			"provider":"ora",
			"connString":"hydra/123456@orcl136",
			"maxOpen":20,
			"maxIdle":10,
			"lifeTime":600
	}`)

	s.Conf.API.Installer(func(c component.IContainer) error {
		if !s.Conf.Confirm("创建数据库表结构?") {
			return nil
		}
		//创建数据库
		db, err := c.GetDB()
		if err != nil {
			return err
		}
		err = sql.CreateDB(db)
		if err != nil {
			return err
		}
		c.GetLogger().Info("\t\t数据表创建成功")
		return nil
	})

}
