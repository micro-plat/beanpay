package account

import (
	"fmt"

	"github.com/micro-plat/beanpay/beanpay/const/sql"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
)

//Query 查询余额变动明细
func query(db db.IDBExecuter, accountID int, startTime string, endTime string, pi int, ps int) (int, db.QueryRows, error) {
	input := map[string]interface{}{
		"account_id": accountID,
		"start":      startTime,
		"end":        endTime,
		"pf":         pi * ps,
		"pi":         pi,
		"ps":         ps,
	}
	count, s, p, err := db.Scalar(sql.QueryBalanceRecordCount, input)
	if err != nil {
		return 0, nil, fmt.Errorf("SQL语句执行出错:%s(%v)", s, p)
	}

	rows, s, p, err := db.Query(sql.QueryBalanceRecord, input)
	if err != nil {
		return 0, nil, fmt.Errorf("SQL语句执行出错:%s(%v)", s, p)
	}
	return types.GetInt(count), rows, nil
}
