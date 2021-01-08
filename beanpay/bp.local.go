package beanpay

import (
	"github.com/micro-plat/beanpay/beanpay/account"
	"github.com/micro-plat/beanpay/beanpay/const/confs"
	"github.com/micro-plat/beanpay/beanpay/const/ecodes"
	"github.com/micro-plat/beanpay/beanpay/const/ttypes"
	"github.com/micro-plat/beanpay/beanpay/pkg"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
)

var _ IBeanpay = &Beanpay{}

//Beanpay 支付对象
type Beanpay struct {
	ident string
	group string
}

//newBeanpay 构建支付对象,传入外部系统标识，帐户类型
func newBeanpay(ident string, group ...string) *Beanpay {
	return &Beanpay{
		ident: ident,
		group: types.GetStringByIndex(group, 0, "-"),
	}
}

//CreateAccount 根据外部用户编号，名称创建资金帐户信息
func (b *Beanpay) CreateAccount(i interface{}, eid string, name string) (*account.AccountResult, error) {
	db, err := getDBExecuter(i)
	if err != nil {
		return nil, err
	}
	return account.Create(db, b.ident, b.group, eid, name)
}

//SetAccountName 根据外部用户编号，名称修改账户名称
func (b *Beanpay) SetAccountName(i interface{}, eid string, name string) (*account.AccountResult, error) {
	db, err := getDBExecuter(i)
	if err != nil {
		return nil, err
	}
	return account.SetAccountName(db, b.ident, b.group, eid, name)
}

//GetAccount 根据eid获取资金帐户编号
func (b *Beanpay) GetAccount(i interface{}, eid string) (*account.Account, error) {
	db, err := getDBExecuter(i)
	if err != nil {
		return nil, err
	}
	return account.GetAccount(db, b.ident, b.group, eid)
}

//QueryAccount 查询账户列表
func (b *Beanpay) QueryAccount(i interface{}, eid, accountType, name, status string, pi, ps int) (r *account.AccountInfoList, err error) {
	db, err := getDBExecuter(i)
	if err != nil {
		return nil, err
	}
	return account.QueryAccount(db, b.ident, b.group, eid, accountType, name, status, types.GetMax(pi, 1), ps)
}

// SetCreditAmount 设置授信金额
func (b *Beanpay) SetCreditAmount(i interface{}, eid string, credit float64) (*account.AccountResult, error) {
	db, err := getDBExecuter(i)
	if err != nil {
		return nil, err
	}
	return account.SetCreditAmount(db, b.ident, b.group, eid, credit)
}

