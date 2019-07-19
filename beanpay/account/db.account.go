package account

import (
	"github.com/micro-plat/beanpay/beanpay/const/sql"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
)

//Create 根据uaid,name创建帐户信息
func create(db db.IDBExecuter, uaid string, name string) (int, error) {
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

//GetAccountID 根据uaid获取帐户编号
func getAccountID(db db.IDBExecuter, uaid string) (int, error) {
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
