package pkg

import (
	"fmt"

	"github.com/micro-plat/beanpay/beanpay/const/sql"
	"github.com/micro-plat/lib4go/db"
)

//Query 查询服务包变动记录
func query(db db.IDBExecuter, pkgID int, startTime string, pi int, ps int) (db.QueryRows, error) {
	input := map[string]interface{}{
		"pkg_id": pkgID,
		"start":  startTime,
		"pi":     pi,
		"ps":     ps,
	}
	rows, s, p, err := db.Query(sql.QueryPackageRecord, input)
	if err != nil {
		return nil, fmt.Errorf("SQL语句执行出错:%s(%v)", s, p)
	}
	return rows, nil
}
