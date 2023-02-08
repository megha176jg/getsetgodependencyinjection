package main

import (
	"context"

	"bitbucket.org/junglee_games/getsetgo/loggerv2"
	"bitbucket.org/junglee_games/getsetgo/utils"
)

func main() {

	loggerv2.Config{AppName: "Test App", Build: "dev"}.InitiateLogger()
	ctx := context.Background()
	loggerv2.AddMiddleLayers(loggerv2.RequestMiddleLayer)
	ctx = context.WithValue(ctx, utils.RequestContextKey, &utils.RequestContext{RequestID: "134234", UserID: "456", IP: "196.1.100.23", URI: "/test/ping", ClientAppID: "RUMMY"})
	loggerv2.Info(ctx, "info level message")
	loggerv2.Warn(ctx, "warn message")
	loggerv2.Debug(ctx, "debug message")
	loggerv2.Error(ctx, "error message")
	fields := &loggerv2.Fields{}
	fields.AddField("time_taken", 49856)
	loggerv2.Infow(ctx, "Request end", fields)
}
