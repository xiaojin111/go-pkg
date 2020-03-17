package log

import (
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

var (
	// std is a shared standard Logger
	std *Logger
)

func init() {
	std = newStdLogger()

	Hijack(std)
}

func Hijack(l *Logger) {
	// 注入标准包 log 输出
	log.SetOutput(l.Writer())
	log.SetFlags(0)
}

func newStdLogger() *Logger {
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

// StandardLogger returns a shared standard Logger.
func StandardLogger() *Logger {
	return std
}
