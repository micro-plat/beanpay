package account

import (
	"github.com/micro-plat/beanpay/beanpay/const/sql"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"
)

//Create 根据eid,name创建帐户信息
func create(db db.IDBExecuter, ident string, groups string, eid string, name string) error {
	input := map[string]interface{}{
		"ident":  ident,
		"groups": groups,
		"eid":    eid,
		"name":   name,
	}
	_, _, _, err := db.Execute(sql.CreateAccount, input)
	if err != nil {
		return err
	}
	return nil
}

//update 根据eid,name修改帐户信息
func update(db db.IDBExecuter, ident string, groups string, eid string, name string) error {
	input := map[string]interface{}{
		"ident":  ident,
		"groups": groups,
		"eid":    eid,
		"name":   name,
	}
	_, _, _, err := db.Execute(sql.UpdateAccount, input)
	if err != nil {
		return err
	}
	return nil
}

//setCreditAmount 设置授信金额
func setCreditAmount(db db.IDBExecuter, credit int, accountID int) error {
	input := map[string]interface{}{
		"credit":     credit,
		"account_id": accountID,
	}
	_, _, _, err := db.Execute(sql.SetCreditAmount, input)
	if err != nil {
		return err
	}
	return nil
}

//GetAccountID 根据eid获取帐户编号
func getAccount(db db.IDBExecuter, ident string, groups string, eid string) (r db.QueryRow, err error) {
	input := map[string]interface{}{
		"ident":  ident,
		"groups": groups,
		"eid":    eid,
	}
	rows, _, _, err := db.Query(sql.GetAccountByeid, input)
	if err != nil {
		return nil, err
	}
	if rows.IsEmpty() {
		return nil, context.NewError(908, "帐户不存在")
	}
	return rows.Get(0), nil

}
