package account

import (
	"fmt"

	"github.com/micro-plat/beanpay/beanpay"
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"
)

//Create 根据uaid,name创建帐户,如果帐户存在直接返回帐户编号
func Create(i interface{}, uaid string, name string) (int, error) {
	db, err := getDBExecuter(i)
	if err != nil {
		return 0, err
	}
	if id, err := getAccountID(db, uaid); id != 0 && err != nil {
		return id, nil
	}
	return create(db, uaid, name)
}

//GetBalance 获取帐户余额
func GetBalance(i interface{}, uaid string) (int, error) {
	db, err := getDBExecuter(i)
	if err != nil {
		return 0, err
	}
	id, err := getAccountID(db, uaid)
	if err != nil {
		return 0, err
	}
	return getBalance(db, id)
}

//GetAccountID 根据uaid获取帐户编号
func GetAccountID(i interface{}, uaid string) (int, error) {
	db, err := getDBExecuter(i)
	if err != nil {
		return 0, err
	}
	return getAccountID(db, uaid)
}

//AddAmount 资金加款
func AddAmount(i interface{}, uaid string, tradeNo string, amount int) error {
	db, err := getDBExecuter(i)
	if err != nil {
		return err
	}
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
func DeductAmount(i interface{}, uaid string, tradeNo string, amount int) error {
	db, err := getDBExecuter(i)
	if err != nil {
		return err
	}
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
func RefundAmount(i interface{}, uaid string, tradeNo string, amount int) error {
	db, err := getDBExecuter(i)
	if err != nil {
		return err
	}
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
func Query(i interface{}, uaid string, startTime string, pi int, ps int) (db.QueryRows, error) {
	db, err := getDBExecuter(i)
	if err != nil {
		return nil, err
	}
	id, err := getAccountID(db, uaid)
	if err != nil {
		return nil, err
	}
	return query(db, id, startTime, pi, ps)
}
func getDBExecuter(c interface{}) (db.IDBExecuter, error) {
	switch v := c.(type) {
	case *context.Context:
		return v.GetContainer().GetDB(beanpay.DBName)
	case component.IContainer:
		return v.GetDB(beanpay.DBName)
	case db.IDBExecuter:
		return v, nil
	default:
		return nil, fmt.Errorf("不支持的参数类型")
	}
}
