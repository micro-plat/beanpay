package main

import "github.com/micro-plat/hydra/hydra"
import _ "github.com/go-sql-driver/mysql"

type apiserver struct {
	*hydra.MicroApp
}

func main() {
	app := &apiserver{
		hydra.NewApp(
			hydra.WithPlatName("beanpay"),
			hydra.WithSystemName("apiserver"),
			hydra.WithServerTypes("api"),
			hydra.WithDebug()),
	}

	app.init()
	app.install()
	app.handling()

	app.Start()
}
