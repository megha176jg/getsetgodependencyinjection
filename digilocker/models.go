package digilocker

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type AccountStatusRequest struct {
	Mobile  string `json:"mobile"`
	Aadhaar string `json:"aadhaar"`
}

type AccountStatusDetails struct {
	Code                 string `json:"code"`
	Mobile               string `json:"mobile"`
	Aadhaar              string `json:"aadhaar"`
	MobileAadhaarLinkage string `json:"mobileAadhaarLinkage"`
}

type AccountStatusResponse struct {
	Status     string               `json:"status"`
	StatusCode string               `json:"statusCode"`
	Result     AccountStatusDetails `json:"result"`
	Error      Error                `json:"error"`
}

type KYCStartRequest struct {
	ReferenceId string `json:"referenceId"`
	RedirectURL string `json:"redirectURL"`
}
type KYCStartDetails struct {
	URL string `json:"url"`
}
type KYCStartResponse struct {
	Status     string          `json:"status"`
	StatusCode string          `json:"statusCode"`
	Result     KYCStartDetails `json:"result"`
	Error      Error           `json:"error"`
}

type EAadhaarDetailsRequest struct {
	ReferenceId string `json:"referenceId"`
	AadhaarFile string `json:"aadhaarFile"`
}

type AadhaarDetails struct {
	Name                string `json:"name"`
	DOB                 string `json:"dob"`
	Address             string `json:"address"`
	MaskedAadhaarNumber string `json:"maskedAadhaarNumber"`
	Photo               string `json:"photo"`
	AadhaarFile         string `json:"aadhaarFile"`
	XMLAadhaarFile      string `json:"xmlAadhaarFile"`
}

type AadhaarDetailsResponse struct {
	Status     string         `json:"status"`
	StatusCode string         `json:"statusCode"`
	Result     AadhaarDetails `json:"result"`
	Error      Error          `json:"error"`
}

type HypervergeHealthcheckResponse struct {
	StatusCode string            `json:"statusCode"`
	Result     HealthcheckResult `json:"result"`
}

type HealthcheckResult struct {
	Message          string `json:"message"`
	Endpoint         string `json:"endpoint"`
	Severity         string `json:"severity"`
	MeanResponseTime string `json:"meanResponseTime"`
	UserErrors       string `json:"userErrors"`
}
