package account

import (
	"fmt"
	"github.com/micro-plat/beanpay/beanpay/const/ecodes"
	"github.com/micro-plat/beanpay/beanpay/const/sql"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
)

//GetBalance 查询帐户金额
func getBalance(db db.IDBExecuter, ident string, group string, eid string) (int, error) {
	input := map[string]interface{}{
		"ident":  ident,
		"groups": group,
		"eid":    eid,
	}
	rows, _, _, err := db.Query(sql.GetAccountByeid, input)
	if err != nil {
		return 0, err
	}
	if rows.IsEmpty() {
		return 0, context.NewError(908, "帐户不存在")
	}
	return rows.Get(0).GetInt("balance"), nil

}

//Change 资金变动
func change(db db.IDBExecuter, accountID int, tradeNo string, extNo string, tradeType int, changeType int, amount int, ext string) (types.XMap, error) {
	input := map[string]interface{}{
		"account_id":  accountID,
		"amount":      amount,
		"trade_no":    tradeNo,
		"ext_no":      extNo,
		"change_type": changeType,
		"trade_type":  tradeType,
		"ext":         ext,
	}
	//修改帐户余额
	row, _, _, err := db.Execute(sql.ChangeAmount, input)
	if err != nil {
		return nil, err
	}
	if row == 0 {
		return nil, context.NewError(ecodes.NotEnough, "帐户余额不足")
	}

	//添加资金变动
	row, _, _, err = db.Execute(sql.AddBalanceRecord, input)
	if err != nil {
		return nil, err
	}
	data, err := getRecordByTradeNo(db, accountID, tradeNo, tradeType, changeType)
	if context.GetCode(err) == ecodes.NotExists {
		return nil, context.NewError(ecodes.Failed, "添加资金变动记录失败")
	}
	return data, nil
}

//Exists 检查交易是否已存在
func exists(db db.IDBExecuter, accountID int, tradeNo string, maxAmount int, tradeType int, changeType int) (bool, error) {
	input := map[string]interface{}{
		"account_id":  accountID,
		"trade_no":    tradeNo,
		"max_amount":  maxAmount,
		"change_type": changeType,
		"trade_type":  tradeType,
	}
	//修改帐户余额
	row, _, _, err := db.Scalar(sql.ExistsBalanceRecord, input)
	if err != nil {
		return false, err
	}
	return types.GetInt(row) != 0, nil
}
func getRecordByID(db db.IDBExecuter, id int64) (db.QueryRow, error) {
	rows, _, _, err := db.Query(sql.GetBalanceRecord, map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return nil, err
	}
	if rows.IsEmpty() {
		return nil, context.NewError(ecodes.NotExists, "记录不存在")
	}
	return rows.Get(0), nil
}
func getRecordByTradeNo(db db.IDBExecuter, accountID int, tradeNo string, tradeType int, changeType int) (db.QueryRow, error) {
	rows, _, _, err := db.Query(sql.GetBalanceRecordByTradeNo, map[string]interface{}{
		"account_id":  accountID,
		"trade_no":    tradeNo,
		"change_type": changeType,
		"trade_type":  tradeType,
	})
	if err != nil {
		return nil, err
	}
	if rows.IsEmpty() {
		return nil, context.NewError(ecodes.NotExists, "记录不存在")
	}
	return rows.Get(0), nil
}

// lockTradeRecord 锁交易记录
func lockTradeRecord(db db.IDBExecuter, accountID int, tradeNo string, tradeType int, changeType int) (int, error) {
	input := map[string]interface{}{
		"account_id":  accountID,
		"trade_no":    tradeNo,
		"change_type": changeType,
		"trade_type":  tradeType,
	}
	row, _, _, err := db.Scalar(sql.LockTradeRecord, input)
	if err != nil {
		return 0, err
	}
	return types.GetInt(row), nil
}

// queryTradedAmount 查询已交易金额
func queryTradedAmount(db db.IDBExecuter, accountID int, extNo string, tradeType int, changeType int) (int, error) {
	input := map[string]interface{}{
		"account_id":  accountID,
		"ext_no":      extNo,
		"change_type": changeType,
		"trade_type":  tradeType,
	}
	row, sqlStr, args, err := db.Scalar(sql.QueryTradedAmount, input)
	if err != nil {
		return 0, err
	}
	fmt.Printf("row:%v,sqlStr:%v, args:%+v\n", row, sqlStr, args)
	return types.GetInt(row), nil
}
