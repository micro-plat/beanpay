package pkg

import (
	"fmt"

	"github.com/micro-plat/beanpay/apiserver/modules/const/sql"
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
func (r *Record) Query(pkgID int, startTime string, pi int, ps int) (db.QueryRows, error) {
	input := map[string]interface{}{
		"pkg_id": pkgID,
		"start":  startTime,
		"pi":     pi,
		"ps":     ps,
	}
	db := r.c.GetRegularDB()
	rows, s, p, err := db.Query(sql.QueryPackageRecord, input)
	if err != nil {
		return nil, fmt.Errorf("SQL语句执行出错:%s(%v)", s, p)
	}
	return rows, nil
}
