package account

import (
	"github.com/micro-plat/beanpay/beanpay/const/sql"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
)

//GetBalance 查询帐户金额
func getBalance(db db.IDBExecuter, accountID int) (int, error) {
	input := map[string]interface{}{
		"account_id": accountID,
	}
	//修改帐户余额
	rows, _, _, err := db.Query(sql.GetAccountByUaid, input)
	if err != nil {
		return 0, err
	}
	if rows.IsEmpty() {
		return 0, context.NewError(908, "帐户不存在")
	}
	return rows.Get(0).GetInt("balance"), nil

}

//Change 资金变动
func change(db db.IDBExecuter, accountID int, tradeNo string, tp int, amount int) error {
	input := map[string]interface{}{
		"account_id": accountID,
		"amount":     amount,
		"trade_no":   tradeNo,
		"tp":         tp,
	}
	//修改帐户余额
	row, _, _, err := db.Execute(sql.ChangeAmount, input)
	if err != nil {
		return err
	}
	if row == 0 {
		return context.NewError(901, "帐户不存在或余额不足")
	}

	//添加资金变动
	row, _, _, err = db.Execute(sql.AddBalanceRecord, input)
	if err != nil {
		return err
	}
	if row > 0 {
		return nil
	}

	//查询记录是否已存在
	e, _, _, err := db.Scalar(sql.ExistsBalanceRecord, input)
	if err != nil {
		return err
	}
	if types.GetInt(e) == 0 {
		return context.NewError(902, "帐户操作失败")
	}
	return nil
}

//Exists 检查交易是否已存在
func exists(db db.IDBExecuter, accountID int, tradeNo string) (bool, error) {
	input := map[string]interface{}{
		"account_id": accountID,
		"trade_no":   tradeNo,
	}
	//修改帐户余额
	row, _, _, err := db.Scalar(sql.ExistsBalanceRecord, input)
	if err != nil {
		return false, err
	}
	return types.GetInt(row) != 0, nil
}
