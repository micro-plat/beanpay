package pkg

import (
	"github.com/micro-plat/beanpay/beanpay/account"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"
)

//Create 创建服务包信息
func Create(db db.IDBExecuter, uaid string, spkgid string, name string, total int, daily int, expires string) (int, error) {
	id, err := account.GetAccountID(db, uaid)
	if err != nil {
		return 0, err
	}

	return create(db, id, spkgid, name, total, daily, expires)
}

//GetPackageRemain 查询包剩余数量
func GetPackageRemain(db db.IDBExecuter, uaid string, spkgID string) (int, error) {
	id, err := account.GetAccountID(db, uaid)
	if err != nil {
		return 0, err
	}
	return getPackageRemain(db, id, spkgID)
}

//AddCapacity 添加服务包数量
func AddCapacity(db db.IDBExecuter, uaid string, spkgid string, tradeNo string, capacity int) error {
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
func DeductCapacity(db db.IDBExecuter, uaid string, spkgid string, tradeNo string, capacity int) error {
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
func RefundCapacity(db db.IDBExecuter, uaid string, spkgid string, tradeNo string, capacity int) error {
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
func Query(db db.IDBExecuter, uaid string, spkgid string, startTime string, pi int, ps int) (db.QueryRows, error) {
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
