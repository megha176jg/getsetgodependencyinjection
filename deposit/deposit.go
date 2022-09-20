package deposit

type DepositResponse struct {
	Amount      float64 `json:"amount"`
	DepositedOn int64   `json:"depositedOn"`
}

type Deposit interface {
	GetFirstDepositFromHouzat(mobile string) (*DepositResponse, error)
}
