package profile

import "errors"

var (
	ErrCreatingRequest        = errors.New("ErrCreatingRequest")
	ErrCallingProfile         = errors.New("ErrCallingProfile")
	ErrReadingBody            = errors.New("ErrReadingBody")
	ErrUnmarshlingResponse    = errors.New("ErrUnmarshlingResponse")
	ErrStatusCodeOtherThan200 = errors.New("ErrStatusCodeOtherThan200")
)
