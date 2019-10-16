// +build !prod
// +build oracle

package main

import (
	"github.com/micro-plat/hydra/conf"
	_ "github.com/zkfy/go-oci8"
)

//bindConf 绑定启动配置， 启动时检查注册中心配置是否存在，不存在则引导用户输入配置参数并自动创建到注册中心
func init() {
	s.IsDebug = true
	app.Conf.Plat.SetDB(conf.NewOracleConf("hydra", "123456", "orcl136"))
}
