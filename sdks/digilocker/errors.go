package digilocker

import "errors"

var (
	ErrNotAvailable      = errors.New("ERROR : hyperverge not available")
	ErrUnmarshalJson     = errors.New("ERROR : unmarshling response")
	ErrReadResponseBody  = errors.New("ERROR : reading response body")
	ErrCallingHyperverge = errors.New("ERROR : calling hyperverge")
	ErrCreatingRequest   = errors.New("ERROR : creating request")
	ErrReqValidate       = errors.New("ERROR : hyperverge request validate")
	ErrHVServer          = errors.New("ERROR : Hyperverge Server")
)
