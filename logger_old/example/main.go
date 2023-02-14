package main

import (
	"context"

	logger "bitbucket.org/junglee_games/getsetgo/logger_old"
)

func main() {
	ctx := context.Background()
	lConfig := logger.Config{
		LoggerService: "zap",
		Level:         5,
		EncoderConfig: logger.EncoderConfig{},
		AppName:       "testing",
	}
	l, err := lConfig.InitiateLogger()
	if err != nil {
		panic(err)
	}
	l.Debug(ctx, "testing")
}
