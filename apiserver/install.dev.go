// +build !prod

package main

//bindConf 绑定启动配置， 启动时检查注册中心配置是否存在，不存在则引导用户输入配置参数并自动创建到注册中心
func (s *apiserver) install() {
	s.IsDebug = true
	s.Conf.API.SetMainConf(`{"address":":9090"}`)
	// s.Conf.Plat.SetVarConf("db", "db", `{
	// 		"provider":"mysql",
	// 		"connString":"mrss:123456@tcp(192.168.0.36)/mrss?charset=utf8",
	// 		"maxOpen":20,
	// 		"maxIdle":10,
	// 		"lifeTime":600
	// }`)

	s.Conf.Plat.SetVarConf("db", "db", `{
			"provider":"ora",
			"connString":"sso/123456@orcl136",
			"maxOpen":20,
			"maxIdle":10,
			"lifeTime":600
	}`)

}
