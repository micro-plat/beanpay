package pkg

import (
	"github.com/micro-plat/beanpay/beanpay/account"
	"github.com/micro-plat/beanpay/beanpay/const/ttypes"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"
)

//Create 创建服务包信息
func Create(db db.IDBExecuter, eid string, sid string, name string, total int, daily int, expires string) (int, error) {
	acc, err := account.GetAccount(db, eid)
	if err != nil {
		return 0, err
	}

	return create(db, acc.ID, sid, name, total, daily, expires)
}

//GetPackageRemain 查询包剩余数量
func GetPackageRemain(db db.IDBExecuter, eid string, sid string) (int, error) {
	acc, err := account.GetAccount(db, eid)
	if err != nil {
		return 0, err
	}
	return getPackageRemain(db, acc.ID, sid)
}

//GetPackageID 获取服务包编号
func GetPackageID(db db.IDBExecuter, eid string, sid string) (int, error) {
	acc, err := account.GetAccount(db, eid)
	if err != nil {
		return 0, err
	}
	pkgid, err := getPackageID(db, acc.ID, sid)
	if err != nil {
		return 0, err
	}
	return pkgid, nil
}

//AddCapacity 添加服务包数量
func AddCapacity(db db.IDBExecuter, eid string, sid string, tradeNo string, capacity int) error {
	if capacity <= 0 {
		return context.NewErrorf(903, "数量错误%d", capacity)
	}
	acc, err := account.GetAccount(db, eid)
	if err != nil {
		return err
	}
	pkgid, err := getPackageID(db, acc.ID, sid)
	if err != nil {
		return err
	}
	return change(db, pkgid, tradeNo, ttypes.Add, capacity)
}

//DeductCapacity 扣减服务包数量
func DeductCapacity(db db.IDBExecuter, eid string, sid string, tradeNo string, capacity int) error {

	if capacity <= 0 {
		return context.NewErrorf(903, "数量错误%d", capacity)
	}
	acc, err := account.GetAccount(db, eid)
	if err != nil {
		return err
	}
	pkgid, err := getPackageID(db, acc.ID, sid)
	if err != nil {
		return err
	}
	return change(db, pkgid, tradeNo, ttypes.Deduct, -capacity)
}

//RefundCapacity 退回服务包数量
func RefundCapacity(db db.IDBExecuter, eid string, sid string, tradeNo string, capacity int) error {
	if capacity <= 0 {
		return context.NewErrorf(903, "数量错误%d", capacity)
	}
	acc, err := account.GetAccount(db, eid)
	if err != nil {
		return err
	}
	pkgid, err := getPackageID(db, acc.ID, sid)
	if err != nil {
		return err
	}
	return change(db, pkgid, tradeNo, ttypes.Refund, capacity)
}

//Query 查询指定服务变的变动明细
func Query(db db.IDBExecuter, eid string, sid string, startTime string, endTime string, pi int, ps int) (db.QueryRows, error) {
	acc, err := account.GetAccount(db, eid)
	if err != nil {
		return nil, err
	}
	pkgid := 0
	if sid != "" {
		pkgid, err = getPackageID(db, acc.ID, sid)
		if err != nil {
			return nil, err
		}
	}
	return query(db, acc.ID, pkgid, startTime, endTime, pi, ps)
}
