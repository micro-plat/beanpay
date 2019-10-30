package account

import (
	"github.com/micro-plat/beanpay/beanpay/const/ecodes"
	"github.com/micro-plat/beanpay/beanpay/const/ttypes"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"
)

//Create 根据eid,name创建帐户,如果帐户存在直接返回帐户编号
func Create(db db.IDBExecuter, eid string, tp string, name string) (*AccountResult, error) {
	acc, err := GetAccount(db, eid)
	if err == nil {
		return NewAccountResult(ecodes.HasExists, acc), nil
	}
	if context.GetCode(err) != ecodes.NotExists {
		return nil, err
	}
	if err = create(db, eid, tp, name); err != nil {
		return nil, err
	}
	acc, err = GetAccount(db, eid)
	if err != nil {
		return nil, err
	}
	return NewAccountResult(200, acc), nil
}

//GetBalance 获取帐户余额
func GetBalance(db db.IDBExecuter, eid string) (int, error) {
	return getBalance(db, eid)
}

//GetAccount 根据eid获取帐户
func GetAccount(db db.IDBExecuter, eid string) (acc *Account, err error) {
	row, err := getAccount(db, eid)
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
func AddAmount(db db.IDBExecuter, eid string, tradeNo string, amount int) (*RecordResult, error) {
	if amount <= 0 {
		return nil, context.NewErrorf(ecodes.AmountErr, "金额错误%d", amount)
	}
	acc, err := GetAccount(db, eid)
	if err != nil {
		return nil, err
	}

	b, err := exists(db, acc.ID, tradeNo, 0, ttypes.Add, ttypes.OrderTrade)
	if err != nil {
		return nil, err
	}
	if b {
		row, err := getRecordByTradeNo(db, acc.ID, tradeNo, ttypes.Add, 1)
		if err != nil {
			return nil, context.NewError(ecodes.Failed, "暂时无法加款")
		}
		return NewRecordResult(ecodes.HasExists, row), nil
	}
	row, err := change(db, acc.ID, tradeNo, "0", ttypes.Add, amount, ttypes.OrderTrade)
	if err != nil {
		return nil, err
	}
	return NewRecordResult(ecodes.Success, row), nil
}

//DrawingAmount 资金提款
func DrawingAmount(db db.IDBExecuter, eid string, tradeNo string, amount int) (*RecordResult, error) {
	if amount <= 0 {
		return nil, context.NewErrorf(ecodes.AmountErr, "金额错误%d", amount)
	}
	acc, err := GetAccount(db, eid)
	if err != nil {
		return nil, err
	}

	b, err := exists(db, acc.ID, tradeNo, 0, ttypes.Drawing, ttypes.OrderTrade)
	if err != nil {
		return nil, err
	}
	if b {
		row, err := getRecordByTradeNo(db, acc.ID, tradeNo, ttypes.Drawing, ttypes.OrderTrade)
		if err != nil {
			return nil, context.NewError(ecodes.Failed, "暂时无法提款")
		}
		return NewRecordResult(ecodes.HasExists, row), nil
	}
	row, err := change(db, acc.ID, tradeNo, "0", ttypes.Drawing, -1*amount, ttypes.OrderTrade)
	if err != nil {
		return nil, err
	}
	return NewRecordResult(ecodes.Success, row), nil
}

//DeductAmount 资金扣款
func DeductAmount(db db.IDBExecuter, eid string, tradeNo string, amount int, tradeType int) (*RecordResult, error) {
	if amount == 0 {
		return nil, context.NewErrorf(ecodes.AmountErr, "金额错误%d", amount)
	}
	acc, err := GetAccount(db, eid)
	if err != nil {
		return nil, err
	}
	b, err := exists(db, acc.ID, tradeNo, 0, ttypes.Deduct, tradeType)
	if err != nil {
		return nil, err
	}
	if b {
		row, err := getRecordByTradeNo(db, acc.ID, tradeNo, ttypes.Deduct, tradeType)
		if err != nil {
			return nil, context.NewError(ecodes.Failed, "暂时无法扣款")
		}
		return NewRecordResult(ecodes.HasExists, row), nil
	}
	row, err := change(db, acc.ID, tradeNo, "0", ttypes.Deduct, -amount, tradeType)
	if err != nil {
		return nil, err
	}
	return NewRecordResult(ecodes.Success, row), nil
}

//RefundAmount 资金退款
func RefundAmount(db db.IDBExecuter, eid string, tradeNo string, deductNo string, amount int, tradeType int) (*RecordResult, error) {
	if amount == 0 {
		return nil, context.NewErrorf(ecodes.AmountErr, "金额错误%d", amount)
	}

	acc, err := GetAccount(db, eid)
	if err != nil {
		return nil, err
	}

	//检查锁扣款记录
	deductAmount, err := lockDuductRecord(db, acc.ID, deductNo, ttypes.Deduct, ttypes.OrderTrade)
	if err != nil {
		return nil, err
	}
	if deductAmount == 0 {
		return nil, context.NewErrorf(ecodes.HasExists, "扣款交易编号(%s)不存在", deductNo)
	}

	// 查询已扣款
	refundAmount, err := queryRefundAmount(db, acc.ID, deductNo, ttypes.Refund, tradeType)
	if err != nil {
		return nil, err
	}

	if deductAmount < refundAmount+amount {
		return nil, context.NewErrorf(ecodes.AmountErr, "扣款金额:%d,已退款金额:%d,本次退款金额:%d", deductAmount, refundAmount, amount)
	}

	row, err := change(db, acc.ID, tradeNo, deductNo, ttypes.Refund, amount, tradeType)
	if err != nil {
		return nil, err
	}
	return NewRecordResult(ecodes.Success, row), nil
}

//Query 查询余额变动明细
func Query(db db.IDBExecuter, eid string, startTime string, endTime string, pi int, ps int) (*RecordResults, error) {
	acc, err := GetAccount(db, eid)
	if err != nil {
		return nil, err
	}
	rows, err := query(db, acc.ID, startTime, endTime, pi, ps)
	if err != nil {
		return nil, err
	}
	return NewRecordResults(200, rows), nil
}
