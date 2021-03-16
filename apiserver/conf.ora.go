// +build oracle

package main

import (
	_ "github.com/mattn/go-oci8"
	"github.com/micro-plat/hydra"
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
