package jwr

type UserProfile struct {
	FirstName  string `json:"firstName"`
	MiddleName string `json:"middleName"`
	LastName   string `json:"lastName"`
	DOB        int64  `json:"dateOfBirth"`
	Pin        string `json:"pin"`
	Address    string `json:"address"`
	Address2   string `json:"address2"`
	City       string `json:"city"`
	State      string `json:"state"`
	Gender     string `json:"gender"`
}

type JWRSDKConfig struct {
	BaseURL    string
	Token      string
	APITimeout int
}
