package idfy

type Idfy interface {
	ExtractPan(documentType string, idfyrequest IdfyRequest) (*IdfyPanResponse, error)
	ExtractAadhar(documentType string, idfyrequest IdfyRequest) (*IdfyAadharResponse, error)
	ExtractDl(documentType string, idfyrequest IdfyRequest) (*IdfyDlResponse, error)
	ExtractVoter(documentType string, idfyrequest IdfyRequest) (*IdfyVoterIdResponse, error)
	ExtractPassport(documentType string, idfyrequest IdfyRequest) (*IdfyPassportResponse, error)
}

type IdfyConfig interface {
	GetIdfyAccountId() string
	GetIdfyApiKey() string
	GetIdfyEndpoint() string
}
