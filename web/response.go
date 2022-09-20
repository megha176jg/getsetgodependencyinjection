package web

import (
	"net/http"

	"github.com/kataras/iris/v12"
	"github.com/pkg/errors"
)

// Respond converts a Go value to JSON and sends it to the client.
func Respond(ctx iris.Context, data interface{}, statusCode int) error {
	v, ok := ctx.Values().Get(KeyValues).(*Values)
	if !ok {
		ctx.Application().Logger().Println("Respond: web value missing from context")
		return errors.New("Respond: web value missing from context")
	}
	v.StatusCode = statusCode
	ctx.StatusCode(statusCode)
	// If there is nothing to marshal then set status code and return.
	if statusCode == http.StatusNoContent {
		return nil
	}
	ctx.JSON(data)
	ctx.StopExecution()
	return nil
}

// RespondInternalServerError responds with 500 and something went wrong message
func RespondInternalServerError(ctx iris.Context, data interface{}) error {
	return Respond(ctx, data, http.StatusInternalServerError)
}
