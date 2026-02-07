package logger

import (
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	instance *logrus.Logger
	once     sync.Once
)

// GetLogger returns the singleton logger instance
func GetLogger() *logrus.Logger {
	once.Do(func() {
		instance = logrus.New()

		// Set output to stdout
		instance.SetOutput(os.Stdout)

		// Set log format
		instance.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})

		// Set log level from environment variable or default to Info
		level := os.Getenv("LOG_LEVEL")
		switch level {
		case "debug":
			instance.SetLevel(logrus.DebugLevel)
		case "warn":
			instance.SetLevel(logrus.WarnLevel)
		case "error":
			instance.SetLevel(logrus.ErrorLevel)
		case "fatal":
			instance.SetLevel(logrus.FatalLevel)
		case "panic":
			instance.SetLevel(logrus.PanicLevel)
		case "trace":
			instance.SetLevel(logrus.TraceLevel)
		default:
			instance.SetLevel(logrus.InfoLevel)
		}

		// Add hooks if needed (optional)
		// instance.AddHook(...)
	})
	return instance
}

func Info(args ...interface{}) {
	GetLogger().Info(args...)
}

func Infof(format string, args ...interface{}) {
	GetLogger().Infof(format, args...)
}

func Debug(args ...interface{}) {
	GetLogger().Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	GetLogger().Debugf(format, args...)
}

func Warn(args ...interface{}) {
	GetLogger().Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	GetLogger().Warnf(format, args...)
}

func Error(args ...interface{}) {
	GetLogger().Error(args...)
}

func Errorf(format string, args ...interface{}) {
	GetLogger().Errorf(format, args...)
}

func Fatal(args ...interface{}) {
	GetLogger().Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	GetLogger().Fatalf(format, args...)
}

func WithFields(fields logrus.Fields) *logrus.Entry {
	return GetLogger().WithFields(fields)
}

func WithField(key string, value interface{}) *logrus.Entry {
	return GetLogger().WithField(key, value)
}
