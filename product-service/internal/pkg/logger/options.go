package custom_logger

import (
	"github.com/kingstonduy/go-core/logger"
	"github.com/sirupsen/logrus"
)

type LogrusOptions struct {
	logger.Options
	Formatter logrus.Formatter
	Hooks     logrus.LevelHooks
	// Flag for whether to log caller info (off by default)
	ReportCaller bool
	// Exit Function to call when FatalLevel log
	ExitFunc func(int)
}

type formatterKey struct{}

func WithTextTextFormatter(formatter *logrus.TextFormatter) logger.Option {
	return logger.SetOption(formatterKey{}, formatter)
}
func WithJSONFormatter(formatter *logrus.JSONFormatter) logger.Option {
	return logger.SetOption(formatterKey{}, formatter)
}
func WithFormatter(formatter logrus.Formatter) logger.Option {
	return logger.SetOption(formatterKey{}, formatter)
}

type hooksKey struct{}

func WithLevelHooks(hooks logrus.LevelHooks) logger.Option {
	return logger.SetOption(hooksKey{}, hooks)
}

// set hooks, override the LevelHooks set before
func WithHooks(hooks ...logrus.Hook) logger.Option {
	levelHooks := logrus.LevelHooks{}
	for _, h := range hooks {
		levelHooks.Add(h)
	}
	return logger.SetOption(hooksKey{}, levelHooks)
}

type reportCallerKey struct{}

// warning to use this option. because logrus doest not open CallerDepth option
// this will only print this package.
func ReportCaller() logger.Option {
	return logger.SetOption(reportCallerKey{}, true)
}

type exitKey struct{}

func WithExitFunc(exit func(int)) logger.Option {
	return logger.SetOption(exitKey{}, exit)
}

type logrusLoggerKey struct{}

func WithLogger(l logrus.StdLogger) logger.Option {
	return logger.SetOption(logrusLoggerKey{}, l)
}
