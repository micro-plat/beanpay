package ttypes

// 变动类型
const (
	Add         = 1
	Drawing     = 2
	Deduct      = 3
	Refund      = 4
	TradeFlat   = 5
	BalanceFlat = 6
)

// 交易类型

type TradeType int

const (
	Account    TradeType = 1
	Free       TradeType = 2
	Commission TradeType = 3
	Reverse    TradeType = 4
)
