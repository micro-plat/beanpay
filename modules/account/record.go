package account

import (
	"fmt"

	"github.com/micro-plat/beanpay/modules/const/sql"
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
)

type IRecord interface {
	Query(accountID int, startTime string, pi int, ps int) (db.QueryRows, error)
}

type Record struct {
	c component.IContainer
}

func NewRecord(c component.IContainer) *Record {
	return &Record{
		c: c,
	}
}
func (r *Record) Query(accountID int, startTime string, pi int, ps int) (db.QueryRows, error) {
	input := map[string]interface{}{
		"account_id": accountID,
		"start":      startTime,
		"pi":         pi,
		"ps":         ps,
	}
	db := r.c.GetRegularDB()
	rows, s, p, err := db.Query(sql.QueryBalanceRecord, input)
	if err != nil {
		return nil, fmt.Errorf("SQL语句执行出错:%s(%v)", s, p)
	}
	return rows, nil
}
