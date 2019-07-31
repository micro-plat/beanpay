package account

import (
	"github.com/micro-plat/beanpay/beanpay/const/ecodes"
	"github.com/micro-plat/beanpay/beanpay/const/sql"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
)

//GetBalance 查询帐户金额
func getBalance(db db.IDBExecuter, uid string) (int, error) {
	input := map[string]interface{}{
		"uid": uid,
	}
	rows, _, _, err := db.Query(sql.GetAccountByUid, input)
	if err != nil {
		return 0, err
	}
	if rows.IsEmpty() {
		return 0, context.NewError(908, "帐户不存在")
	}
	return rows.Get(0).GetInt("balance"), nil

}

//Change 资金变动
func change(db db.IDBExecuter, accountID int, tradeNo string, tp int, amount int) (types.XMap, error) {
	input := map[string]interface{}{
		"account_id": accountID,
		"amount":     amount,
		"trade_no":   tradeNo,
		"tp":         tp,
	}
	//修改帐户余额
	row, _, _, err := db.Execute(sql.ChangeAmount, input)
	if err != nil {
		return nil, err
	}
	if row == 0 {
		return nil, context.NewError(ecodes.NotEnough, "帐户不存在或余额不足")
	}

	//添加资金变动
	row, _, _, err = db.Execute(sql.AddBalanceRecord, input)
	if err != nil {
		return nil, err
	}
	data, err := getRecordByTradeNo(db, accountID, tradeNo, tp)
	if context.GetCode(err) == ecodes.NotExists {
		return nil, context.NewError(ecodes.Failed, "添加资金变动失败")
	}
	return data, nil
}

//Exists 检查交易是否已存在
func exists(db db.IDBExecuter, accountID int, tradeNo string, maxAmount int, tp int) (bool, error) {
	input := map[string]interface{}{
		"account_id": accountID,
		"trade_no":   tradeNo,
		"max_amount": maxAmount,
		"tp":         tp,
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
func getRecordByTradeNo(db db.IDBExecuter, accountID int, tradeNo string, tp int) (db.QueryRow, error) {
	rows, _, _, err := db.Query(sql.GetBalanceRecordByTradeNo, map[string]interface{}{
		"account_id": accountID,
		"trade_no":   tradeNo,
		"tp":         tp,
	})
	if err != nil {
		return nil, err
	}
	if rows.IsEmpty() {
		return nil, context.NewError(ecodes.NotExists, "记录不存在")
	}
	return rows.Get(0), nil
}
