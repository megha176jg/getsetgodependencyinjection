package web

import (
	"mime/multipart"

	"github.com/kataras/iris/v12"
)

// Param returns the web call parameters from the request.
func Param(ctx iris.Context, key string) string {
	m := ctx.Params()
	return m.Get(key)
}

// Decode reads the body of an HTTP request looking for a JSON document. The
// body is decoded into the provided value.
//
// If the provided value is a struct then it is checked for validation tags.
func Decode(ctx iris.Context, val interface{}) error {
	if err := ctx.ReadJSON(val); err != nil {
		return err
	}

	return nil
}

// DecodeForm reads the form of an HTTP request. The
// form is decoded into the provided value.
//
// If the provided value is a struct then it is checked for validation tags.
func DecodeForm(ctx iris.Context, val interface{}) error {
	if err := ctx.ReadForm(val); err != nil {
		return err
	}

	return nil
}

// DecodeFormFile reads the form of an HTTP request to find type "file". The
// file is returned along with it's headers.
// The field is looked up in form for type file.
func DecodeFormFile(ctx iris.Context, field string) (multipart.File, *multipart.FileHeader, error) {
	file, fh, err := ctx.FormFile(field)
	if err != nil {
		return nil, nil, err
	}

	return file, fh, err
}
