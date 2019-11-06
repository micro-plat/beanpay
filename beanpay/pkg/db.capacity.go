package pkg

import (
	"github.com/micro-plat/beanpay/beanpay/const/ecodes"
	"github.com/micro-plat/beanpay/beanpay/const/sql"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
)

//Change 服务包数量变动
func change(db db.IDBExecuter, pkgID int64, tradeNo string, tp int, capacity int, ext string) (types.XMap, error) {
	input := map[string]interface{}{
		"pkg_id":   pkgID,
		"capacity": capacity,
		"total":    types.DecodeInt(tp, 1, capacity, 0),
		"trade_no": tradeNo,
		"tp":       tp,
		"ext":      ext,
	}
	//修改包数量
	row, _, _, err := db.Execute(sql.ChangePackage, input)
	if err != nil {
		return nil, err
	}
	if row == 0 {
		return nil, context.NewError(ecodes.NotEnough, "服务包剩余数量不足")
	}

	//添加变动记录
	row, _, _, err = db.Execute(sql.AddPackageRecord, input)
	if err != nil {
		return nil, err
	}
	data, err := getRecordByTradeNo(db, pkgID, tradeNo, tp)
	if context.GetCode(err) == ecodes.NotExists {
		return nil, context.NewError(ecodes.Failed, "添加资金变动失败")
	}
	return data, nil
}

//Exists 检查记录是否已存在
func exists(db db.IDBExecuter, pkgID int64, tradeNo string, num int, tp int) (bool, error) {
	input := map[string]interface{}{
		"pkg_id":   pkgID,
		"trade_no": tradeNo,
		"tp":       tp,
		"max_num":  num,
	}
	row, _, _, err := db.Scalar(sql.ExistsPackageRecord, input)
	if err != nil {
		return false, err
	}
	return types.GetInt(row) != 0, nil
}
