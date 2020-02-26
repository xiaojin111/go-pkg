package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Logger is logrus.Logger.
type Logger = logrus.Logger

// NewLogger creates a new MicroLogger.
func NewLogger() *Logger {
	var l = &logrus.Logger{
		Out:          os.Stderr,
		Formatter:    DefaultTextFormatter(),
		Hooks:        make(logrus.LevelHooks),
		Level:        logrus.InfoLevel,
		ExitFunc:     os.Exit,
		ReportCaller: false,
	}

	// 默认值参考 logrus.Logger 设定
	// https://godoc.org/github.com/sirupsen/logrus#Logger
	return l
}

// Log makes use of github.com/go-log/log.Log
func Log(v ...interface{}) {
	std.Log(std.GetLevel(), v...)
}

// Logf makes use of github.com/go-log/log.Logf
func Logf(format string, v ...interface{}) {
	std.Logf(std.GetLevel(), format, v...)
}

// Fatal logs with Log and then exits with os.Exit(1)
func Fatal(v ...interface{}) {
	std.Fatal(v...)
}

// Fatalf logs with Logf and then exits with os.Exit(1)
func Fatalf(format string, v ...interface{}) {
	std.Fatalf(format, v...)
}
