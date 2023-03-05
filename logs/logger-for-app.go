package logs

import (
	"go.uber.org/zap"
)

var (
	configForApp zap.Config
	atomForApp   zap.AtomicLevel
	loggerForApp *zap.Logger
)

func init() {
	var err error
	atomForApp = zap.NewAtomicLevel()
	configForApp = zap.NewDevelopmentConfig()
	configForApp.Level = atomForApp
	loggerForApp, err = configForApp.Build(zap.AddCallerSkip(3))
	if err != nil {
		panic(err)
	}
}

func SetCallerSkipForApp(skip int) {
	var err error
	loggerForApp, err = configForApp.Build(zap.AddCallerSkip(skip))
	if err != nil {
		panic(err)
	}
}

func GetZapLoggerForApp() *zap.Logger {
	return loggerForApp
}

// func SetLevel(level zapcore.Level) {
// 	atomForApp.SetLevel(level)
// }

// func Debug(msg string, fields ...zap.Field) {
// 	loggerForApp.Debug(msg, fields...)
// }

// func Info(msg string, fields ...zap.Field) {
// 	loggerForApp.Info(msg, fields...)
// }

// func Warn(msg string, fields ...zap.Field) {
// 	loggerForApp.Warn(msg, fields...)
// }

// func Error(msg string, fields ...zap.Field) {
// 	loggerForApp.Error(msg, fields...)
// }
