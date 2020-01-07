package beanpay

import (
	"github.com/micro-plat/beanpay/beanpay/account"
	"github.com/micro-plat/beanpay/beanpay/pkg"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"
)

//IBeanpay Beanpay接口
type IBeanpay interface {
	IAccount
	IPackage
}

//IAccount Account接口
type IAccount interface {
	CreateAccount(i interface{}, eid string, name string) (*account.AccountResult, error)
	GetAccount(i interface{}, eid string) (*account.Account, error)
	AddAmount(i interface{}, eid string, tradeNo string, amount int, ext ...string) (*account.RecordResult, error)
	DrawingAmount(i interface{}, eid string, tradeNo string, amount int, ext ...string) (*account.RecordResult, error)
	DeductAmount(i interface{}, eid string, tradeNo string, tradeType int, amount int, ext ...string) (*account.RecordResult, error)
	RefundAmount(i interface{}, eid string, tradeNo string, extNo string, tradeType int, amount int, ext ...string) (*account.RecordResult, error)
	QueryAccountRecords(i interface{}, eid string, startTime string, endTime string, pi int, ps int) (*account.RecordResults, error)
	TradeFlatAmount(i interface{}, eid string, tradeNo string, tradeType int, amount int, ext ...string) (*account.RecordResult, error)
	BalanceFlatAmount(i interface{}, eid string, tradeNo string, tradeType int, amount int, ext ...string) (*account.RecordResult, error)
	ReverseAddAmount(i interface{}, eid string, tradeNo string, extNo string, tradeType int, ext ...string) (*account.RecordResult, error)
	ReverseDrawingAmount(i interface{}, eid string, tradeNo string, extNo string, tradeType int, ext ...string) (*account.RecordResult, error)
}

//IPackage Package接口
type IPackage interface {
	CreatePackage(i interface{}, eid string, spid string, name string, total int, daily int, expires string) (interface{}, error)
	GetPackage(i interface{}, eid string, spid string) (*pkg.PKG, error)
	AddCapacity(i interface{}, eid string, spid string, tradeNo string, capacity int, ext ...string) (*context.Result, error)
	DrawingCapacity(i interface{}, eid string, spid string, tradeNo string, capacity int, ext ...string) (*context.Result, error)
	DeductCapacity(i interface{}, eid string, spid string, tradeNo string, capacity int, ext ...string) (*context.Result, error)
	RefundCapacity(i interface{}, eid string, spid string, tradeNo string, capacity int, ext ...string) (*context.Result, error)
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
