// +build !oracle

package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/micro-plat/hydra"
)

func init() {
	hydra.OnReady(func() error {
		// if hydra.G.IsDebug() {
		hydra.Conf.Vars().DB().MySQL("db", "hydra", "123456", "192.168.0.36", "hydra")
		// return nil
		// }
		// hydra.Conf.Vars().DB().MySQLByConnStr("db", hydra.ByInstall)
		return nil
	})
}
