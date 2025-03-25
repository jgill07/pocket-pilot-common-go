// Logger is a wrapper around the zap logger
package log

import "go.uber.org/zap"

var logger *zap.Logger = safeInitLogger()

func safeInitLogger() *zap.Logger {
	l, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	return l
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}

func Sync() error {
	return logger.Sync()
}

func WithField(key string, value interface{}) *zap.Logger {
	return logger.With(zap.Any(key, value))
}

func WithFields(fields ...zap.Field) *zap.Logger {
	return logger.With(fields...)
}

func WithError(err error) *zap.Logger {
	return logger.With(zap.Error(err))
}
