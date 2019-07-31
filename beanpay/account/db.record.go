package account

import (
	"fmt"

	"github.com/micro-plat/beanpay/beanpay/const/sql"
	"github.com/micro-plat/lib4go/db"
)

//Query 查询余额变动明细
func query(db db.IDBExecuter, accountID int, startTime string, endTime string, pi int, ps int) (db.QueryRows, error) {
	input := map[string]interface{}{
		"account_id": accountID,
		"start":      startTime,
		"end":        endTime,
		"pf":         pi * ps,
		"pi":         pi,
		"ps":         ps,
	}
	rows, s, p, err := db.Query(sql.QueryBalanceRecord, input)
	if err != nil {
		return nil, fmt.Errorf("SQL语句执行出错:%s(%v)", s, p)
	}
	return rows, nil
}
