package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.Logger
}

func NewLogger(debug bool) *Logger {
	var config zap.Config

	if debug {
		// Development config (human-readable, debug level)
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		// Production config (JSON format, info level)
		config = zap.NewProductionConfig()
	}

	config.OutputPaths = []string{"stdout"}

	logger, err := config.Build()
	if err != nil {
		panic("failed to create logger: " + err.Error())
	}

	return &Logger{logger}
}

func (l *Logger) Sync() {
	_ = l.Logger.Sync()
}

func (l *Logger) WithField(key string, value interface{}) *Logger {
	return &Logger{l.Logger.With(zap.Any(key, value))}
}

func (l *Logger) WithError(err error) *Logger {
	return &Logger{l.Logger.With(zap.Error(err))}
}
