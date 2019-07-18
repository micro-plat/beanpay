// +build !oci

package sql

import _ "github.com/go-sql-driver/mysql"
import "github.com/micro-plat/lib4go/db"

func CreateDB(xdb db.IDB) error {
	return db.CreateDB(xdb, "src/github.com/micro-plat/beanpay/modules/const/sql/mysql")
}
