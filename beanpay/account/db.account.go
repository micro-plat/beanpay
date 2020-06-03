package account

import (
	"github.com/micro-plat/beanpay/beanpay/const/sql"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
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
func setCreditAmount(db db.IDBExecuter, credit float64, accountID int) error {
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

//queryAccount 查询账户
func queryAccount(db db.IDBExecuter, ident, group, eid, accountType, name string, pi, ps, status int) (r *AccountInfoList, err error) {
	input := map[string]interface{}{
		"ident":        ident,
		"eid":          eid,
		"groups":       group,
		"status":       status,
		"account_name": name,
		"types":        accountType,
		"currentPage":  (pi - 1) * ps,
		"size":         pi * ps,
		"pageSize":     ps,
	}
	count, _, _, err := db.Scalar(sql.QueryAccountListCount, input)
	if err != nil {
		return nil, err
	}
	if types.GetInt(count) == 0 {
		return nil, nil
	}
	rows, _, _, err := db.Query(sql.QueryAccountList, input)
	if err != nil {
		return nil, err
	}
	accounts := make([]*AccountInfo, rows.Len())

	if err := types.Map2Struct(rows, &accounts); err != nil {
		return nil, err
	}

	return &AccountInfoList{
		Count: types.GetInt(count),
		Data:  accounts,
	}, nil

}
