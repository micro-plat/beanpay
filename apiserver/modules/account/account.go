package account

import (
	"github.com/micro-plat/beanpay/apiserver/modules/const/sql"
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/types"
)

type IAccount interface {
	Create(uaid string, name string) (int, error)
	Get(uaid string) (int, error)
}

type Account struct {
	c component.IContainer
}

func NewAccount(c component.IContainer) *Account {
	return &Account{
		c: c,
	}
}

//Create 根据uaid,name创建帐户信息
func (m *Account) Create(uaid string, name string) (int, error) {
	db := m.c.GetRegularDB()
	input := map[string]interface{}{
		"uaid": uaid,
		"name": name,
	}
	_, _, _, err := db.Execute(sql.CreateAccount, input)
	if err != nil {
		return 0, err
	}

	accountID, _, _, err := db.Scalar(sql.GetAccountByUaid, input)
	if err != nil {
		return 0, err
	}
	return types.GetInt(accountID), nil
}

//Get 根据uaid获取帐户编号
func (m *Account) Get(uaid string) (int, error) {
	db := m.c.GetRegularDB()
	input := map[string]interface{}{
		"uaid": uaid,
	}
	rows, _, _, err := db.Query(sql.GetAccountByUaid, input)
	if err != nil {
		return 0, err
	}
	if rows.IsEmpty() {
		return 0, context.NewError(908, "帐户不存在")
	}
	return rows.Get(0).GetInt("account_id"), nil
}
