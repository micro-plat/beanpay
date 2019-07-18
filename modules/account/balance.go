package account

import (
	"github.com/micro-plat/beanpay/modules/const/sql"
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/types"
)

type IBalance interface {

	//Change 资金变动
	Change(accountID int, tradeNo string, tp int, amount int) error

	//Exists 检查交易是否已存在
	Exists(tradeNo string) (bool, error)

	//Query 查询帐户全额
	Query(accountID int) (int, error)
}

type Balance struct {
	c component.IContainer
}

func NewBalance(c component.IContainer) *Balance {
	return &Balance{
		c: c,
	}
}

//Query 查询帐户全额
func (b *Balance) Query(accountID int) (int, error) {
	input := map[string]interface{}{
		"account_id": accountID,
	}

	db := b.c.GetRegularDB()
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
func (b *Balance) Change(accountID int, tradeNo string, tp int, amount int) error {
	input := map[string]interface{}{
		"account_id": accountID,
		"amount":     amount,
		"trade_no":   tradeNo,
		"tp":         tp,
	}
	db, err := b.c.GetRegularDB().Begin()
	if err != nil {
		return err
	}

	//修改帐户余额
	row, _, _, err := db.Execute(sql.ChangeAmount, input)
	if err != nil {
		db.Rollback()
		return err
	}
	if row == 0 {
		db.Rollback()
		return context.NewError(901, "帐户不存在或余额不足")
	}

	//添加资金变动
	row, _, _, err = db.Execute(sql.AddBalanceRecord, input)
	if err != nil {
		db.Rollback()
		return err
	}
	if row > 0 {
		db.Commit()
		return nil
	}

	//查询记录是否已存在
	e, _, _, err := db.Scalar(sql.ExistsBalanceRecord, input)
	if err != nil {
		db.Rollback()
		return err
	}
	if types.GetInt(e) == 0 {
		db.Rollback()
		return context.NewError(902, "帐户操作失败")
	}
	db.Commit()
	return nil
}

//Exists 检查交易是否已存在
func (b *Balance) Exists(accountID int, tradeNo string) (bool, error) {
	input := map[string]interface{}{
		"account_id": accountID,
		"trade_no":   tradeNo,
	}
	db := b.c.GetRegularDB()

	//修改帐户余额
	row, _, _, err := db.Scalar(sql.ExistsBalanceRecord, input)
	if err != nil {
		return false, err
	}
	return types.GetInt(row) != 0, nil
}
