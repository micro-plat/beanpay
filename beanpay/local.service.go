package beanpay

import (
	"fmt"
	"strings"

	"github.com/micro-plat/hydra/context"

	"github.com/micro-plat/beanpay/beanpay/account"
	"github.com/micro-plat/beanpay/beanpay/const/confs"
	"github.com/micro-plat/beanpay/beanpay/const/ttypes"
	"github.com/micro-plat/beanpay/beanpay/pkg"
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
)

//Beanpay 支付对象
type Beanpay struct {
	ident string
	tp    string
}

//NewBeanpay 构建支付对象,传入外部系统标识，帐户类型
func NewBeanpay(ident string, accountType ...string) *Beanpay {
	return &Beanpay{
		ident: ident,
		tp:    types.GetStringByIndex(accountType, 0),
	}
}

//NewBeanpayByInternal 构建支付对象，内部系统构建时使用
func NewBeanpayByInternal(accountType ...string) *Beanpay {
	return &Beanpay{
		tp: types.GetStringByIndex(accountType, 0),
	}
}
func (b *Beanpay) makeEID(eid string) string {
	var buff strings.Builder
	if b.ident != "" {
		buff.WriteString(b.ident)
		buff.WriteString(":")
	}
	if b.tp != "" {
		buff.WriteString(b.tp)
		buff.WriteString(":")
	}
	buff.WriteString(eid)
	return buff.String()
}

//CreateAccount 根据外部用户编号，名称创建资金帐户信息
func (b *Beanpay) CreateAccount(i interface{}, eid string, name string) (*account.AccountResult, error) {
	db, err := getDBExecuter(i)
	if err != nil {
		return nil, err
	}
	return account.Create(db, b.makeEID(eid), b.tp, name)
}

//GetAccount 根据eid获取资金帐户编号
func (b *Beanpay) GetAccount(i interface{}, eid string) (*account.Account, error) {
	db, err := getDBExecuter(i)
	if err != nil {
		return nil, err
	}
	return account.GetAccount(db, b.makeEID(eid))
}

//AddAmount 指定用户编号，交易变号，金额进行资金帐户加款
func (b *Beanpay) AddAmount(i interface{}, eid string, tradeNo string, amount int) (*account.RecordResult, error) {
	m, db, err := getTrans(i)
	if err != nil {
		return nil, err
	}
	row, err := account.AddAmount(db, b.makeEID(eid), tradeNo, amount)
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

//DrawingAmount 指定用户编号，交易变号，金额进行资金帐户提款
func (b *Beanpay) DrawingAmount(i interface{}, eid string, tradeNo string, amount int) (*account.RecordResult, error) {
	m, db, err := getTrans(i)
	if err != nil {
		return nil, err
	}
	row, err := account.DrawingAmount(db, b.makeEID(eid), tradeNo, amount)
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
func (b *Beanpay) DeductAmount(i interface{}, eid string, tradeNo string, amount int, tradeTypes ...int) (*account.RecordResult, error) {
	tradeType := ttypes.OrderTrade
	if len(tradeTypes) > 0 {
		tradeType = tradeTypes[0]
	}
	m, db, err := getTrans(i)
	if err != nil {
		return nil, err
	}
	row, err := account.DeductAmount(db, b.makeEID(eid), tradeNo, amount, tradeType)
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
func (b *Beanpay) RefundAmount(i interface{}, eid string, tradeNo string, reductNo string, amount int, tradeTypes ...int) (*account.RecordResult, error) {
	tradeType := ttypes.OrderTrade
	if len(tradeTypes) > 0 {
		tradeType = tradeTypes[0]
	}
	m, db, err := getTrans(i)
	if err != nil {
		return nil, err
	}
	row, err := account.RefundAmount(db, b.makeEID(eid), tradeNo, reductNo, amount, tradeType)
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
func (b *Beanpay) QueryAccountRecords(i interface{}, eid string, startTime string, endTime string, pi int, ps int) (*account.RecordResults, error) {
	db, err := getDBExecuter(i)
	if err != nil {
		return nil, err
	}
	return account.Query(db, b.makeEID(eid), startTime, endTime, pi, ps)
}

//CreatePackage 根据用户编号， 服务编号，服务名称，服务包可用总数，日限制使用次数，过期时间创建服务包
//用户必须先创建资金帐户
func (b *Beanpay) CreatePackage(i interface{}, eid string, spid string, name string, total int, daily int, expires string) (interface{}, error) {
	db, err := getDBExecuter(i)
	if err != nil {
		return 0, err
	}
	return pkg.Create(db, b.makeEID(eid), spid, name, total, daily, expires)
}

//GetPackage 根据用户编号，服务编号获取服务包编号
func (b *Beanpay) GetPackage(i interface{}, eid string, spid string) (*pkg.PKG, error) {
	db, err := getDBExecuter(i)
	if err != nil {
		return nil, err
	}
	return pkg.GetPackage(db, b.makeEID(eid), spid)
}

//AddCapacity 指定用户编号，交易变号，金额进行服务包数量追加
func (b *Beanpay) AddCapacity(i interface{}, eid string, spid string, tradeNo string, capacity int) (*context.Result, error) {
	m, db, err := getTrans(i)
	if err != nil {
		return nil, err
	}
	row, err := pkg.AddCapacity(db, b.makeEID(eid), spid, tradeNo, capacity)
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

//DrawingCapacity 指定用户编号，交易变号，金额进行服务包数量提取
func (b *Beanpay) DrawingCapacity(i interface{}, eid string, spid string, tradeNo string, capacity int) (*context.Result, error) {
	m, db, err := getTrans(i)
	if err != nil {
		return nil, err
	}
	row, err := pkg.DrawingCapacity(db, b.makeEID(eid), spid, tradeNo, capacity)
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
func (b *Beanpay) DeductCapacity(i interface{}, eid string, spid string, tradeNo string, capacity int) (*context.Result, error) {
	m, db, err := getTrans(i)
	if err != nil {
		return nil, err
	}
	row, err := pkg.DeductCapacity(db, b.makeEID(eid), spid, tradeNo, capacity)
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
func (b *Beanpay) RefundCapacity(i interface{}, eid string, spid string, tradeNo string, capacity int) (*context.Result, error) {
	m, db, err := getTrans(i)
	if err != nil {
		return nil, err
	}
	row, err := pkg.RefundCapacity(db, b.makeEID(eid), spid, tradeNo, capacity)
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
func (b *Beanpay) QueryPackageRecords(i interface{}, eid string, spid string, startTime string, endTime string, pi int, ps int) (db.QueryRows, error) {
	db, err := getDBExecuter(i)
	if err != nil {
		return nil, err
	}
	return pkg.Query(db, b.makeEID(eid), spid, startTime, endTime, pi, ps)
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
