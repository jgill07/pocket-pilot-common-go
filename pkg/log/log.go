// Logger is a wrapper around the zap logger
package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger = safeInitLogger()

func safeInitLogger() *zap.Logger {
	l, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	return l
}

func InitLogger(level string) {
	lvl, err := zap.ParseAtomicLevel(level)
	if err != nil {
		panic(err)
	}

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderCfg.LevelKey = "level"
	encoderCfg.MessageKey = "msg"
	encoderCfg.CallerKey = "caller"
	encoderCfg.StacktraceKey = "stacktrace"
	encoderCfg.LineEnding = zapcore.DefaultLineEnding
	encoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder

	cfg := zap.Config{
		Level:             lvl,
		Development:       false,
		Encoding:          "json",
		DisableCaller:     false,
		DisableStacktrace: false,
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stderr"},
		EncoderConfig:     encoderCfg,
		InitialFields: map[string]interface{}{
			"pid": os.Getpid(),
		},
	}
	logger = zap.Must(cfg.Build())
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
