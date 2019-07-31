package account

import (
	"github.com/micro-plat/beanpay/beanpay/const/sql"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"
)

//Create 根据uid,name创建帐户信息
func create(db db.IDBExecuter, uid string, name string) error {
	input := map[string]interface{}{
		"uid":  uid,
		"name": name,
	}
	_, _, _, err := db.Execute(sql.CreateAccount, input)
	if err != nil {
		return err
	}
	return nil
}

//GetAccountID 根据uid获取帐户编号
func getAccount(db db.IDBExecuter, uid string) (r db.QueryRow, err error) {
	input := map[string]interface{}{
		"uid": uid,
	}
	rows, _, _, err := db.Query(sql.GetAccountByUid, input)
	if err != nil {
		return nil, err
	}
	if rows.IsEmpty() {
		return nil, context.NewError(908, "帐户不存在")
	}
	return rows.Get(0), nil

}
