package logger

import (
	"context"

	"bitbucket.org/junglee_games/getsetgo/logger/message"
)

var _ LogInterface = (*zapImpl)(nil)

type LogInterface interface {
	Trace(ctx context.Context, msg message.LogMsg)
	Warning(ctx context.Context, msg message.LogMsg)
	Info(ctx context.Context, msg message.LogMsg)
	Error(ctx context.Context, msg message.LogMsg)
	Debug(ctx context.Context, msg message.LogMsg)
	Profile(ctx context.Context, msg message.LogMsg)
}
