package loggerv2

import (
	"context"
	"fmt"

	"go.uber.org/zap/zapcore"
)

func executeMiddleLayers(ctx context.Context, msg string, fields *Fields) (context.Context, string, *Fields) {

	for _, f := range middleLayers {
		ctx, msg, fields = f(ctx, msg, fields)
	}
	return ctx, msg, fields
}

func getEnv(env string) Env {
	if env == "prod" {
		return PROD
	}
	return DEV
}

func setColour(level zapcore.Level) {
	if env == PROD {
		return
	}
	switch level {
	case zapcore.ErrorLevel:
		fmt.Print(redColor)
	case zapcore.WarnLevel:
		fmt.Print(yellowColor)
	case zapcore.DebugLevel:
		fmt.Print(greenColor)
	case zapcore.PanicLevel:
		fmt.Print(redColor)
	default:
		return
	}
}

func resetColour() {
	if env == PROD {
		return
	}
	fmt.Print(defaultStyle)
}
