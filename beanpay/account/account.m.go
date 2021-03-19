package account

import "encoding/json"

type Account struct {
	ID      int     `json:"account_id"`
	Name    string  `json:"account_name"`
	eid     string  `json:"eid"`
	Balance float64 `json:"balance"`
	Credit  float64 `json:"credit"`
	Status  int     `json:"status"`
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
	ID      int     `json:"account_id"`
	Name    string  `json:"account_name"`
	EID     string  `json:"eid"`
	Balance float64 `json:"balance"`
	Credit  float64 `json:"credit"`
	Group   string  `json:"groupx"`
	Status  int     `json:"status"`
}

type AccountInfoList struct {
	Count int            `json:"count"`
	Data  []*AccountInfo `json:"data"`
}
