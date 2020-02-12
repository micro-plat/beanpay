package account

import (
	"github.com/micro-plat/beanpay/beanpay/const/ecodes"
	"github.com/micro-plat/beanpay/beanpay/const/ttypes"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"
)

//Create 根据eid,name创建帐户,如果帐户存在直接返回帐户编号
func Create(db db.IDBExecuter, ident string, group string, eid string, name string) (*AccountResult, error) {
	acc, err := GetAccount(db, ident, group, eid)
	if err == nil {
		return NewAccountResult(ecodes.HasExists, acc), nil
	}
	if context.GetCode(err) != ecodes.NotExists {
		return nil, err
	}
	if err = create(db, ident, group, eid, name); err != nil {
		return nil, err
	}
	acc, err = GetAccount(db, ident, group, eid)
	if err != nil {
		return nil, err
	}
	return NewAccountResult(ecodes.Success, acc), nil
}

//SetCreditAmount 设置授信金额
func SetCreditAmount(db db.IDBExecuter, ident string, group string, eid string, credit int) (*AccountResult, error) {
	acc, err := GetAccount(db, ident, group, eid)
	if err != nil {
		return nil, err
	}
	if err = setCreditAmount(db, credit, acc.ID); err != nil {
		return nil, err
	}
	acc, err = GetAccount(db, ident, group, eid)
	if err != nil {
		return nil, err
	}
	return NewAccountResult(ecodes.Success, acc), nil
}

//GetBalance 获取帐户余额
func GetBalance(db db.IDBExecuter, ident string, group string, eid string) (int, error) {
	return getBalance(db, ident, group, eid)
}

//GetAccount 根据eid获取帐户
func GetAccount(db db.IDBExecuter, ident string, group string, eid string) (acc *Account, err error) {
	row, err := getAccount(db, ident, group, eid)
	if err != nil {
		return nil, err
	}
	acc = &Account{}
	if err = row.ToStruct(acc); err != nil {
		return nil, err
	}
	return acc, nil
}

//AddAmount 资金加款
func AddAmount(db db.IDBExecuter, ident string, group string, eid string, tradeNo string, tradeType int, changeType int, amount int, ext string) (*RecordResult, error) {

	acc, err := GetAccount(db, ident, group, eid)
	if err != nil {
		return nil, err
	}

	b, err := exists(db, acc.ID, tradeNo, 0, tradeType, changeType)
	if err != nil {
		return nil, err
	}
	if b {
		row, err := getRecordByTradeNo(db, acc.ID, tradeNo, tradeType, changeType)
		if err != nil {
			return nil, context.NewError(ecodes.Failed, "暂时无法加款")
		}
		return NewRecordResult(ecodes.HasExists, row), nil
	}
	row, err := change(db, acc.ID, tradeNo, "0", tradeType, changeType, amount, ext)
	if err != nil {
		return nil, err
	}
	return NewRecordResult(ecodes.Success, row), nil
}

//DrawingAmount 资金提款
func DrawingAmount(db db.IDBExecuter, ident string, group string, eid string, tradeNo string, tradeType int, changeType int, amount int, ext string) (*RecordResult, error) {

	acc, err := GetAccount(db, ident, group, eid)
	if err != nil {
		return nil, err
	}

	b, err := exists(db, acc.ID, tradeNo, 0, tradeType, changeType)
	if err != nil {
		return nil, err
	}
	if b {
		row, err := getRecordByTradeNo(db, acc.ID, tradeNo, tradeType, changeType)
		if err != nil {
			return nil, context.NewError(ecodes.Failed, "暂时无法提款")
		}
		return NewRecordResult(ecodes.HasExists, row), nil
	}
	row, err := change(db, acc.ID, tradeNo, "0", tradeType, changeType, -1*amount, ext)
	if err != nil {
		return nil, err
	}
	return NewRecordResult(ecodes.Success, row), nil
}

