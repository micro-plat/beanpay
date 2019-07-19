package account

import (
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"
)

//Create 根据uaid,name创建帐户,如果帐户存在直接返回帐户编号
func Create(db db.IDBExecuter, uaid string, name string) (int, error) {
	if id, err := getAccountID(db, uaid); id != 0 && err != nil {
		return id, nil
	}
	return create(db, uaid, name)
}

//GetBalance 获取帐户余额
func GetBalance(db db.IDBExecuter, uaid string) (int, error) {
	id, err := getAccountID(db, uaid)
	if err != nil {
		return 0, err
	}
	return getBalance(db, id)
}

//GetAccountID 根据uaid获取帐户编号
func GetAccountID(db db.IDBExecuter, uaid string) (int, error) {
	return getAccountID(db, uaid)
}

//AddAmount 资金加款
func AddAmount(db db.IDBExecuter, uaid string, tradeNo string, amount int) error {
	if amount <= 0 {
		return context.NewErrorf(903, "金额错误%d", amount)
	}
	id, err := getAccountID(db, uaid)
	if err != nil {
		return err
	}
	b, err := exists(db, id, tradeNo)
	if err != nil {
		return err
	}
	if b {
		return nil
	}
	return change(db, id, tradeNo, 0, amount)
}

//DeductAmount 资金扣款
func DeductAmount(db db.IDBExecuter, uaid string, tradeNo string, amount int) error {
	if amount <= 0 {
		return context.NewErrorf(903, "金额错误%d", amount)
	}
	id, err := getAccountID(db, uaid)
	if err != nil {
		return err
	}
	b, err := exists(db, id, tradeNo)
	if err != nil {
		return err
	}
	if b {
		return nil
	}
	return change(db, id, tradeNo, 1, -amount)
}

//RefundAmount 资金退款
func RefundAmount(db db.IDBExecuter, uaid string, tradeNo string, amount int) error {
	if amount <= 0 {
		return context.NewErrorf(903, "金额错误%d", amount)
	}
	id, err := getAccountID(db, uaid)
	if err != nil {
		return err
	}
	b, err := exists(db, id, tradeNo)
	if err != nil {
		return err
	}
	if b {
		return nil
	}
	return change(db, id, tradeNo, 2, amount)
}

//Query 查询余额变动明细
func Query(db db.IDBExecuter, uaid string, startTime string, pi int, ps int) (db.QueryRows, error) {
	id, err := getAccountID(db, uaid)
	if err != nil {
		return nil, err
	}
	return query(db, id, startTime, pi, ps)
}
