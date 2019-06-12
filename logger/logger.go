package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// MicroLogger is a logger wrapper of logrus for Go Micro.
// It implements go-log Logger interface.
type Logger struct {
	*logrus.Logger
}

// NewLogger creates a new MicroLogger.
func NewLogger() *Logger {
	var log = &logrus.Logger{
		Out:          os.Stderr,
		Formatter:    DefaultTextFormatter(),
		Hooks:        make(logrus.LevelHooks),
		Level:        logrus.InfoLevel,
		ExitFunc:     os.Exit,
		ReportCaller: false,
	}

	// 默认值参考 logrus.Logger 设定
	// https://godoc.org/github.com/sirupsen/logrus#Logger
	return &Logger{
		log,
	}
}

// Log implements go-log Log interface
func (l *Logger) Log(v ...interface{}) {
	l.Print(v...)
}

// Logf implements go-log Logf interface
func (l *Logger) Logf(format string, v ...interface{}) {
	l.Printf(format, v...)
}

// Log makes use of github.com/go-log/log.Log
func Log(v ...interface{}) {
	std.Log(v...)
}

// Logf makes use of github.com/go-log/log.Logf
func Logf(format string, v ...interface{}) {
	std.Logf(format, v...)
}

// Fatal logs with Log and then exits with os.Exit(1)
func Fatal(v ...interface{}) {
	std.Fatal(v...)
}

// Fatalf logs with Logf and then exits with os.Exit(1)
func Fatalf(format string, v ...interface{}) {
	std.Fatalf(format, v...)
}
