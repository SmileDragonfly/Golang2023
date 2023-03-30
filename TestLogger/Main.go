package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"time"
)

func main() {
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.EncoderConfig.TimeKey = "timestamp"
	loggerConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	logger, err := loggerConfig.Build()
	if err != nil {
		log.Fatal(err)
	}
	sugar := logger.Sugar()
	sugar.Info("Hello from zap logger")
}
