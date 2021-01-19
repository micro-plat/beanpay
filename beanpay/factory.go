package beanpay

import (
	"github.com/micro-plat/beanpay/beanpay/account"
	"github.com/micro-plat/beanpay/beanpay/const/ecodes"
	"github.com/micro-plat/beanpay/beanpay/const/ttypes"
	"github.com/micro-plat/beanpay/beanpay/pkg"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/errs"
)

//TradeType 交易类型
type TradeType = ttypes.TradeType

// 交易类型
var (

	//AccountTradeType 账户交易
	AccountTradeType TradeType = ttypes.Account

	//FreeTradeType 费用交易
	FreeTradeType TradeType = ttypes.Free

	//CommissionTradeType 佣金交易
	CommissionTradeType TradeType = ttypes.Commission

	//ReverseTradeType 红冲交易
	ReverseTradeType TradeType = ttypes.Reverse
)

var (

	//Success 成功
	Success = ecodes.Success

	//Failed　901失败
	Failed = ecodes.Failed

	//AmountErr　903金额错误
	AmountErr = ecodes.AmountErr

	//NotEnough　904余额不足
	NotEnough = ecodes.NotEnough

	//NotExists　908交易不存在
	NotExists = ecodes.NotExists

	//HasExists　201交易已存在
	HasExists = ecodes.HasExists
)

//IBeanpay Beanpay接口
type IBeanpay interface {
	IAccount
	IPackage
}

//IAccount Account接口
type IAccount interface {
	CreateAccount(i interface{}, eid string, name string) (*account.AccountResult, error)
	SetAccountName(i interface{}, eid string, name string) (*account.AccountResult, error)
	GetAccount(i interface{}, eid string) (*account.Account, error)
	QueryAccount(i interface{}, eid, accountType, name, status string, pi, ps int) (r *account.AccountInfoList, err error)
	AddAmount(i interface{}, eid string, tradeNo string, amount float64, memo string, ext ...string) (*account.RecordResult, error)
	DrawingAmount(i interface{}, eid string, tradeNo string, amount float64, memo string, ext ...string) (*account.RecordResult, error)
	DeductAmount(i interface{}, eid string, tradeNo string, tradeType TradeType, amount float64, memo string, ext ...string) (*account.RecordResult, error)
	RefundAmount(i interface{}, eid string, tradeNo string, extNo string, tradeType TradeType, amount float64, memo string, ext ...string) (*account.RecordResult, error)
	QueryAccountRecords(i interface{}, accountType string, accountID string, accountName string, group string, changeType string, tradeType string, eid string, startTime string, endTime string, pi int, ps int) (*account.RecordResults, error)
	TradeFlatAmount(i interface{}, eid string, tradeNo string, tradeType TradeType, amount float64, memo string, ext ...string) (*account.RecordResult, error)
	BalanceFlatAmount(i interface{}, eid string, tradeNo string, tradeType TradeType, amount float64, memo string, ext ...string) (*account.RecordResult, error)
	ReverseAddAmount(i interface{}, eid string, tradeNo string, extNo string, memo string, ext ...string) (*account.RecordResult, error)
	ReverseDrawingAmount(i interface{}, eid string, tradeNo string, extNo string, memo string, ext ...string) (*account.RecordResult, error)
	SetCreditAmount(i interface{}, eid string, credit float64) (*account.AccountResult, error)
}

//IPackage Package接口
type IPackage interface {
	CreatePackage(i interface{}, eid string, spid string, name string, total int, daily int, expires string) (interface{}, error)
	GetPackage(i interface{}, eid string, spid string) (*pkg.PKG, error)
	AddCapacity(i interface{}, eid string, spid string, tradeNo string, capacity int, ext ...string) (*errs.Result, error)
	DrawingCapacity(i interface{}, eid string, spid string, tradeNo string, capacity int, ext ...string) (*errs.Result, error)
	DeductCapacity(i interface{}, eid string, spid string, tradeNo string, capacity int, ext ...string) (*errs.Result, error)
	RefundCapacity(i interface{}, eid string, spid string, tradeNo string, capacity int, ext ...string) (*errs.Result, error)
	QueryPackageRecords(i interface{}, eid string, spid string, startTime string, endTime string, pi int, ps int) (db.QueryRows, error)
}

const (
	//TPTrade 交易
	TPTrade = 1
	//TPFree 手续费
	TPFree = 2
	//TPCommission 佣金
	TPCommission = 3
	// TPReverse 红冲
	TPReverse = 4
)

//GetAccount 获取Account操作类
func GetAccount(ident string, group string) IAccount {
	return newBeanpay(ident, group)
}

//GetPackage 获取package操作类
func GetPackage(ident string, group string) IPackage {
	return newBeanpay(ident, group)
}
