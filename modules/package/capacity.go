package pkg

import (
	"github.com/micro-plat/beanpay/modules/const/sql"
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/types"
)

type ICapacity interface {
}

type Capacity struct {
	c component.IContainer
}

func NewCapacity(c component.IContainer) *Capacity {
	return &Capacity{
		c: c,
	}
}

//Query 查询包剩余数量
func (b *Capacity) Query(accountID int, spkgID string) (int, error) {
	input := map[string]interface{}{
		"account_id": accountID,
		"spkg_id":    spkgID,
	}

	db := b.c.GetRegularDB()

	//修改帐户余额
	rows, _, _, err := db.Query(sql.GetPackageBySPKG, input)
	if err != nil {
		return 0, err
	}
	if rows.IsEmpty() {
		return 0, context.NewError(908, "包不存在")
	}
	return rows.Get(0).GetInt("total_remain"), nil

}

//Change 资金变动
func (b *Capacity) Change(pkgID int, tradeNo string, tp int, capacity int) error {
	input := map[string]interface{}{
		"pkg_id":   pkgID,
		"capacity": capacity,
		"trade_no": tradeNo,
		"tp":       tp,
	}
	db, err := b.c.GetRegularDB().Begin()
	if err != nil {
		return err
	}

	//修改帐户余额
	row, _, _, err := db.Execute(sql.ChangePackage, input)
	if err != nil {
		db.Rollback()
		return err
	}
	if row == 0 {
		db.Rollback()
		return context.NewError(901, "服务包剩余数量不足")
	}

	//添加资金变动
	row, _, _, err = db.Execute(sql.AddPackageRecord, input)
	if err != nil {
		db.Rollback()
		return err
	}
	if row > 0 {
		db.Commit()
		return nil
	}

	//查询记录是否已存在
	e, _, _, err := db.Scalar(sql.ExistsPackageRecord, input)
	if err != nil {
		db.Rollback()
		return err
	}
	if types.GetInt(e) == 0 {
		db.Rollback()
		return context.NewError(902, "服务包操作失败")
	}
	db.Commit()
	return nil
}

//Exists 检查交易是否已存在
func (b *Capacity) Exists(pkgID int, tradeNo string) (bool, error) {
	input := map[string]interface{}{
		"pkg_id":   pkgID,
		"trade_no": tradeNo,
	}
	db := b.c.GetRegularDB()

	//修改帐户余额
	row, _, _, err := db.Scalar(sql.ExistsPackageRecord, input)
	if err != nil {
		return false, err
	}
	return types.GetInt(row) != 0, nil
}
