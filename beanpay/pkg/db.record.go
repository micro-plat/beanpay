package pkg

import (
	"fmt"

	"github.com/micro-plat/beanpay/beanpay/const/sql"
	"github.com/micro-plat/lib4go/db"
)

//Query 查询服务包变动记录
func query(db db.IDBExecuter, aid int, pkg_id int64, startTime string, endTime string, pi int, ps int) (db.QueryRows, error) {
	input := map[string]interface{}{
		"account_id": aid,
		"pkg_id":     pkg_id,
		"start":      startTime,
		"end":        endTime,
		"pf":         pi * ps,
		"pi":         pi,
		"ps":         ps,
	}
	rows, err := db.Query(sql.QueryPackageRecord, input)
	if err != nil {
		return nil, fmt.Errorf("SQL语句执行出错:%w", err)
	}
	return rows, nil
}
