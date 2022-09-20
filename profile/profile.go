package profile

type ProfileResponse struct {
	RegistrationTime   int64   `json:"registrationTime"`
	UserId             int     `json:"userId"`
	FirstName          string  `json:"firstName"`
	LastName           string  `json:"lastName"`
	Pin                string  `json:"pin"`
	DateOfBirth        int64   `json:"dateOfBirth"`
	State              string  `json:"state"`
	Gender             string  `json:"gender"`
	Mobile             string  `json:"mobile"`
	FirstDepositDate   int64   `json:"firstDepositDate"`
	FirstDepositAmount float64 `json:"firstDepositAmount"`
}

type Profile interface {
	GetUserByID(userId int) (*ProfileResponse, error)
}
