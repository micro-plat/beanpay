package pkg

import (
	"github.com/micro-plat/beanpay/beanpay/const/ecodes"
	"github.com/micro-plat/beanpay/beanpay/const/sql"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
)

//Create 根据帐户编号，包编号，名称，总数，日限制数，过期时间
func create(db db.IDBExecuter, accountID int, spkgID string, name string, total int, daily int, expires string) (int, error) {
	input := map[string]interface{}{
		"account_id": accountID,
		"spkg_id":    spkgID,
		"name":       name,
		"total":      total,
		"daily":      daily,
		"expires":    expires,
	}
	_, _, _, err := db.Execute(sql.CreatePackage, input)
	if err != nil {
		return 0, err
	}

	pkgID, _, _, err := db.Scalar(sql.GetPackageBySPKG, input)
	if err != nil {
		return 0, err
	}
	return types.GetInt(pkgID), nil
}

//GetPackageID 根据帐户编号，外部包编号获取当前系统包编号
func getPackage(db db.IDBExecuter, accountID int, spkgID string) (db.QueryRow, error) {
	input := map[string]interface{}{
		"account_id": accountID,
		"spkg_id":    spkgID,
	}
	rows, _, _, err := db.Query(sql.GetPackageBySPKG, input)
	if err != nil {
		return nil, err
	}
	if rows.IsEmpty() {
		return nil, context.NewError(ecodes.NotExists, "服务包不存在")
	}
	return rows.Get(0), nil
}
func getRecordByTradeNo(db db.IDBExecuter, pkgID int64, tradeNo string, tp int) (db.QueryRow, error) {
	rows, _, _, err := db.Query(sql.GetPackageRecordByTradeNo, map[string]interface{}{
		"pkg_id":   pkgID,
		"trade_no": tradeNo,
		"tp":       tp,
	})
	if err != nil {
		return nil, err
	}
	if rows.IsEmpty() {
		return nil, context.NewError(ecodes.NotExists, "记录不存在")
	}
	return rows.Get(0), nil
}
