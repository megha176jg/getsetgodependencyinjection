package loggerv2

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	zapLogger    *zap.Logger
	appNameField zap.Field
	middleLayers []MiddleLayer
	env          Env
)

func init() {
	Config{AppName: "default", Env: "dev"}.InitiateLogger()
	appNameField = zap.Field{Key: "App", Type: zapcore.StringType, String: "default"}
	middleLayers = make([]MiddleLayer, 0)
}

func (c Config) InitiateLogger() error {
	var err error
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = timeKey
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	var zapConfig zap.Config
	env = getEnv(c.Env)
	if env == PROD {
		zapConfig = zap.NewProductionConfig()
	} else {
		zapConfig = zap.NewDevelopmentConfig()
	}
	zapConfig.DisableStacktrace = true

	zapConfig.EncoderConfig = encoderConfig
	zapLogger, err = zapConfig.Build(zap.AddStacktrace(zapcore.ErrorLevel))

	appNameField.String = c.AppName
	return err
}

func AddMiddleLayers(middlelayers ...MiddleLayer) {
	middleLayers = append(middleLayers, middlelayers...)
}

func Info(ctx context.Context, format string, a ...any) {
	_, msg, fields := executeMiddleLayers(ctx, fmt.Sprintf(format, a...), &Fields{fields: []zap.Field{appNameField}})
	zapLogger.Info(msg, fields.fields...)
}

func Infow(ctx context.Context, message string, fs *Fields) {
	fs.fields = append(fs.fields, appNameField)
	_, msg, fields := executeMiddleLayers(ctx, message, fs)
	zapLogger.Info(msg, fields.fields...)
}

func Error(ctx context.Context, format string, a ...any) {
	setColour(zapcore.ErrorLevel)
	_, msg, fields := executeMiddleLayers(ctx, fmt.Sprintf(format, a...), &Fields{fields: []zap.Field{appNameField}})
	zapLogger.Error(msg, fields.fields...)
	resetColour()
}

func Warn(ctx context.Context, format string, a ...any) {
	setColour(zapcore.WarnLevel)
	_, msg, fields := executeMiddleLayers(ctx, fmt.Sprintf(format, a...), &Fields{fields: []zap.Field{appNameField}})
	zapLogger.Warn(msg, fields.fields...)
	resetColour()
}

func Debug(ctx context.Context, format string, a ...any) {
	setColour(zapcore.DebugLevel)
	_, msg, fields := executeMiddleLayers(ctx, fmt.Sprintf(format, a...), &Fields{fields: []zap.Field{appNameField}})
	zapLogger.Debug(msg, fields.fields...)
	resetColour()
}

func Panic(ctx context.Context, format string, a ...any) {
	setColour(zap.PanicLevel)
	_, msg, fields := executeMiddleLayers(ctx, fmt.Sprintf(format, a...), &Fields{fields: []zap.Field{appNameField}})
	zapLogger.Panic(msg, fields.fields...)
	resetColour()
}

func Fatal(ctx context.Context, format string, a ...any) {
	setColour(zap.FatalLevel)
	_, msg, fields := executeMiddleLayers(ctx, fmt.Sprintf(format, a...), &Fields{fields: []zap.Field{appNameField}})
	zapLogger.Fatal(msg, fields.fields...)
	resetColour()
}
