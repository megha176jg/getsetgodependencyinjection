package onfido

import "errors"

var (
	ErrRequestCreation         = errors.New("ERROR : ONFIDO SDK : creating request for onfido")
	ErrCallingOnfido           = errors.New("ERROR : ONFIDO SDK : calling onfido")
	ErrReadingResponseBody     = errors.New("ERROR : ONFIDO SDK : reading response body")
	ErrUnmarshalingResponse    = errors.New("ERROR : ONFIDO SDK : unmarshaling response")
	ErrReadingFile             = errors.New("ERROR : ONFIDO SDK : reading upload file")
	ErrWritingFormField        = errors.New("ERROR : ONFIDO SDK : writing form field")
	ErrFileOrDirectoryNotFound = errors.New("ERROR : ONFIDO SDK : file or directory not found")
	ErrStatusCodeOtherThan2XX  = errors.New("ERROR : ONFIDO SDK : status code other than 2XX")
)
