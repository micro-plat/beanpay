package main

import "github.com/micro-plat/hydra/hydra"

type apiserver struct {
	*hydra.MicroApp
}

var app = &apiserver{
	hydra.NewApp(
		hydra.WithPlatName("beanpay"),
		hydra.WithSystemName("apiserver"),
		hydra.WithServerTypes("api-rpc")),
}

func main() {
	app.Start()
}
