package loggerv2

import (
	"context"

	"bitbucket.org/junglee_games/getsetgo/utils"
)

func RequestMiddleLayer(ctx context.Context, msg string, fields *Fields) (context.Context, string, *Fields) {
	vRc, rcOk := ctx.Value(utils.RequestContextKey).(*utils.RequestContext)
	if !rcOk {
		return ctx, msg, fields
	}
	fields.AddField(requestIDKey, vRc.RequestID)
	fields.AddField(appIDKey, vRc.ClientAppID)
	fields.AddField(userIDKey, vRc.UserID)
	fields.AddField(uriKey, vRc.URI)
	fields.AddField(ipKey, vRc.IP)
	return ctx, msg, fields
}
