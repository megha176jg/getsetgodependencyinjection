package loggerv2

import (
	"context"

	"go.uber.org/zap"
)

type MiddleLayer func(ctx context.Context, msg string, fields []zap.Field) (context.Context, string, []zap.Field)

// Build : if prod it will set to prod else dev
type Config struct {
	AppName string
	Build   string
}
