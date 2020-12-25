package pkg

import (
	"github.com/micro-plat/beanpay/beanpay/account"
	"github.com/micro-plat/beanpay/beanpay/const/ecodes"
	"github.com/micro-plat/beanpay/beanpay/const/ttypes"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/errs"
)

//Create 创建服务包信息
func Create(db db.IDBExecuter, ident string, group string, eid string, sid string, name string, total int, daily int, expires string) (interface{}, error) {

	pkg, err := GetPackage(db, ident, group, eid, sid)
	if err == nil {
		return errs.NewResult(ecodes.HasExists, pkg), nil
	}
	if errs.GetCode(err) != ecodes.NotExists {
		return nil, err
	}
	acc, err := account.GetAccount(db, ident, group, eid)
	if err != nil {
		return nil, err
	}

	err = create(db, acc.ID, sid, name, total, daily, expires)
	if err != nil {
		return nil, err
	}
	return GetPackage(db, ident, group, eid, sid)
}

//GetPackage 获取服务包编号
func GetPackage(db db.IDBExecuter, ident string, group string, eid string, sid string) (pkg *PKG, err error) {
	acc, err := account.GetAccount(db, ident, group, eid)
	if err != nil {
		return nil, err
	}
	row, err := getPackage(db, acc.ID, sid)
	if err != nil {
		return nil, err
	}
	pkg = &PKG{}
	if err = row.ToAnyStruct(pkg); err != nil {
		return nil, err
	}
	return pkg, nil
}

//AddCapacity 添加服务包数量
func AddCapacity(db db.IDBExecuter, ident string, group string, eid string, sid string, tradeNo string, capacity int, ext string) (*errs.Result, error) {
	if capacity <= 0 {
		return nil, errs.NewErrorf(ecodes.AmountErr, "数量错误%d", capacity)
	}
	pkg, err := GetPackage(db, ident, group, eid, sid)
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
			return nil, errs.NewError(ecodes.Failed, "暂时无法添加服务包数量")
		}
		return errs.NewResult(ecodes.HasExists, row), nil
	}

	row, err := change(db, pkg.ID, tradeNo, ttypes.Add, capacity, ext)
	if err != nil {
		return nil, err
	}
	return errs.NewResult(ecodes.Success, row), nil
}

//DrawingCapacity 添加服务包数量
func DrawingCapacity(db db.IDBExecuter, ident string, group string, eid string, sid string, tradeNo string, capacity int, ext string) (*errs.Result, error) {
	if capacity <= 0 {
		return nil, errs.NewErrorf(ecodes.AmountErr, "数量错误%d", capacity)
	}
	pkg, err := GetPackage(db, ident, group, eid, sid)
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
			return nil, errs.NewError(ecodes.Failed, "暂时无法提取服务包数量")
		}
		return errs.NewResult(ecodes.HasExists, row), nil
	}

	row, err := change(db, pkg.ID, tradeNo, ttypes.Drawing, -1*capacity, ext)
	if err != nil {
		return nil, err
	}
	return errs.NewResult(ecodes.Success, row), nil
}

//DeductCapacity 扣减服务包数量
func DeductCapacity(db db.IDBExecuter, ident string, group string, eid string, sid string, tradeNo string, capacity int, ext string) (*errs.Result, error) {

	if capacity <= 0 {
		return nil, errs.NewErrorf(ecodes.AmountErr, "数量错误%d", capacity)
	}
	pkg, err := GetPackage(db, ident, group, eid, sid)
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
			return nil, errs.NewError(ecodes.Failed, "暂时无法扣减服务包数量")
		}
		return errs.NewResult(ecodes.HasExists, row), nil
	}

	row, err := change(db, pkg.ID, tradeNo, ttypes.Deduct, -capacity, ext)
	if err != nil {
		return nil, err
	}
	return errs.NewResult(ecodes.Success, row), nil
}

//RefundCapacity 退回服务包数量
func RefundCapacity(db db.IDBExecuter, ident string, group string, eid string, sid string, tradeNo string, capacity int, ext string) (*errs.Result, error) {
	if capacity <= 0 {
		return nil, errs.NewErrorf(ecodes.AmountErr, "数量错误%d", capacity)
	}
	pkg, err := GetPackage(db, ident, group, eid, sid)
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
			return nil, errs.NewError(ecodes.Failed, "暂时无法退款")
		}
		return errs.NewResult(ecodes.HasExists, row), nil
	}
	//检查是否存在加款记录
	b, err = exists(db, pkg.ID, tradeNo, capacity, ttypes.Deduct)
	if err != nil {
		return nil, err
	}
	if !b {
		return nil, errs.NewError(ecodes.HasExists, "扣款交易编号不存在")
	}

	row, err := change(db, pkg.ID, tradeNo, ttypes.Refund, capacity, ext)
	if err != nil {
		return nil, err
	}
	return errs.NewResult(ecodes.Success, row), nil
}

//Query 查询指定服务变的变动明细
func Query(db db.IDBExecuter, ident string, group string, eid string, sid string, startTime string, endTime string, pi int, ps int) (db.QueryRows, error) {
	pkg, err := GetPackage(db, ident, group, eid, sid)
	if err != nil {
		return nil, err
	}
	return query(db, pkg.AccountID, pkg.ID, startTime, endTime, pi, ps)
}
