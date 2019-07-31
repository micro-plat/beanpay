package beanpay

import (
	"fmt"

	"github.com/micro-plat/hydra/context"

	"github.com/micro-plat/beanpay/beanpay/account"
	"github.com/micro-plat/beanpay/beanpay/const/confs"
	"github.com/micro-plat/beanpay/beanpay/pkg"
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
)

//CreateAccount 根据外部用户编号，名称创建资金帐户信息
func CreateAccount(i interface{}, eid string, name string) (interface{}, error) {
	db, err := getDBExecuter(i)
	if err != nil {
		return nil, err
	}
	return account.Create(db, eid, name)
}

//GetAccount 根据eid获取资金帐户编号
func GetAccount(i interface{}, eid string) (*account.Account, error) {
	db, err := getDBExecuter(i)
	if err != nil {
		return nil, err
	}
	return account.GetAccount(db, eid)
}

//AddAmount 指定用户编号，交易变号，金额进行资金帐户加款
func AddAmount(i interface{}, eid string, tradeNo string, amount int) (*context.Result, error) {
	m, db, err := getTrans(i)
	if err != nil {
		return nil, err
	}
	row, err := account.AddAmount(db, eid, tradeNo, amount)
	if !m {
		return row, err
	}
	if err != nil {
		db.Rollback()
		return nil, err
	}
	db.Commit()
	return row, nil
}

//DeductAmount 指定用户编号，交易变号，金额进行资金帐户扣款
func DeductAmount(i interface{}, eid string, tradeNo string, amount int) (*context.Result, error) {
	m, db, err := getTrans(i)
	if err != nil {
		return nil, err
	}
	row, err := account.DeductAmount(db, eid, tradeNo, amount)
	if !m {
		return row, err
	}
	if err != nil {
		db.Rollback()
		return nil, err
	}
	db.Commit()
	return row, nil
}

//RefundAmount 指定用户编号，交易变号，金额进行资金帐户退款
func RefundAmount(i interface{}, eid string, tradeNo string, amount int) (*context.Result, error) {
	m, db, err := getTrans(i)
	if err != nil {
		return nil, err
	}
	row, err := account.RefundAmount(db, eid, tradeNo, amount)
	if !m {
		return row, err
	}
	if err != nil {
		db.Rollback()
		return nil, err
	}
	db.Commit()
	return row, nil
}

//QueryAccountRecords 查询指定用户在一段时间内的资金变动信息
func QueryAccountRecords(i interface{}, eid string, startTime string, endTime string, pi int, ps int) (db.QueryRows, error) {
	db, err := getDBExecuter(i)
	if err != nil {
		return nil, err
	}
	return account.Query(db, eid, startTime, endTime, pi, ps)
}

//CreatePackage 根据用户编号， 服务编号，服务名称，服务包可用总数，日限制使用次数，过期时间创建服务包
//用户必须先创建资金帐户
func CreatePackage(i interface{}, eid string, sid string, name string, total int, daily int, expires string) (interface{}, error) {
	db, err := getDBExecuter(i)
	if err != nil {
		return 0, err
	}
	return pkg.Create(db, eid, sid, name, total, daily, expires)
}

//GetPackage 根据用户编号，服务编号获取服务包编号
func GetPackage(i interface{}, eid string, sid string) (*pkg.PKG, error) {
	db, err := getDBExecuter(i)
	if err != nil {
		return nil, err
	}
	return pkg.GetPackage(db, eid, sid)
}

//AddCapacity 指定用户编号，交易变号，金额进行服务包数量追加
func AddCapacity(i interface{}, eid string, sid string, tradeNo string, capacity int) (*context.Result, error) {
	m, db, err := getTrans(i)
	if err != nil {
		return nil, err
	}
	row, err := pkg.AddCapacity(db, eid, sid, tradeNo, capacity)
	if !m {
		return row, err
	}
	if err != nil {
		db.Rollback()
		return nil, err
	}
	db.Commit()
	return row, nil
}

//DeductCapacity 指定用户编号，交易变号，金额进行服务包数量扣减
func DeductCapacity(i interface{}, eid string, sid string, tradeNo string, capacity int) (*context.Result, error) {
	m, db, err := getTrans(i)
	if err != nil {
		return nil, err
	}
	row, err := pkg.DeductCapacity(db, eid, sid, tradeNo, capacity)
	if !m {
		return row, err
	}
	if err != nil {
		db.Rollback()
		return nil, err
	}
	db.Commit()
	return row, nil

}

//RefundCapacity 指定用户编号，交易变号，金额进行服务包数量退回
func RefundCapacity(i interface{}, eid string, sid string, tradeNo string, capacity int) (*context.Result, error) {
	m, db, err := getTrans(i)
	if err != nil {
		return nil, err
	}
	row, err := pkg.RefundCapacity(db, eid, sid, tradeNo, capacity)
	if !m {
		return row, err
	}
	if err != nil {
		db.Rollback()
		return nil, err
	}
	db.Commit()
	return row, nil
}

//QueryPackageRecords 查询指定用户在一段时间内的服务包的变动记录
func QueryPackageRecords(i interface{}, eid string, sid string, startTime string, endTime string, pi int, ps int) (db.QueryRows, error) {
	db, err := getDBExecuter(i)
	if err != nil {
		return nil, err
	}
	return pkg.Query(db, eid, sid, startTime, endTime, pi, ps)
}

func getTrans(c interface{}) (bool, db.IDBTrans, error) {
	b, e, err := getDB(c)
	if err != nil {
		return false, nil, err
	}
	if b {
		return false, e.(db.IDBTrans), nil
	}
	t, err := e.(db.IDB).Begin()
	if err != nil {
		return false, nil, err
	}
	return true, t, nil
}

func getDBExecuter(c interface{}) (db.IDBExecuter, error) {
	_, e, err := getDB(c)
	return e, err
}

func getDB(c interface{}) (bool, db.IDBExecuter, error) {
	switch v := c.(type) {
	case *context.Context:
		db, err := v.GetContainer().GetDB(confs.DBName)
		return false, db, err
	case component.IContainer:
		db, err := v.GetDB(confs.DBName)
		return false, db, err
	case db.IDB:
		return false, v, nil
	case db.IDBTrans:
		return true, v, nil
	default:
		return false, nil, fmt.Errorf("不支持的参数类型")
	}
}
