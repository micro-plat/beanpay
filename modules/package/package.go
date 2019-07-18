package pkg

import (
	"github.com/micro-plat/beanpay/modules/const/sql"
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/types"
)

type IPackage interface {

	//Create 根据帐户编号，包编号，名称，总数，日限制数，过期时间
	Create(accountID int, spkgID string, name string, total int, daily int, expires string) (int, error)

	//Get 根据帐户编号，外部包编号获取当前系统包编号
	Get(accountID int, spkgID string) (int, error)
}

type Package struct {
	c component.IContainer
}

func NewPackage(c component.IContainer) *Package {
	return &Package{
		c: c,
	}
}

//Create 根据帐户编号，包编号，名称，总数，日限制数，过期时间
func (m *Package) Create(accountID int, spkgID string, name string, total int, daily int, expires string) (int, error) {
	db := m.c.GetRegularDB()
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

//Get 根据帐户编号，外部包编号获取当前系统包编号
func (m *Package) Get(accountID int, spkgID string) (int, error) {
	db := m.c.GetRegularDB()
	input := map[string]interface{}{
		"account_id": accountID,
		"spkg_id":    spkgID,
	}
	rows, _, _, err := db.Query(sql.GetPackageBySPKG, input)
	if err != nil {
		return 0, err
	}
	if rows.IsEmpty() {
		return 0, context.NewError(908, "服务包不存在")
	}
	return rows.Get(0).GetInt("pkg_id"), nil
}
