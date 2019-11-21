// +build !oracle

package creator

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/micro-plat/lib4go/db"
)

func CreateDB(xdb db.IDBExecuter) error {
	return db.CreateDB(xdb, "src/github.com/micro-plat/beanpay/beanpay/const/sql/mysql")
}
