// +build prod
// +build !oci

package main

import _ "github.com/go-sql-driver/mysql"

//bindConf 绑定启动配置， 启动时检查注册中心配置是否存在，不存在则引导用户输入配置参数并自动创建到注册中心
func (s *apiserver) installs() {
	s.Conf.SetInput("db_cstring", "数据库连接串", "mysql:uName:pwd@tcp(serverip)/dbName?charset=utf8")
	s.Conf.Plat.SetVarConf("db", "db", `{
			"provider":"mysql",
			"connString":"#db_cstring",
			"maxOpen":20,
			"maxIdle":10,
			"lifeTime":600
	}`)
}
