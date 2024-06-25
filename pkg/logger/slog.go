package logger

import (
	"fmt"
	"log/slog"
)

type sLogLogger struct {
	logLevel slog.Level
	instance *slog.Logger
}

func New() Logger {
	return &sLogLogger{}
}

func (slogLogger *sLogLogger) Debug(msg string, any ...interface{}) {
	slogLogger.instance.Debug(msg, any...)
}

func (slogLogger *sLogLogger) Info(msg string, any ...interface{}) {
	slogLogger.instance.Info(msg, any...)
}

func (slogLogger *sLogLogger) Warn(msg string, any ...interface{}) {
	slogLogger.instance.Warn(msg, any...)
}

func (slogLogger *sLogLogger) Error(msg string, any ...interface{}) {
	slogLogger.instance.Error(msg, any...)
}

func (slogLogger *sLogLogger) WithAttribute(attribute Attribute) Logger {
	newInstance := &sLogLogger{
		instance: slogLogger.instance.With(slog.Attr{
			Key:   attribute.Key,
			Value: slog.AnyValue(attribute.Value),
		}),
	}

	return newInstance
}

func (slogLogger *sLogLogger) Setup(config *Config) error {

	switch config.level {
	case "info":
		slogLogger.logLevel = slog.LevelInfo
	case "error":
		slogLogger.logLevel = slog.LevelError
	case "debug":
		slogLogger.logLevel = slog.LevelDebug
	case "warn":
		slogLogger.logLevel = slog.LevelWarn
	default:
		return fmt.Errorf("unknown logger level: %s", config.level)
	}

	// Set logging level option
	var handler slog.Handler
	opts := &slog.HandlerOptions{
		Level: slogLogger.logLevel,
	}

	if !config.withTimestamp {
		opts.ReplaceAttr = func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == "time" {
				a.Value = slog.Value{}
				a.Key = ""
			}
			return a
		}
	}

	switch config.outType {
	case "json":
		handler = slog.NewJSONHandler(config.writer, opts)
	case "text":
		handler = slog.NewTextHandler(config.writer, opts)
	default:
		return fmt.Errorf("unknown logger type: %s", config.outType)
	}

	instance := slog.New(handler)

	slogLogger.instance = instance
	return nil
}
