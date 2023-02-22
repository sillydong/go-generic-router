package logx

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Level zapcore.Level

const (
	DebugLevel Level = iota - 1
	InfoLevel
	WarnLevel
	ErrorLevel
)

var logger *zap.Logger

func init() {
	New("demo", InfoLevel)
}

func New(app string, level Level) {
	conf := zap.NewProductionConfig()
	conf.Level = zap.NewAtomicLevelAt(zapcore.Level(level))
	if level == DebugLevel {
		conf = zap.NewDevelopmentConfig()
	}
	conf.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	conf.InitialFields = map[string]interface{}{
		"app": app,
	}
	base, err := conf.Build(zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.DPanicLevel))
	if err != nil {
		panic(err)
	}

	logger = base
}

func Clone() *zap.Logger {
	return logger.WithOptions(zap.AddCallerSkip(-1))
}

func Sync() {
	logger.Sync()
}

func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	logger.Panic(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}
