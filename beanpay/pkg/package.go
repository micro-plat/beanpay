package pkg

import (
	"fmt"

	"github.com/micro-plat/beanpay/beanpay"
	"github.com/micro-plat/beanpay/beanpay/account"
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"
)

//Create 创建服务包信息
func Create(i interface{}, uaid string, spkgid string, name string, total int, daily int, expires string) (int, error) {
	db, err := getDBExecuter(i)
	if err != nil {
		return 0, err
	}
	id, err := account.GetAccountID(db, uaid)
	if err != nil {
		return 0, err
	}

	return create(db, id, spkgid, name, total, daily, expires)
}

//GetPackageRemain 查询包剩余数量
func GetPackageRemain(i interface{}, uaid string, spkgID string) (int, error) {
	db, err := getDBExecuter(i)
	if err != nil {
		return 0, err
	}
	id, err := account.GetAccountID(db, uaid)
	if err != nil {
		return 0, err
	}
	return getPackageRemain(db, id, spkgID)
}

//AddCapacity 添加服务包数量
func AddCapacity(i interface{}, uaid string, spkgid string, tradeNo string, capacity int) error {
	db, err := getDBExecuter(i)
	if err != nil {
		return err
	}
	if capacity <= 0 {
		return context.NewErrorf(903, "数量错误%d", capacity)
	}
	id, err := account.GetAccountID(db, uaid)
	if err != nil {
		return err
	}
	pkgid, err := getPackageID(db, id, spkgid)
	if err != nil {
		return err
	}
	return change(db, pkgid, tradeNo, 0, capacity)
}

//DeductCapacity 扣减服务包数量
func DeductCapacity(i interface{}, uaid string, spkgid string, tradeNo string, capacity int) error {
	db, err := getDBExecuter(i)
	if err != nil {
		return err
	}
	if capacity <= 0 {
		return context.NewErrorf(903, "数量错误%d", capacity)
	}
	id, err := account.GetAccountID(db, uaid)
	if err != nil {
		return err
	}
	pkgid, err := getPackageID(db, id, spkgid)
	if err != nil {
		return err
	}
	return change(db, pkgid, tradeNo, 1, -capacity)
}

//RefundCapacity 退回服务包数量
func RefundCapacity(i interface{}, uaid string, spkgid string, tradeNo string, capacity int) error {
	db, err := getDBExecuter(i)
	if err != nil {
		return err
	}
	if capacity <= 0 {
		return context.NewErrorf(903, "数量错误%d", capacity)
	}
	id, err := account.GetAccountID(db, uaid)
	if err != nil {
		return err
	}
	pkgid, err := getPackageID(db, id, spkgid)
	if err != nil {
		return err
	}
	return change(db, pkgid, tradeNo, 2, capacity)
}

//Query 查询指定服务变的变动明细
func Query(i interface{}, uaid string, spkgid string, startTime string, pi int, ps int) (db.QueryRows, error) {
	db, err := getDBExecuter(i)
	if err != nil {
		return nil, err
	}
	id, err := account.GetAccountID(db, uaid)
	if err != nil {
		return nil, err
	}
	pkgid, err := getPackageID(db, id, spkgid)
	if err != nil {
		return nil, err
	}
	return query(db, pkgid, startTime, pi, ps)
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
