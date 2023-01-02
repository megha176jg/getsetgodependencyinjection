package logger

import (
	"context"
)

var _ LogInterface = (*zapImpl)(nil)

type LogInterface interface {
	Trace(ctx context.Context, a ...interface{})
	Warning(ctx context.Context, a ...interface{})
	Info(ctx context.Context, a ...interface{})
	Error(ctx context.Context, a ...interface{})
	Debug(ctx context.Context, a ...interface{})
	Profile(ctx context.Context, a ...interface{})
}
