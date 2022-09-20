package digilocker

import "net/http"

type Digilocker interface {
	GetRedirectURL() string
	StartKYC(transactionId, referenceId, redirectURL string) (*KYCStartDetails, error)
	CheckAccountstatus(mobile, aadhaar string) (*AccountStatusDetails, error)
	GetAddharDetails(transactionId, referenceId string) (*AadhaarDetails, error)
	Healthcheck() (*HealthcheckResult, error)
	addHeaders(req *http.Request)
}
