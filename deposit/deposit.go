package deposit

type DepositResponse struct {
	UserId             int     `json:"userId"`
	MobileNumber       string  `json:"mobileNumber"`
	FirstDepositTime   int64   `json:"firstDepositTime"`
	FirstDepositAmount float64 `json:"firstDepositAmount"`
	LastDepositTime    int64   `json:"lastDepositTime"`
	LastDepositAmount  float64 `json:"lastDepositAmount"`
	TotalDepositAmount float64 `json:"totalDepositAmount"`
}

type Deposit interface {
	GetFirstDepositFromHouzat(mobile string) (*DepositResponse, error)
}

type DepositConfig interface {
	GetDepositEndpoint() string
	GetDepositAuthToken() string
	GetDepositAPIKey() string
}
