package account

import (
	"github.com/micro-plat/lib4go/db"
)

type AccountRecord struct {
	AccountRecordID int    `json:"AccountRecord_id" m2s:"AccountRecord_id"`
	RecordID        int    `json:"record_id" m2s:"record_id"`
	TradeNo         string `json:"trade_no" m2s:"trade_no"`
	ChangeType      int    `json:"change_type" m2s:"change_type"`
	Amount          int    `json:"amount" m2s:"amount"`
	Balance         int    `json:"balance" m2s:"balance"`
	CreateTime      string `json:"create_time" m2s:"create_time"`
}
type RecordResult struct {
	*AccountRecord
	code int
}

func NewRecordResult(code int, m db.QueryRow) *RecordResult {
	var account AccountRecord
	m.ToStruct(&account)
	return &RecordResult{AccountRecord: &account, code: code}
}

func (r *RecordResult) GetResult() interface{} {
	return r.AccountRecord
}
func (r *RecordResult) GetCode() int {
	return r.code
}

type RecordResults struct {
	List []*AccountRecord
	code int
}

func NewRecordResults(code int, m db.QueryRows) *RecordResults {
	var accounts []*AccountRecord
	m.ToStruct(&accounts)
	return &RecordResults{List: accounts, code: code}
}

func (r *RecordResults) GetResult() interface{} {
	return r.List
}
func (r *RecordResults) GetCode() int {
	return r.code
}
