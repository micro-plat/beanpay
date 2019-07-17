// +build oci

package sql

import _ "github.com/zkfy/go-oci8"
import "github.com/micro-plat/lib4go/db"

func CreateDB(xdb db.IDB) error {
	return db.CreateDB(xdb, "src/github.com/micro-plat/beanpay/apiserver/modules/const/sql/oracle")
}
