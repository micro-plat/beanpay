package account

import (
	"fmt"

	"github.com/micro-plat/beanpay/beanpay/const/sql"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
)

//Query 查询余额变动明细
func query(db db.IDBExecuter, accountType string, group string, accountID string, accountName string, changeType string, tradeType string, startTime string, endTime string, pi int, ps int) (int, db.QueryRows, error) {
	input := map[string]interface{}{
		"types":        accountType,
		"account_id":   accountID,
		"start":        startTime,
		"end":          endTime,
		"change_type":  changeType,
		"groupx":       group,
		"account_name": accountName,
		"trade_type":   tradeType,
		"currentPage":  (pi - 1) * ps,
		"size":         pi * ps,
		"pageSize":     ps,
	}
	count, err := db.Scalar(sql.QueryBalanceRecordCount, input)
	if err != nil {
		return 0, nil, fmt.Errorf("SQL语句执行出错:%w", err)
	}

	rows, err := db.Query(sql.QueryBalanceRecord, input)
	if err != nil {
		return 0, nil, fmt.Errorf("SQL语句执行出错:%w", err)
	}
	return types.GetInt(count), rows, nil
}
