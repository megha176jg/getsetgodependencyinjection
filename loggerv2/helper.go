package loggerv2

import (
	"context"

	"go.uber.org/zap"
)

func executeMiddleLayers(ctx context.Context, msg string, fields []zap.Field) (context.Context, string, []zap.Field) {

	for _, f := range middleLayers {
		ctx, msg, fields = f(ctx, msg, fields)
	}
	return ctx, msg, fields
}
