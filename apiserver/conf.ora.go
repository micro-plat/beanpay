// +build oracle

package main

import (
	"github.com/micro-plat/hydra"
	_ "github.com/zkfy/go-oci8"
)

func init() {

	hydra.OnReady(func() error {
		if hydra.G.IsDebug() {
			hydra.Conf.Vars().DB().Oracle("db", "hydra", "123456", "orcl136")
			return nil
		}
		hydra.Conf.Vars().DB().OracleByConnStr("db", hydra.ByInstall)
		return nil
	})

}
