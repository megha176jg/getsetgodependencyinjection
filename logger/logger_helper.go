package logger

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"strings"

	"bitbucket.org/junglee_games/getsetgo/logger/message"
	"bitbucket.org/junglee_games/getsetgo/utils"
)

func InfoSpecific(ctx context.Context, logType string, a ...interface{}) {

	loggerHandle, err := GetLoggerHandle(ctx, logType)
	if err != nil {
		return
	}
	msg := Convert(ctx, a...)
	loggerHandle.Info(ctx, msg)
}

func Info(ctx context.Context, a ...interface{}) {
	InfoSpecific(ctx, defaultLoggerType, a...)
}

func DebugSpecific(ctx context.Context, logType string, a ...interface{}) {

	loggerHandle, err := GetLoggerHandle(ctx, logType)
	if err != nil {
		return
	}
	msg := Convert(ctx, a...)
	loggerHandle.Debug(ctx, msg)
}

func Debug(ctx context.Context, a ...interface{}) {
	DebugSpecific(ctx, defaultLoggerType, a...)
}

func ErrorSpecific(ctx context.Context, logType string, a ...interface{}) {

	loggerHandle, err := GetLoggerHandle(ctx, logType)
	if err != nil {
		return
	}
	msg := Convert(ctx, a...)
	loggerHandle.Error(ctx, msg)
}

func Error(ctx context.Context, a ...interface{}) {
	ErrorSpecific(ctx, defaultLoggerType, a...)
}

func WarningSpecific(ctx context.Context, logType string, a ...interface{}) {

	loggerHandle, err := GetLoggerHandle(ctx, logType)
	if err != nil {
		return
	}
	msg := Convert(ctx, a...)
	loggerHandle.Warning(ctx, msg)
}

func Warning(ctx context.Context, a ...interface{}) {
	WarningSpecific(ctx, defaultLoggerType, a...)
}

// Convert converts application log to LogMsg format suitable for Logger
func Convert(ctx context.Context, a ...interface{}) message.LogMsg {
	paramLength := len(a)
	if paramLength == 0 {
		return message.LogMsg{
			Message: "Empty log param",
		}
	}
	if paramLength == 1 {
		//Only Log Message string is passed
		return message.LogMsg{
			Message: fmt.Sprintf("%s", a[0]),
		}
	}

	//First param is message string; Second param is request context
	vMsg, msgOk := a[0].(string)
	if !msgOk {
		return message.LogMsg{
			Message: fmt.Sprintf("Erorr in parsing logging params for %v", a),
		}
	}

	vRc, rcOk := ctx.Value("RequestContext").(utils.RequestContext)
	if !rcOk {
		return message.LogMsg{
			Message: fmt.Sprintf("Erorr in parsing logging params for %v", a),
		}
	}
	return message.LogMsg{
		Message:       vMsg,
		TransactionID: vRc.TransactionID,
		SessionID:     vRc.SessionID,
		RequestID:     vRc.RequestID,
		AppID:         vRc.ClientAppID,
		UserID:        vRc.UserID,
		URI:           fmt.Sprintf("%s %s", strings.ToUpper(vRc.Method), vRc.URI),
		IP:            vRc.IP,
	}
}

func GetLoggerHandle(ctx context.Context, logType string) (LogInterface, error) {
	loggerHandle, ok := loggerImpls[logType]
	if !ok {
		return nil, errors.New("Undefined log type requested " + logType)
	}
	return loggerHandle, nil
}

// getStackTrace gets the stack trace for a called function.
func getStackTrace() []string {
	var sf []string
	j := 0
	for i := Skip; ; i++ {
		_, filePath, lineNumber, ok := runtime.Caller(i)
		if !ok || j >= CallingDepth {
			break
		}
		sf = append(sf, fmt.Sprintf("%s(%d)", filePath, lineNumber))
		j++
	}
	return sf
}