//DeductAmount 资金扣款
func DeductAmount(db db.IDBExecuter, ident string, group string, eid string, tradeNo string, tradeType int, amount int, ext string) (*RecordResult, error) {
	if amount == 0 {
		return nil, context.NewErrorf(ecodes.AmountErr, "金额错误%d", amount)
	}
	acc, err := GetAccount(db, ident, group, eid)
	if err != nil {
		return nil, err
	}
	b, err := exists(db, acc.ID, tradeNo, 0, tradeType, ttypes.Deduct)
	if err != nil {
		return nil, err
	}
	if b {
		row, err := getRecordByTradeNo(db, acc.ID, tradeNo, tradeType, ttypes.Deduct)
		if err != nil {
			return nil, context.NewError(ecodes.Failed, "暂时无法扣款")
		}
		return NewRecordResult(ecodes.HasExists, row), nil
	}
	row, err := change(db, acc.ID, tradeNo, "0", tradeType, ttypes.Deduct, -amount, ext)
	if err != nil {
		return nil, err
	}
	return NewRecordResult(ecodes.Success, row), nil
}

//RefundAmount 资金退款
func RefundAmount(db db.IDBExecuter, ident string, group string, eid string, tradeNo string, extNo string, tradeType int, amount int, ext string) (*RecordResult, error) {
	if amount == 0 {
		return nil, context.NewErrorf(ecodes.AmountErr, "金额错误%d", amount)
	}

	acc, err := GetAccount(db, ident, group, eid)
	if err != nil {
		return nil, err
	}

	//检查锁交易记录
	deductAmount, err := lockTradeRecord(db, acc.ID, extNo, tradeType, ttypes.Deduct)
	if err != nil {
		return nil, err
	}
	if deductAmount == 0 {
		return nil, context.NewErrorf(ecodes.NotExists, "交易编号(%s)不存在", extNo)
	}

	// 查询已扣款
	refundAmount, err := queryTradedAmount(db, acc.ID, extNo, tradeType, ttypes.Refund)
	if err != nil {
		return nil, err
	}

	if deductAmount < refundAmount+amount {
		return nil, context.NewErrorf(ecodes.AmountErr, "扣款金额:%d,已退款金额:%d,本次退款金额:%d", deductAmount, refundAmount, amount)
	}

	row, err := change(db, acc.ID, tradeNo, extNo, tradeType, ttypes.Refund, amount, ext)
	if err != nil {
		return nil, err
	}
	return NewRecordResult(ecodes.Success, row), nil
}

//ReverseAmount 红冲
func ReverseAmount(db db.IDBExecuter, ident string, group string, eid string, tradeNo string, extNo string, tradeType int, changeType int, ext string) (*RecordResult, error) {

	acc, err := GetAccount(db, ident, group, eid)
	if err != nil {
		return nil, err
	}
	amount, err := queryTradedAmount(db, acc.ID, extNo, ttypes.Reverse, changeType)
	if err != nil {
		return nil, err
	}
	if amount != 0 {
		return nil, context.NewErrorf(ecodes.HasExists, "红冲交易编号(%s)已存在", extNo)
	}
	//锁交易记录
	tradeAmount, err := lockTradeRecord(db, acc.ID, extNo, tradeType, changeType)
	if err != nil {
		return nil, err
	}
	if tradeAmount == 0 {
		return nil, context.NewErrorf(ecodes.NotExists, "交易编号(%s)不存在", extNo)
	}

	row, err := change(db, acc.ID, tradeNo, extNo, ttypes.Reverse, changeType, tradeAmount, ext)
	if err != nil {
		return nil, err
	}
	return NewRecordResult(ecodes.Success, row), nil
}

//Query 查询余额变动明细
func Query(db db.IDBExecuter, ident string, group string, eid string, startTime string, endTime string, pi int, ps int) (*RecordResults, error) {
	acc, err := GetAccount(db, ident, group, eid)
	if err != nil {
		return nil, err
	}
	rows, err := query(db, acc.ID, startTime, endTime, pi, ps)
	if err != nil {
		return nil, err
	}
	return NewRecordResults(ecodes.Success, rows), nil
}
