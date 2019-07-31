package main

import "github.com/micro-plat/hydra/hydra"

type apiserver struct {
	*hydra.MicroApp
}

func main() {
	app := &apiserver{
		hydra.NewApp(
			hydra.WithPlatName("beanpay"),
			hydra.WithSystemName("apiserver"),
			hydra.WithServerTypes("api")),
	}

	app.init()
	app.install()
	app.handling()

	app.Start()
}
