package loggerv2

import (
	"context"

	"bitbucket.org/junglee_games/getsetgo/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func RequestMiddleLayer(ctx context.Context, msg string, fields []zap.Field) (context.Context, string, []zap.Field) {
	vRc, rcOk := ctx.Value("RequestContext").(*utils.RequestContext)
	if !rcOk {
		return ctx, msg, fields
	}
	fields = append(fields, zap.Field{Key: requestIDKey, Type: zapcore.StringType, String: vRc.RequestID})
	fields = append(fields, zap.Field{Key: appIDKey, Type: zapcore.StringType, String: vRc.ClientAppID})
	fields = append(fields, zap.Field{Key: userIDKey, Type: zapcore.StringType, String: vRc.UserID})
	fields = append(fields, zap.Field{Key: uriKey, Type: zapcore.StringType, String: vRc.URI})
	fields = append(fields, zap.Field{Key: ipKey, Type: zapcore.StringType, String: vRc.IP})

	return ctx, msg, fields
}
