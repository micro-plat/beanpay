package pkg

import (
	"github.com/micro-plat/beanpay/beanpay/const/sql"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
)

//GetPackageRemain 查询包剩余数量
func getPackageRemain(db db.IDBExecuter, accountID int, spkgID string) (int, error) {
	input := map[string]interface{}{
		"account_id": accountID,
		"spkg_id":    spkgID,
	}

	rows, _, _, err := db.Query(sql.GetPackageBySPKG, input)
	if err != nil {
		return 0, err
	}
	if rows.IsEmpty() {
		return 0, context.NewError(908, "包不存在")
	}
	return rows.Get(0).GetInt("total_remain"), nil

}

//Change 服务包数量变动
func change(db db.IDBExecuter, pkgID int, tradeNo string, tp int, capacity int) error {
	input := map[string]interface{}{
		"pkg_id":   pkgID,
		"capacity": capacity,
		"trade_no": tradeNo,
		"tp":       tp,
	}
	//修改包数量
	row, _, _, err := db.Execute(sql.ChangePackage, input)
	if err != nil {
		return err
	}
	if row == 0 {
		return context.NewError(901, "服务包剩余数量不足")
	}

	//添加变动记录
	row, _, _, err = db.Execute(sql.AddPackageRecord, input)
	if err != nil {
		return err
	}
	if row > 0 {
		return nil
	}

	//查询记录是否已存在
	e, _, _, err := db.Scalar(sql.ExistsPackageRecord, input)
	if err != nil {
		return err
	}
	if types.GetInt(e) == 0 {
		return context.NewError(902, "服务包操作失败")
	}
	return nil
}

//Exists 检查记录是否已存在
func exists(db db.IDBExecuter, pkgID int, tradeNo string) (bool, error) {
	input := map[string]interface{}{
		"pkg_id":   pkgID,
		"trade_no": tradeNo,
	}
	row, _, _, err := db.Scalar(sql.ExistsPackageRecord, input)
	if err != nil {
		return false, err
	}
	return types.GetInt(row) != 0, nil
}
