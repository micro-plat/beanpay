// +build oracle

package creator

import (
	"github.com/micro-plat/lib4go/db"
	_ "github.com/zkfy/go-oci8"
)

func CreateDB(xdb db.IDBExecuter) error {
	return db.CreateDB(xdb, "src/github.com/micro-plat/beanpay/beanpay/const/sql/oracle")
}
