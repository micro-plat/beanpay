package pkg

import (
	"github.com/micro-plat/beanpay/beanpay/const/ecodes"
	"github.com/micro-plat/beanpay/beanpay/const/sql"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
)

//Change 服务包数量变动
func change(db db.IDBExecuter, pkgID int64, tradeNo string, changeType int, capacity int, ext string) (types.IXMap, error) {
	input := map[string]interface{}{
		"pkg_id":      pkgID,
		"capacity":    capacity,
		"total":       types.DecodeInt(changeType, 1, capacity, 0),
		"trade_no":    tradeNo,
		"change_type": changeType,
		"ext":         ext,
	}
	//修改包数量
	row, err := db.Execute(sql.ChangePackage, input)
	if err != nil {
		return nil, err
	}
	if row == 0 {
		return nil, errs.NewError(ecodes.NotEnough, "服务包剩余数量不足")
	}

	//添加变动记录
	row, err = db.Execute(sql.AddPackageRecord, input)
	if err != nil {
		return nil, err
	}
	data, err := getRecordByTradeNo(db, pkgID, tradeNo, changeType)
	if errs.GetCode(err) == ecodes.NotExists {
		return nil, errs.NewError(ecodes.Failed, "添加资金变动失败")
	}
	return data, nil
}

//Exists 检查记录是否已存在
func exists(db db.IDBExecuter, pkgID int64, tradeNo string, num int, changeType int) (bool, error) {
	input := map[string]interface{}{
		"pkg_id":      pkgID,
		"trade_no":    tradeNo,
		"change_type": changeType,
		"max_num":     num,
	}
	row, err := db.Scalar(sql.ExistsPackageRecord, input)
	if err != nil {
		return false, err
	}
	return types.GetInt(row) != 0, nil
}
