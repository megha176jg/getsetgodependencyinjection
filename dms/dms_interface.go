package dms

type DMS interface {
	Initiate(req IntiateRequest) (IntiateResponse, error)
}
