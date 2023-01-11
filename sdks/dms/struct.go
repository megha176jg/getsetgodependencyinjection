package dms

type IntiateRequest struct {
	UserID     int
	DocType    string
	FrontImage string
	BackImage  string
	XProductID string
}

type IntiateResponse struct {
	ID string
}
