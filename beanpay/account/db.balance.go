package account

import (
	"fmt"

	"github.com/micro-plat/beanpay/beanpay/const/ecodes"
	"github.com/micro-plat/beanpay/beanpay/const/sql"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
)

//GetBalance 查询帐户金额
func getBalance(db db.IDBExecuter, ident string, group string, eid string) (int, error) {
	input := map[string]interface{}{
		"ident":  ident,
		"groups": group,
		"eid":    eid,
	}
	rows, err := db.Query(sql.GetAccountByeid, input)
	if err != nil {
		return 0, err
	}
	if rows.IsEmpty() {
		return 0, errs.NewError(908, "帐户不存在")
	}
	return rows.Get(0).GetInt("balance"), nil

}

//Change 资金变动
func change(db db.IDBExecuter, accountID int, tradeNo string, extNo string, tradeType int, changeType int, amount float64, memo, ext string) (types.IXMap, error) {
	input := map[string]interface{}{
		"account_id":  accountID,
		"amount":      amount,
		"trade_no":    tradeNo,
		"ext_no":      extNo,
		"change_type": changeType,
		"trade_type":  tradeType,
		"ext":         ext,
		"memo":        memo,
	}
	//修改帐户余额
	row, err := db.Execute(sql.ChangeAmount, input)
	if err != nil {
		return nil, err
	}
	if row == 0 {
		return nil, errs.NewError(ecodes.NotEnough, "帐户余额不足")
	}

	//添加资金变动
	row, err = db.Execute(sql.AddBalanceRecord, input)
	if err != nil {
		return nil, err
	}
	data, err := getRecordByTradeNo(db, accountID, tradeNo, tradeType, changeType)
	if errs.GetCode(err) == ecodes.NotExists {
		return nil, errs.NewError(ecodes.Failed, "添加资金变动记录失败")
	}
	return data, nil
}

//Exists 检查交易是否已存在
func exists(db db.IDBExecuter, accountID int, tradeNo string, tradeType int, changeType int) (bool, error) {

	input := map[string]interface{}{
		"account_id":  accountID,
		"trade_no":    tradeNo,
		"change_type": changeType,
		"trade_type":  tradeType,
	}

	// 锁账户
	accID, err := db.Scalar(sql.LockAccount, input)
	if err != nil || types.GetInt64(accID) == 0 {
		return false, fmt.Errorf("锁账户失败，account_id:%v,accID:%v,err:%v", accountID, accID, err)
	}
	// 检查交易是否已存在
	row, err := db.Scalar(sql.ExistsBalanceRecord, input)
	if err != nil {
		return false, err
	}
	return types.GetInt(row) != 0, nil
}
func getRecordByID(db db.IDBExecuter, id int64) (types.IXMap, error) {
	rows, err := db.Query(sql.GetBalanceRecord, map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return nil, err
	}
	if rows.IsEmpty() {
		return nil, errs.NewError(ecodes.NotExists, "记录不存在")
	}
	return rows.Get(0), nil
}
func getRecordByTradeNo(db db.IDBExecuter, accountID int, tradeNo string, tradeType int, changeType int) (types.IXMap, error) {
	rows, err := db.Query(sql.GetBalanceRecordByTradeNo, map[string]interface{}{
		"account_id":  accountID,
		"trade_no":    tradeNo,
		"change_type": changeType,
		"trade_type":  tradeType,
	})
	if err != nil {
		return nil, err
	}
	if rows.IsEmpty() {
		return nil, errs.NewError(ecodes.NotExists, "记录不存在")
	}
	return rows.Get(0), nil
}

// lockTradeRecord 锁交易记录
func lockTradeRecord(db db.IDBExecuter, accountID int, tradeNo string, tradeType int, changeType int) (float64, error) {
	input := map[string]interface{}{
		"account_id":  accountID,
		"trade_no":    tradeNo,
		"change_type": changeType,
		"trade_type":  tradeType,
	}
	row, err := db.Scalar(sql.LockTradeRecord, input)
	if err != nil {
		return 0, err
	}
	return types.GetFloat64(row), nil
}

// queryTradedAmount 查询已交易金额
func queryTradedAmount(db db.IDBExecuter, accountID int, extNo string, tradeType int, changeType int) (float64, error) {
	input := map[string]interface{}{
		"account_id":  accountID,
		"ext_no":      extNo,
		"change_type": changeType,
		"trade_type":  tradeType,
	}
	row, err := db.Scalar(sql.QueryTradedAmount, input)
	if err != nil {
		return 0, err
	}
	return types.GetFloat64(row), nil
}

// checkRefundAmount 查询已退款金额
func checkRefundAmount(db db.IDBExecuter, accountID int, tradeNo, extNo string, tradeType int, changeType int, deductAmount, amount float64) (bool, error) {
	input := map[string]interface{}{
		"account_id":    accountID,
		"ext_no":        extNo,
		"trade_no":      tradeNo,
		"change_type":   changeType,
		"trade_type":    tradeType,
		"deduct_amount": deductAmount,
		"amount":        amount,
	}
	// 检查交易是否已存在
	count, err := db.Scalar(sql.ExistsBalanceRecord, input)
	if err != nil {
		return false, err
	}

	row, err := db.Query(sql.CheckRefundAmount, input)
	if err != nil || row.IsEmpty() {
		return false, fmt.Errorf("查询已退款金额发生异常,count:%v,err:%v", row.Len(), err)
	}
	if !row.Get(0).GetBool("can_refund") {
		return false, errs.NewErrorf(ecodes.AmountErr, "扣款金额:%v,已退款金额:%v,本次退款金额:%v", deductAmount, row.Get(0).GetFloat64("refund_amount"), amount)
	}
	return types.GetInt(count) != 0, nil
}
