package logger

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapImpl struct {
	zap         *zap.Logger
	PushMetrics func(ctx context.Context, msg string) error
}

func (this *zapImpl) Trace(ctx context.Context, a ...interface{}) {

	msg := Convert(ctx, a...)
	fields := make([]zap.Field, 0)
	for _, eachField := range logFields {
		var value string
		switch eachField {
		case "URI":
			value = msg.URI
		case "RequestID":
			value = msg.RequestID
		case "IP":
			value = msg.IP
		case "TimeTaken":
			if msg.TimeTaken == "" {
				continue
			}
			value = msg.TimeTaken
		case "Caller":
			trace := getStackTrace()
			value = fmt.Sprintf("%s", trace)
		}
		fields = append(fields, zap.Field{
			Key:    eachField,
			Type:   zapcore.StringType,
			String: value,
		})
	}
	if this.PushMetrics != nil {
		this.PushMetrics(ctx, msg.Message)
	}
	this.zap.Error(msg.Message, fields...)
}

func (this *zapImpl) Warning(ctx context.Context, a ...interface{}) {

	msg := Convert(ctx, a...)

	fields := make([]zap.Field, 0)
	for _, eachField := range logFields {
		var value string
		switch eachField {
		case "URI":
			value = msg.URI
		case "RequestID":
			value = msg.RequestID
		case "IP":
			value = msg.IP
		case "TimeTaken":
			if msg.TimeTaken == "" {
				continue
			}
			value = msg.TimeTaken
		case "Caller":
			trace := getStackTrace()
			value = fmt.Sprintf("%s", trace[0])
		}
		fields = append(fields, zap.Field{
			Key:    eachField,
			Type:   zapcore.StringType,
			String: value,
		})
	}

	if this.PushMetrics != nil {
		this.PushMetrics(ctx, msg.Message)
	}

	this.zap.Warn(msg.Message, fields...)
}

func (this *zapImpl) Info(ctx context.Context, a ...interface{}) {

	msg := Convert(ctx, a...)
	fields := make([]zap.Field, 0)
	for _, eachField := range logFields {
		var value string
		switch eachField {
		case "URI":
			value = msg.URI
		case "RequestID":
			value = msg.RequestID
		case "IP":
			value = msg.IP
		case "TimeTaken":
			if msg.TimeTaken == "" {
				continue
			}
			value = msg.TimeTaken
		case "Caller":
			trace := getStackTrace()
			value = fmt.Sprintf("%s", trace[0])
		}
		fields = append(fields, zap.Field{
			Key:    eachField,
			Type:   zapcore.StringType,
			String: value,
		})
	}
	if this.PushMetrics != nil {
		this.PushMetrics(ctx, msg.Message)
	}

	this.zap.Info(msg.Message, fields...)
}

func (this *zapImpl) Error(ctx context.Context, a ...interface{}) {

	msg := Convert(ctx, a...)
	fields := make([]zap.Field, 0)
	for _, eachField := range logFields {
		var value string
		switch eachField {
		case "URI":
			value = msg.URI
		case "RequestID":
			value = msg.RequestID
		case "IP":
			value = msg.IP
		case "TimeTaken":
			if msg.TimeTaken == "" {
				continue
			}
			value = msg.TimeTaken
		case "Caller":
			trace := getStackTrace()
			value = fmt.Sprintf("%s", trace)
		}
		fields = append(fields, zap.Field{
			Key:    eachField,
			Type:   zapcore.StringType,
			String: value,
		})
	}

	if this.PushMetrics != nil {
		this.PushMetrics(ctx, msg.Message)
	}
	this.zap.Error(msg.Message, fields...)
}

func (this *zapImpl) Debug(ctx context.Context, a ...interface{}) {

	msg := Convert(ctx, a...)
	fields := make([]zap.Field, 0)
	for _, eachField := range logFields {
		var value string
		switch eachField {
		case "URI":
			value = msg.URI
		case "RequestID":
			value = msg.RequestID
		case "IP":
			value = msg.IP
		case "TimeTaken":
			if msg.TimeTaken == "" {
				continue
			}
			value = msg.TimeTaken
		case "Caller":
			trace := getStackTrace()
			value = fmt.Sprintf("%s", trace[0])
		}
		fields = append(fields, zap.Field{
			Key:    eachField,
			Type:   zapcore.StringType,
			String: value,
		})
	}

	if this.PushMetrics != nil {
		this.PushMetrics(ctx, msg.Message)
	}

	this.zap.Debug(msg.Message, fields...)
}

func (this *zapImpl) Profile(ctx context.Context, a ...interface{}) {

	//todo
}

func (this *zapImpl) getTime(time string) string {
	return fmt.Sprintf("%s %s %s", pinkColor, time, defaultStyle)
}

func (this *zapImpl) getURI(uri string) string {
	return fmt.Sprintf("%s %s %s", lightBlueColor, uri, defaultStyle)
}
