package pkg

import (
	"github.com/micro-plat/beanpay/beanpay/const/ecodes"
	"github.com/micro-plat/beanpay/beanpay/const/sql"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
)

//Create 根据帐户编号，包编号，名称，总数，日限制数，过期时间
func create(db db.IDBExecuter, accountID int, spkgID string, name string, total int, daily int, expires string) error {
	input := map[string]interface{}{
		"account_id": accountID,
		"spkg_id":    spkgID,
		"name":       name,
		"total":      total,
		"daily":      types.DecodeInt(daily, 0, total, daily),
		"expires":    expires,
	}
	_, err := db.Execute(sql.CreatePackage, input)
	if err != nil {
		return err
	}
	return nil

}

//GetPackageID 根据帐户编号，外部包编号获取当前系统包编号
func getPackage(db db.IDBExecuter, accountID int, spkgID string) (types.IXMap, error) {
	input := map[string]interface{}{
		"account_id": accountID,
		"spkg_id":    spkgID,
	}
	rows, err := db.Query(sql.GetPackageBySPKG, input)
	if err != nil {
		return nil, err
	}
	if rows.IsEmpty() {
		return nil, errs.NewError(ecodes.NotExists, "服务包不存在")
	}
	return rows.Get(0), nil
}
func getRecordByTradeNo(db db.IDBExecuter, pkgID int64, tradeNo string, changeType int) (types.IXMap, error) {
	rows, err := db.Query(sql.GetPackageRecordByTradeNo, map[string]interface{}{
		"pkg_id":      pkgID,
		"trade_no":    tradeNo,
		"change_type": changeType,
	})
	if err != nil {
		return nil, err
	}
	if rows.IsEmpty() {
		return nil, errs.NewError(ecodes.NotExists, "记录不存在")
	}
	return rows.Get(0), nil
}
