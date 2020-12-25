package account

import (
	"github.com/micro-plat/beanpay/beanpay/const/ecodes"
	"github.com/micro-plat/beanpay/beanpay/const/ttypes"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/errs"
)

//Create 根据eid,name创建帐户,如果帐户存在直接返回帐户编号
func Create(db db.IDBExecuter, ident string, group string, eid string, name string) (*AccountResult, error) {
	acc, err := GetAccount(db, ident, group, eid)
	if err == nil {
		return NewAccountResult(ecodes.HasExists, acc), nil
	}
	if errs.GetCode(err) != ecodes.NotExists {
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

//SetAccountName 根据eid,name创建帐户,如果帐户存在直接返回帐户编号
func SetAccountName(db db.IDBExecuter, ident string, group string, eid string, name string) (*AccountResult, error) {
	acc, err := GetAccount(db, ident, group, eid)
	if err != nil {
		return nil, err
	}

	if err = update(db, ident, group, eid, name); err != nil {
		return nil, err
	}

	acc, err = GetAccount(db, ident, group, eid)
	if err != nil {
		return nil, err
	}
	return NewAccountResult(ecodes.Success, acc), nil
}

//SetCreditAmount 设置授信金额
func SetCreditAmount(db db.IDBExecuter, ident string, group string, eid string, credit float64) (*AccountResult, error) {
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
	if err = row.ToAnyStruct(acc); err != nil {
		return nil, err
	}
	return acc, nil
}

// QueryAccount 查询账户列表
func QueryAccount(db db.IDBExecuter, ident, group, eid, accountType, name, status string, pi, ps int) (r *AccountInfoList, err error) {
	return queryAccount(db, ident, group, eid, accountType, name, status, pi, ps)
}

//AddAmount 资金加款
func AddAmount(db db.IDBExecuter, ident string, group string, eid string, tradeNo string, tradeType int, changeType int, amount float64, memo, ext string) (*RecordResult, error) {

	acc, err := GetAccount(db, ident, group, eid)
	if err != nil {
		return nil, err
	}

	b, err := exists(db, acc.ID, tradeNo, tradeType, changeType)
	if err != nil {
		return nil, err
	}
	if b {
		row, err := getRecordByTradeNo(db, acc.ID, tradeNo, tradeType, changeType)
		if err != nil {
			return nil, errs.NewError(ecodes.Failed, "暂时无法加款")
		}
		return NewRecordResult(ecodes.HasExists, row), nil
	}
	row, err := change(db, acc.ID, tradeNo, "0", tradeType, changeType, amount, memo, ext)
	if err != nil {
		return nil, err
	}
	return NewRecordResult(ecodes.Success, row), nil
}

//DrawingAmount 资金提款
func DrawingAmount(db db.IDBExecuter, ident string, group string, eid string, tradeNo string, tradeType int, changeType int, amount float64, memo, ext string) (*RecordResult, error) {

	acc, err := GetAccount(db, ident, group, eid)
	if err != nil {
		return nil, err
	}

	b, err := exists(db, acc.ID, tradeNo, tradeType, changeType)
	if err != nil {
		return nil, err
	}
	if b {
		row, err := getRecordByTradeNo(db, acc.ID, tradeNo, tradeType, changeType)
		if err != nil {
			return nil, errs.NewError(ecodes.Failed, "暂时无法提款")
		}
		return NewRecordResult(ecodes.HasExists, row), nil
	}
	row, err := change(db, acc.ID, tradeNo, "0", tradeType, changeType, -1*amount, memo, ext)
	if err != nil {
		return nil, err
	}
	return NewRecordResult(ecodes.Success, row), nil
}

//DeductAmount 资金扣款
func DeductAmount(db db.IDBExecuter, ident string, group string, eid string, tradeNo string, tradeType int, amount float64, memo, ext string) (*RecordResult, error) {
	if amount == 0 {
		return nil, errs.NewErrorf(ecodes.AmountErr, "金额错误%d", amount)
	}
	acc, err := GetAccount(db, ident, group, eid)
	if err != nil {
		return nil, err
	}
	b, err := exists(db, acc.ID, tradeNo, tradeType, ttypes.Deduct)
	if err != nil {
		return nil, err
	}
	if b {
		row, err := getRecordByTradeNo(db, acc.ID, tradeNo, tradeType, ttypes.Deduct)
		if err != nil {
			return nil, errs.NewError(ecodes.Failed, "暂时无法扣款")
		}
		return NewRecordResult(ecodes.HasExists, row), nil
	}
	row, err := change(db, acc.ID, tradeNo, "0", tradeType, ttypes.Deduct, -amount, memo, ext)
	if err != nil {
		return nil, err
	}
	return NewRecordResult(ecodes.Success, row), nil
}

//RefundAmount 资金退款
func RefundAmount(db db.IDBExecuter, ident string, group string, eid string, tradeNo string, extNo string, tradeType int, amount float64, memo, ext string) (*RecordResult, error) {
	if amount == 0 {
		return nil, errs.NewErrorf(ecodes.AmountErr, "金额错误%d", amount)
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
		return nil, errs.NewErrorf(ecodes.NotExists, "交易编号(%s)不存在", extNo)
	}

	// 检查扣款金额
	b, err := checkRefundAmount(db, acc.ID, tradeNo, extNo, tradeType, ttypes.Refund, deductAmount, amount)
	if err != nil {
		return nil, err
	}
	if b {
		row, err := getRecordByTradeNo(db, acc.ID, tradeNo, tradeType, ttypes.Refund)
		if err != nil {
			return nil, errs.NewError(ecodes.Failed, "暂时无法扣款")
		}
		return NewRecordResult(ecodes.HasExists, row), nil
	}

	row, err := change(db, acc.ID, tradeNo, extNo, tradeType, ttypes.Refund, amount, memo, ext)
	if err != nil {
		return nil, err
	}
	return NewRecordResult(ecodes.Success, row), nil
}

//ReverseAmount 红冲
func ReverseAmount(db db.IDBExecuter, ident string, group string, eid string, tradeNo string, extNo string, tradeType int, changeType int, memo, ext string) (*RecordResult, error) {

	acc, err := GetAccount(db, ident, group, eid)
	if err != nil {
		return nil, err
	}
	amount, err := queryTradedAmount(db, acc.ID, extNo, ttypes.Reverse, changeType)
	if err != nil {
		return nil, err
	}
	if amount != 0 {
		return nil, errs.NewErrorf(ecodes.HasExists, "红冲交易编号(%s)已存在", extNo)
	}
	//锁交易记录
	tradeAmount, err := lockTradeRecord(db, acc.ID, extNo, tradeType, changeType)
	if err != nil {
		return nil, err
	}
	if tradeAmount == 0 {
		return nil, errs.NewErrorf(ecodes.NotExists, "交易编号(%s)不存在", extNo)
	}

	row, err := change(db, acc.ID, tradeNo, extNo, ttypes.Reverse, changeType, tradeAmount, memo, ext)
	if err != nil {
		return nil, err
	}
	return NewRecordResult(ecodes.Success, row), nil
}

//Query 查询余额变动明细
func Query(db db.IDBExecuter, accountType string, accountID string, accountName string, group string, changeType string, tradeType string, eid string, startTime string, endTime string, pi int, ps int) (*RecordResults, error) {
	count, rows, err := query(db, accountType, group, accountID, accountName, changeType, tradeType, startTime, endTime, pi, ps)
	if err != nil {
		return nil, err
	}
	return NewRecordResults(ecodes.Success, count, rows), nil
}