//AddAmount 指定用户编号，交易变号，金额进行资金帐户加款,memo 第一个参数为交易描述, ext拓展信息
func (b *Beanpay) AddAmount(i interface{}, eid string, tradeNo string, amount float64, memo string, ext ...string) (*account.RecordResult, error) {
	m, db, err := getTrans(i)
	if err != nil {
		return nil, err
	}

	if amount == 0 {
		return nil, errs.NewErrorf(ecodes.AmountErr, "金额错误%v", amount)
	}
	row, err := account.AddAmount(db, b.ident, b.group, eid, tradeNo, TPTrade, ttypes.Add, amount, memo, types.GetStringByIndex(ext, 0, "{}"))
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
func (b *Beanpay) DrawingAmount(i interface{}, eid string, tradeNo string, amount float64, memo string, ext ...string) (*account.RecordResult, error) {
	m, db, err := getTrans(i)
	if err != nil {
		return nil, err
	}
	if amount == 0 {
		return nil, errs.NewErrorf(ecodes.AmountErr, "金额错误%v", amount)
	}
	row, err := account.DrawingAmount(db, b.ident, b.group, eid, tradeNo, TPTrade, ttypes.Drawing, amount, memo, types.GetStringByIndex(ext, 0, "{}"))
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
func (b *Beanpay) DeductAmount(i interface{}, eid string, tradeNo string, tradeType TradeType, amount float64, memo string, ext ...string) (*account.RecordResult, error) {
	m, db, err := getTrans(i)
	if err != nil {
		return nil, err
	}
	row, err := account.DeductAmount(db, b.ident, b.group, eid, tradeNo, int(tradeType), amount, memo, types.GetStringByIndex(ext, 0, "{}"))
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
func (b *Beanpay) RefundAmount(i interface{}, eid string, tradeNo string, extNo string, tradeType TradeType, amount float64, memo string, ext ...string) (*account.RecordResult, error) {

	m, db, err := getTrans(i)
	if err != nil {
		return nil, err
	}
	row, err := account.RefundAmount(db, b.ident, b.group, eid, tradeNo, extNo, int(tradeType), amount, memo, types.GetStringByIndex(ext, 0, "{}"))
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

//TradeFlatAmount 指定用户编号，交易遍号,交易类型,变动类型(只能时交易平账和余额平账)，金额进行资金帐户交易平账
func (b *Beanpay) TradeFlatAmount(i interface{}, eid string, tradeNo string, tradeType TradeType, amount float64, memo string, ext ...string) (*account.RecordResult, error) {
	m, db, err := getTrans(i)
	if err != nil {
		return nil, err
	}

	var row *account.RecordResult
	row, err = account.AddAmount(db, b.ident, b.group, eid, tradeNo, int(tradeType), ttypes.TradeFlat, -amount, memo, types.GetStringByIndex(ext, 0, "{}"))
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

//BalanceFlatAmount 指定用户编号，交易遍号,交易类型,变动类型(只能时交易平账和余额平账)，金额进行资金帐户余额平账
func (b *Beanpay) BalanceFlatAmount(i interface{}, eid string, tradeNo string, tradeType TradeType, amount float64, memo string, ext ...string) (*account.RecordResult, error) {
	m, db, err := getTrans(i)
	if err != nil {
		return nil, err
	}

	var row *account.RecordResult
	row, err = account.AddAmount(db, b.ident, b.group, eid, tradeNo, int(tradeType), ttypes.BalanceFlat, amount, memo, types.GetStringByIndex(ext, 0, "{}"))
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

// ReverseAddAmount 红冲加款
func (b *Beanpay) ReverseAddAmount(i interface{}, eid string, tradeNo string, extNo string, memo string, ext ...string) (*account.RecordResult, error) {
	m, db, err := getTrans(i)
	if err != nil {
		return nil, err
	}

	row, err := account.ReverseAmount(db, b.ident, b.group, eid, tradeNo, extNo, int(ttypes.Account), ttypes.Add, memo, types.GetStringByIndex(ext, 0, "{}"))
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

// ReverseDrawingAmount 红冲提款
func (b *Beanpay) ReverseDrawingAmount(i interface{}, eid string, tradeNo string, extNo string, memo string, ext ...string) (*account.RecordResult, error) {
	m, db, err := getTrans(i)
	if err != nil {
		return nil, err
	}

	row, err := account.ReverseAmount(db, b.ident, b.group, eid, tradeNo, extNo, int(ttypes.Account), ttypes.Drawing, memo, types.GetStringByIndex(ext, 0, "{}"))
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
func (b *Beanpay) QueryAccountRecords(i interface{}, accountType string, accountID string, accountName string, group string, changeType string, tradeType string, eid string, startTime string, endTime string, pi int, ps int) (*account.RecordResults, error) {
	db, err := getDBExecuter(i)
	if err != nil {
		return nil, err
	}
	return account.Query(db, accountType, accountID, accountName, group, changeType, tradeType, eid, startTime, endTime, types.GetMax(pi, 1), ps)
}

//CreatePackage 根据用户编号， 服务编号，服务名称，服务包可用总数，日限制使用次数，过期时间创建服务包
//用户必须先创建资金帐户
func (b *Beanpay) CreatePackage(i interface{}, eid string, spid string, name string, total int, daily int, expires string) (interface{}, error) {
	db, err := getDBExecuter(i)
	if err != nil {
		return 0, err
	}
	return pkg.Create(db, b.ident, b.group, eid, spid, name, total, daily, expires)
}

//GetPackage 根据用户编号，服务编号获取服务包编号
func (b *Beanpay) GetPackage(i interface{}, eid string, spid string) (*pkg.PKG, error) {
	db, err := getDBExecuter(i)
	if err != nil {
		return nil, err
	}
	return pkg.GetPackage(db, b.ident, b.group, eid, spid)
}

//AddCapacity 指定用户编号，交易变号，金额进行服务包数量追加
func (b *Beanpay) AddCapacity(i interface{}, eid string, spid string, tradeNo string, capacity int, ext ...string) (*errs.Result, error) {
	m, db, err := getTrans(i)
	if err != nil {
		return nil, err
	}
	row, err := pkg.AddCapacity(db, b.ident, b.group, eid, spid, tradeNo, capacity, types.GetStringByIndex(ext, 0, "{}"))
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
func (b *Beanpay) DrawingCapacity(i interface{}, eid string, spid string, tradeNo string, capacity int, ext ...string) (*errs.Result, error) {
	m, db, err := getTrans(i)
	if err != nil {
		return nil, err
	}
	row, err := pkg.DrawingCapacity(db, b.ident, b.group, eid, spid, tradeNo, capacity, types.GetStringByIndex(ext, 0, "{}"))
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
func (b *Beanpay) DeductCapacity(i interface{}, eid string, spid string, tradeNo string, capacity int, ext ...string) (*errs.Result, error) {
	m, db, err := getTrans(i)
	if err != nil {
		return nil, err
	}
	row, err := pkg.DeductCapacity(db, b.ident, b.group, eid, spid, tradeNo, capacity, types.GetStringByIndex(ext, 0, "{}"))
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
func (b *Beanpay) RefundCapacity(i interface{}, eid string, spid string, tradeNo string, capacity int, ext ...string) (*errs.Result, error) {
	m, db, err := getTrans(i)
	if err != nil {
		return nil, err
	}
	row, err := pkg.RefundCapacity(db, b.ident, b.group, eid, spid, tradeNo, capacity, types.GetStringByIndex(ext, 0, "{}"))
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
	return pkg.Query(db, b.ident, b.group, eid, spid, startTime, endTime, types.GetMax(pi, 1), ps)
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
	case db.IDB:
		return false, v, nil
	case db.IDBTrans:
		return true, v, nil
	default:
		c, err := hydra.C.DB().GetDB(confs.DBName)
		return false, c, err
	}
}
