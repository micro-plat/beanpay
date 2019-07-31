package account

type Account struct {
	ID      int    `json:"account_id" m2s:"account_id"`
	Name    string `json:"account_name" m2s:"account_name"`
	eid     string `json:"eid" m2s:"eid"`
	Balance int    `json:"balance" m2s:"balance"`
	Credit  int    `json:"credit" m2s:"credit"`
}
