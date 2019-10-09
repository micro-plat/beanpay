package pkg

import (
	"github.com/micro-plat/beanpay/beanpay/account"
	"github.com/micro-plat/beanpay/beanpay/const/ecodes"
	"github.com/micro-plat/beanpay/beanpay/const/ttypes"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"
)

//Create 创建服务包信息
func Create(db db.IDBExecuter, eid string, sid string, name string, total int, daily int, expires string) (interface{}, error) {

	pkg, err := GetPackage(db, eid, sid)
	if err == nil {
		return context.NewResult(ecodes.HasExists, pkg), nil
	}
	if context.GetCode(err) != ecodes.NotExists {
		return nil, err
	}
	acc, err := account.GetAccount(db, eid)
	if err != nil {
		return nil, err
	}

	err = create(db, acc.ID, sid, name, total, daily, expires)
	if err != nil {
		return nil, err
	}
	return GetPackage(db, eid, sid)
}

//GetPackage 获取服务包编号
func GetPackage(db db.IDBExecuter, eid string, sid string) (pkg *PKG, err error) {
	acc, err := account.GetAccount(db, eid)
	if err != nil {
		return nil, err
	}
	row, err := getPackage(db, acc.ID, sid)
	if err != nil {
		return nil, err
	}
	pkg = &PKG{}
	if err = row.ToStruct(pkg); err != nil {
		return nil, err
	}
	return pkg, nil
}

//AddCapacity 添加服务包数量
func AddCapacity(db db.IDBExecuter, eid string, sid string, tradeNo string, capacity int) (*context.Result, error) {
	if capacity <= 0 {
		return nil, context.NewErrorf(ecodes.AmountErr, "数量错误%d", capacity)
	}
	pkg, err := GetPackage(db, eid, sid)
	if err != nil {
		return nil, err
	}
	b, err := exists(db, pkg.ID, tradeNo, 0, ttypes.Add)
	if err != nil {
		return nil, err
	}
	if b {
		row, err := getRecordByTradeNo(db, pkg.ID, tradeNo, ttypes.Add)
		if err != nil {
			return nil, context.NewError(ecodes.Failed, "暂时无法添加服务包数量")
		}
		return context.NewResult(ecodes.HasExists, row), nil
	}

	row, err := change(db, pkg.ID, tradeNo, ttypes.Add, capacity)
	if err != nil {
		return nil, err
	}
	return context.NewResult(ecodes.Success, row), nil
}

//DrawingCapacity 添加服务包数量
func DrawingCapacity(db db.IDBExecuter, eid string, sid string, tradeNo string, capacity int) (*context.Result, error) {
	if capacity <= 0 {
		return nil, context.NewErrorf(ecodes.AmountErr, "数量错误%d", capacity)
	}
	pkg, err := GetPackage(db, eid, sid)
	if err != nil {
		return nil, err
	}
	b, err := exists(db, pkg.ID, tradeNo, 0, ttypes.Drawing)
	if err != nil {
		return nil, err
	}
	if b {
		row, err := getRecordByTradeNo(db, pkg.ID, tradeNo, ttypes.Drawing)
		if err != nil {
			return nil, context.NewError(ecodes.Failed, "暂时无法提取服务包数量")
		}
		return context.NewResult(ecodes.HasExists, row), nil
	}

	row, err := change(db, pkg.ID, tradeNo, ttypes.Drawing, -1*capacity)
	if err != nil {
		return nil, err
	}
	return context.NewResult(ecodes.Success, row), nil
}

//DeductCapacity 扣减服务包数量
func DeductCapacity(db db.IDBExecuter, eid string, sid string, tradeNo string, capacity int) (*context.Result, error) {

	if capacity <= 0 {
		return nil, context.NewErrorf(ecodes.AmountErr, "数量错误%d", capacity)
	}
	pkg, err := GetPackage(db, eid, sid)
	if err != nil {
		return nil, err
	}

	b, err := exists(db, pkg.ID, tradeNo, 0, ttypes.Deduct)
	if err != nil {
		return nil, err
	}
	if b {
		row, err := getRecordByTradeNo(db, pkg.ID, tradeNo, ttypes.Deduct)
		if err != nil {
			return nil, context.NewError(ecodes.Failed, "暂时无法扣减服务包数量")
		}
		return context.NewResult(ecodes.HasExists, row), nil
	}

	row, err := change(db, pkg.ID, tradeNo, ttypes.Deduct, -capacity)
	if err != nil {
		return nil, err
	}
	return context.NewResult(ecodes.Success, row), nil
}

//RefundCapacity 退回服务包数量
func RefundCapacity(db db.IDBExecuter, eid string, sid string, tradeNo string, capacity int) (*context.Result, error) {
	if capacity <= 0 {
		return nil, context.NewErrorf(ecodes.AmountErr, "数量错误%d", capacity)
	}
	pkg, err := GetPackage(db, eid, sid)
	if err != nil {
		return nil, err
	}

	//检查是否已退款
	b, err := exists(db, pkg.ID, tradeNo, 0, ttypes.Refund)
	if err != nil {
		return nil, err
	}
	if b {
		row, err := getRecordByTradeNo(db, pkg.ID, tradeNo, ttypes.Refund)
		if err != nil {
			return nil, context.NewError(ecodes.Failed, "暂时无法退款")
		}
		return context.NewResult(ecodes.HasExists, row), nil
	}
	//检查是否存在加款记录
	b, err = exists(db, pkg.ID, tradeNo, capacity, ttypes.Deduct)
	if err != nil {
		return nil, err
	}
	if !b {
		return nil, context.NewError(ecodes.HasExists, "扣款交易编号不存在")
	}

	row, err := change(db, pkg.ID, tradeNo, ttypes.Refund, capacity)
	if err != nil {
		return nil, err
	}
	return context.NewResult(ecodes.Success, row), nil
}

//Query 查询指定服务变的变动明细
func Query(db db.IDBExecuter, eid string, sid string, startTime string, endTime string, pi int, ps int) (db.QueryRows, error) {
	pkg, err := GetPackage(db, eid, sid)
	if err != nil {
		return nil, err
	}
	return query(db, pkg.AccountID, pkg.ID, startTime, endTime, pi, ps)
}
