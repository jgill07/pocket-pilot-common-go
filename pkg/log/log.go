// Logger is a wrapper around the zap logger
package log

import "go.uber.org/zap"

type Logger struct {
	zap *zap.Logger
}

func NewLogger() (*Logger, error) {
	l, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	return &Logger{zap: l}, nil
}

func (l *Logger) Info(msg string, fields ...zap.Field) {
	l.zap.Info(msg, fields...)
}

func (l *Logger) Debug(msg string, fields ...zap.Field) {
	l.zap.Debug(msg, fields...)
}

func (l *Logger) Error(msg string, fields ...zap.Field) {
	l.zap.Error(msg, fields...)
}

func (l *Logger) Fatal(msg string, fields ...zap.Field) {
	l.zap.Fatal(msg, fields...)
}

func (l *Logger) Sync() error {
	return l.zap.Sync()
}

func (l *Logger) WithField(key string, value interface{}) *Logger {
	return &Logger{zap: l.zap.With(zap.Any(key, value))}
}

func (l *Logger) WithFields(fields ...zap.Field) *Logger {
	return &Logger{zap: l.zap.With(fields...)}
}

func (l *Logger) WithError(err error) *Logger {
	return &Logger{zap: l.zap.With(zap.Error(err))}
}
