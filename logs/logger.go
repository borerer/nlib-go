package logs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	atom   zap.AtomicLevel
	logger *zap.Logger
)

func init() {
	var err error
	atom = zap.NewAtomicLevel()
	cfg := zap.NewDevelopmentConfig()
	cfg.Level = atom
	logger, err = cfg.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

func GetZapLogger() *zap.Logger {
	return logger
}

func SetLevel(level zapcore.Level) {
	atom.SetLevel(level)
}

func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}
