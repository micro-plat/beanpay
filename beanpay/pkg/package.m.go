package pkg

import "encoding/json"

type PKG struct {
	ID             int64  `json:"pkg_id" m2s:"pkg_id"`
	AccountID      int    `json:"account_id" m2s:"account_id"`
	SPKGID         string `json:"spkg_id" m2s:"spkg_id"`
	Name           string `json:"pkg_name" m2s:"pkg_name"`
	TotalCapacity  int    `json:"total_capacity" m2s:"total_capacity"`
	RemainCapacity int    `json:"total_remain" m2s:"total_remain"`
	DailyCapacity  int    `json:"capacity_daily" m2s:"capacity_daily"`
	TodayCapacity  int    `json:"deduct_today" m2s:"deduct_today"`
	Expires        string `json:"expires" m2s:"expires"`
	BookTime       string `json:"book_time" m2s:"book_time"`
	LastUpdate     string `json:"last_update" m2s:"last_update"`
}

func NewPKG(j string) (*PKG, error) {
	var pkg PKG
	err := json.Unmarshal([]byte(j), &pkg)
	return &pkg, err
}
