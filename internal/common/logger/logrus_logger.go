package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

func init() {
	lv := os.Getenv("LOG_lEVEL")
	format := os.Getenv("LOG_FORMATTER")

	logrus.SetFormatter(formatter(format))
	logrus.SetLevel(level(lv))
}

func NewLogrusLogger() *logrus.Logger {
	lv := os.Getenv("LOG_lEVEL")
	format := os.Getenv("LOG_FORMATTER")

	logger := logrus.New()
	logger.SetFormatter(formatter(format))
	logger.SetLevel(level(lv))

	return logger
}

func level(lv string) logrus.Level {
	switch lv {
	case "trace":
		return logrus.TraceLevel
	case "debug":
		return logrus.DebugLevel
	case "info":
		return logrus.InfoLevel
	case "warn":
		return logrus.WarnLevel
	case "error":
		return logrus.ErrorLevel
	case "fatal":
		return logrus.FatalLevel
	case "panic":
		return logrus.PanicLevel
	default:
		return logrus.InfoLevel
	}
}

func formatter(format string) logrus.Formatter {
	switch format {
	case "json":
		return &logrus.JSONFormatter{}
	default:
		return &logrus.TextFormatter{}
	}
}