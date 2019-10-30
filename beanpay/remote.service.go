package beanpay

import (
	"fmt"
	"strings"

	"github.com/micro-plat/beanpay/beanpay/account"
	"github.com/micro-plat/beanpay/beanpay/pkg"
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
)

//RemoteBeanpay 支付对象
type RemoteBeanpay struct {
	ident string
	tp    string
}

//NewRemoteBeanpay 构建支付对象,传入外部系统标识，帐户类型
func NewRemoteBeanpay(ident string, accountType ...string) *RemoteBeanpay {
	return &RemoteBeanpay{
		ident: ident,
		tp:    types.GetStringByIndex(accountType, 0),
	}
}

//NewRemoteBeanpayByInternal 构建支付对象，内部系统构建时使用
func NewRemoteBeanpayByInternal(accountType ...string) *RemoteBeanpay {
	return &RemoteBeanpay{
		tp: types.GetStringByIndex(accountType, 0),
	}
}
func (b *RemoteBeanpay) makeEID(eid string) string {
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
func (b *RemoteBeanpay) CreateAccount(i interface{}, eid string, name string, kv ...string) (interface{}, error) {
	request, err := getRequest(i)
	if err != nil {
		return nil, err
	}
	status, result, _, err := request("/account/create", "GET", nil, types.Copy(map[string]interface{}{
		"eid":  b.makeEID(eid),
		"sid":  b.ident,
		"name": name,
	}, kv...), true)
	if err != nil {
		return nil, err
	}
	if status != 200 {
		return nil, context.NewError(status, err)
	}
	return result, nil
}

//GetAccount 根据eid获取资金帐户编号
func (b *RemoteBeanpay) GetAccount(i interface{}, eid string, kv ...string) (*account.Account, error) {
	request, err := getRequest(i)
	if err != nil {
		return nil, err
	}
	status, result, _, err := request("/account/query", "GET", nil, types.Copy(map[string]interface{}{
		"eid": b.makeEID(eid),
		"sid": b.ident,
	}, kv...), true)
	if err != nil {
		return nil, err
	}
	if status != 200 {
		return nil, context.NewError(status, err)
	}
	return account.NewAccount(result)
}

//AddAmount 指定用户编号，交易变号，金额进行资金帐户加款
func (b *RemoteBeanpay) AddAmount(i interface{}, eid string, tradeNo string, amount int, kv ...string) (*context.Result, error) {
	request, err := getRequest(i)
	if err != nil {
		return nil, err
	}
	status, result, _, err := request("/account/balance/add", "GET", nil, types.Copy(map[string]interface{}{
		"eid":      b.makeEID(eid),
		"sid":      b.ident,
		"trade_no": tradeNo,
		"amount":   amount,
	}, kv...), true)
	if err != nil {
		return nil, err
	}
	if status != 200 {
		return nil, context.NewError(status, err)
	}
	return context.NewResult(200, result), nil
}

//DrawingAmount 指定用户编号，交易变号，金额进行资金帐户提款
func (b *RemoteBeanpay) DrawingAmount(i interface{}, eid string, tradeNo string, amount int, kv ...string) (*context.Result, error) {
	request, err := getRequest(i)
	if err != nil {
		return nil, err
	}
	status, result, _, err := request("/account/balance/drawing", "GET", nil, types.Copy(map[string]interface{}{
		"eid":      b.makeEID(eid),
		"sid":      b.ident,
		"trade_no": tradeNo,
		"amount":   amount,
	}, kv...), true)
	if err != nil {
		return nil, err
	}
	if status != 200 {
		return nil, context.NewError(status, err)
	}
	return context.NewResult(200, result), nil
}

//DeductAmount 指定用户编号，交易变号，金额进行资金帐户扣款
func (b *RemoteBeanpay) DeductAmount(i interface{}, eid string, tradeNo string, amount int, kv ...string) (*context.Result, error) {

	request, err := getRequest(i)
	if err != nil {
		return nil, err
	}
	status, result, _, err := request("/account/balance/deduct", "GET", nil, types.Copy(map[string]interface{}{
		"eid":      b.makeEID(eid),
		"sid":      b.ident,
		"trade_no": tradeNo,
		"amount":   amount,
	}, kv...), true)
	if err != nil {
		return nil, err
	}
	if status != 200 {
		return nil, context.NewError(status, err)
	}
	return context.NewResult(200, result), nil
}

//RefundAmount 指定用户编号，交易变号，金额进行资金帐户退款
func (b *RemoteBeanpay) RefundAmount(i interface{}, eid string, tradeNo string, reductNo string, amount int, kv ...string) (*context.Result, error) {
	request, err := getRequest(i)
	if err != nil {
		return nil, err
	}
	status, result, _, err := request("/account/balance/refund", "GET", nil, types.Copy(map[string]interface{}{
		"eid":       b.makeEID(eid),
		"sid":       b.ident,
		"trade_no":  tradeNo,
		"reduct_no": reductNo,
		"amount":    amount,
	}, kv...), true)
	if err != nil {
		return nil, err
	}
	if status != 200 {
		return nil, context.NewError(status, err)
	}
	return context.NewResult(200, result), nil
}

//QueryAccountRecords 查询指定用户在一段时间内的资金变动信息
func (b *RemoteBeanpay) QueryAccountRecords(i interface{}, eid string, startTime string, endTime string, pi int, ps int, kv ...string) (db.QueryRows, error) {
	request, err := getRequest(i)
	if err != nil {
		return nil, err
	}
	status, result, _, err := request("/account/record/query", "GET", nil, types.Copy(map[string]interface{}{
		"eid":   b.makeEID(eid),
		"sid":   b.ident,
		"start": startTime,
		"end":   endTime,
		"pi":    pi,
		"ps":    ps,
	}, kv...), true)
	if err != nil {
		return nil, err
	}
	if status != 200 {
		return nil, context.NewError(status, err)
	}
	return db.NewQueryRowsByJSON(result)
}

//CreatePackage 根据用户编号， 服务编号，服务名称，服务包可用总数，日限制使用次数，过期时间创建服务包
//用户必须先创建资金帐户
func (b *RemoteBeanpay) CreatePackage(i interface{}, eid string, spid string, name string, total int, daily int, expires string, kv ...string) (interface{}, error) {
	request, err := getRequest(i)
	if err != nil {
		return nil, err
	}
	status, result, _, err := request("/package/create", "GET", nil, types.Copy(map[string]interface{}{
		"eid":     b.makeEID(eid),
		"sid":     b.ident,
		"spid":    spid,
		"name":    name,
		"total":   total,
		"daily":   daily,
		"expires": expires,
	}, kv...), true)
	if err != nil {
		return nil, err
	}
	if status != 200 {
		return nil, context.NewError(status, err)
	}
	return context.NewResult(200, result), nil
}

//GetPackage 根据用户编号，服务编号获取服务包编号
func (b *RemoteBeanpay) GetPackage(i interface{}, eid string, spid string, kv ...string) (*pkg.PKG, error) {
	request, err := getRequest(i)
	if err != nil {
		return nil, err
	}
	status, result, _, err := request("/package/query", "GET", nil, types.Copy(map[string]interface{}{
		"eid":  b.makeEID(eid),
		"sid":  b.ident,
		"spid": spid,
	}, kv...), true)
	if err != nil {
		return nil, err
	}
	if status != 200 {
		return nil, context.NewError(status, err)
	}
	return pkg.NewPKG(result)
}

//AddCapacity 指定用户编号，交易变号，金额进行服务包数量追加
func (b *RemoteBeanpay) AddCapacity(i interface{}, eid string, spid string, tradeNo string, capacity int, kv ...string) (*context.Result, error) {
	request, err := getRequest(i)
	if err != nil {
		return nil, err
	}
	status, result, _, err := request("/package/capacity/add", "GET", nil, types.Copy(map[string]interface{}{
		"eid":      b.makeEID(eid),
		"sid":      b.ident,
		"trade_no": tradeNo,
		"capacity": capacity,
	}, kv...), true)
	if err != nil {
		return nil, err
	}
	if status != 200 {
		return nil, context.NewError(status, err)
	}
	return context.NewResult(200, result), nil
}

//DrawingCapacity 指定用户编号，交易变号，金额进行服务包数量追加
func (b *RemoteBeanpay) DrawingCapacity(i interface{}, eid string, spid string, tradeNo string, capacity int, kv ...string) (*context.Result, error) {
	request, err := getRequest(i)
	if err != nil {
		return nil, err
	}
	status, result, _, err := request("/package/capacity/drawing", "GET", nil, types.Copy(map[string]interface{}{
		"eid":      b.makeEID(eid),
		"sid":      b.ident,
		"trade_no": tradeNo,
		"capacity": capacity,
	}, kv...), true)
	if err != nil {
		return nil, err
	}
	if status != 200 {
		return nil, context.NewError(status, err)
	}
	return context.NewResult(200, result), nil
}

//DeductCapacity 指定用户编号，交易变号，金额进行服务包数量扣减
func (b *RemoteBeanpay) DeductCapacity(i interface{}, eid string, spid string, tradeNo string, capacity int, kv ...string) (*context.Result, error) {
	request, err := getRequest(i)
	if err != nil {
		return nil, err
	}
	status, result, _, err := request("/package/capacity/deduct", "GET", nil, types.Copy(map[string]interface{}{
		"eid":      b.makeEID(eid),
		"sid":      b.ident,
		"spid":     spid,
		"trade_no": tradeNo,
		"capacity": capacity,
	}, kv...), true)
	if err != nil {
		return nil, err
	}
	if status != 200 {
		return nil, context.NewError(status, err)
	}
	return context.NewResult(200, result), nil

}

//RefundCapacity 指定用户编号，交易变号，金额进行服务包数量退回
func (b *RemoteBeanpay) RefundCapacity(i interface{}, eid string, spid string, tradeNo string, capacity int, kv ...string) (*context.Result, error) {
	request, err := getRequest(i)
	if err != nil {
		return nil, err
	}
	status, result, _, err := request("/package/capacity/refund", "GET", nil, types.Copy(map[string]interface{}{
		"eid":      b.makeEID(eid),
		"sid":      b.ident,
		"spid":     spid,
		"trade_no": tradeNo,
		"capacity": capacity,
	}, kv...), true)
	if err != nil {
		return nil, err
	}
	if status != 200 {
		return nil, context.NewError(status, err)
	}
	return context.NewResult(200, result), nil
}

//QueryPackageRecords 查询指定用户在一段时间内的服务包的变动记录
func (b *RemoteBeanpay) QueryPackageRecords(i interface{}, eid string, spid string, startTime string, endTime string, pi int, ps int, kv ...string) (db.QueryRows, error) {
	request, err := getRequest(i)
	if err != nil {
		return nil, err
	}
	status, result, _, err := request("/package/record/query", "GET", nil, types.Copy(map[string]interface{}{
		"eid":   b.makeEID(eid),
		"sid":   b.ident,
		"spid":  spid,
		"start": startTime,
		"end":   endTime,
		"pi":    pi,
		"ps":    ps,
	}, kv...), true)
	if err != nil {
		return nil, err
	}
	if status != 200 {
		return nil, context.NewError(status, err)
	}
	return db.NewQueryRowsByJSON(result)
}

func getRequest(c interface{}) (rpcRequest, error) {
	switch v := c.(type) {
	case *context.Context:
		return v.GetContainer().Request, nil
	case component.IContainer:
		return v.Request, nil
	case rpcRequest:
		return v, nil
	default:
		return nil, fmt.Errorf("不支持的参数类型")
	}
}

type rpcRequest func(service string, method string, header map[string]string, input map[string]interface{}, failFast bool) (status int, result string, param map[string]string, err error)
