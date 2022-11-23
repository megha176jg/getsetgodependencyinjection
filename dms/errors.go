package dms

import "errors"

var (
	ErrFileNotFound        = errors.New("ErrFileNotFound")
	ErrCreatingRequest     = errors.New("ErrCreatingRequest")
	ErrCallingDMS          = errors.New("ErrCallingDMS")
	ErrReadingResponseBody = errors.New("ErrReadingResponseBody")
)
