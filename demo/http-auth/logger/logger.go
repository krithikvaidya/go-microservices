package logger

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var zapLog *zap.Logger

func init() {

	var err error
	zap.NewProduction()
	config := zap.NewProductionConfig()

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.StacktraceKey = ""

	config.EncoderConfig = encoderConfig

	zapLog, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		log.Fatal("Not able to initialize logger....")
	}
}

func Info(msg string, fields ...zapcore.Field) {
	zapLog.Info(msg, fields...)
}

func Debug(msg string, fields ...zapcore.Field) {
	zapLog.Debug(msg, fields...)
}

func Fatal(msg string, fields ...zapcore.Field) {
	zapLog.Fatal(msg, fields...)
}

func Error(msg string, fields ...zapcore.Field) {
	zapLog.Error(msg, fields...)
}
