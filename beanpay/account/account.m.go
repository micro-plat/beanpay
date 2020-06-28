package account

import "encoding/json"

type Account struct {
	ID      int     `json:"account_id" m2s:"account_id"`
	Name    string  `json:"account_name" m2s:"account_name"`
	eid     string  `json:"eid" m2s:"eid"`
	Balance float64 `json:"balance" m2s:"balance"`
	Credit  float64 `json:"credit" m2s:"credit"`
}

func NewAccount(s string) (*Account, error) {
	var account Account
	err := json.Unmarshal([]byte(s), &account)
	return &account, err
}

type AccountResult struct {
	*Account
	code int
}

func NewAccountResult(code int, account *Account) *AccountResult {
	return &AccountResult{Account: account, code: code}
}

func (r *AccountResult) GetResult() interface{} {
	return r.Account
}
func (r *AccountResult) GetCode() int {
	return r.code
}

type AccountInfo struct {
	ID      int     `json:"account_id" m2s:"account_id"`
	Name    string  `json:"account_name" m2s:"account_name"`
	EID     string  `json:"eid" m2s:"eid"`
	Balance float64 `json:"balance" m2s:"balance"`
	Credit  float64 `json:"credit" m2s:"credit"`
	Group   string  `json:"groups" m2s:"groups"`
	Status  int     `json:"status" m2s:"status"`
}

type AccountInfoList struct {
	Count int            `json:"count" m2s:"count"`
	Data  []*AccountInfo `json:"data" m2s:"data"`
}
