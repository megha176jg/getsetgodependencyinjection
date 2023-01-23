package common_errors

import "fmt"

var (
	ErrS3FileBodyReader = fmt.Errorf("error reading S3 file body")
	ErrS3FileDownload   = fmt.Errorf("error downloading s3 file")
	ErrS3PresignFailed  = fmt.Errorf("error presigning file")
)
