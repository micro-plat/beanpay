// +build !prod
// +build !oracle

package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/micro-plat/hydra/conf"
)

//bindConf 绑定启动配置， 启动时检查注册中心配置是否存在，不存在则引导用户输入配置参数并自动创建到注册中心
func init() {
	app.IsDebug = true
	app.Conf.Plat.SetDB(conf.NewMysqlConf("hydra", "123456", "192.168.0.36", "hydra"))
}
