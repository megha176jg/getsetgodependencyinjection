package main

import (
	"context"

	"bitbucket.org/junglee_games/getsetgo/loggerv2"
)

func main() {

	loggerv2.Config{AppName: "test", Build: "dev"}.InitiateLogger()
	ctx := context.Background()
	loggerv2.AddMiddleLayers(loggerv2.RequestMiddleLayer)
	loggerv2.Info(ctx, "info level message")
	loggerv2.Warn(ctx, "warn message")
	loggerv2.Debug(ctx, "debug message")
	loggerv2.Error(ctx, "error message")
}
