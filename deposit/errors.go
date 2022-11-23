package deposit

import "errors"

var (
	ErrCreatingRequest             = errors.New("ErrCreatingRequest")
	ErrCallingHouzatPaymentService = errors.New("ErrCallingHouzatPaymentService")
	ErrReadingBody                 = errors.New("ErrReadingBody")
	ErrUnmarshlingResponse         = errors.New("ErrUnmarshlingResponse")
	ErrStatusCodeOtherThan200      = errors.New("ErrStatusCodeOtherThan200")
)
