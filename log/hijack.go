package log

import (
	microlog "github.com/micro/go-micro/util/log"
	"github.com/sirupsen/logrus"
)

func Hijack() {
	// 将 std 设定到 go-micro 的 logger 之中
	microlog.SetLogger(std)
}

func hijackLevel(level logrus.Level) {
	l := microlog.LevelInfo
	switch level {
	case logrus.TraceLevel:
		l = microlog.LevelTrace
	case logrus.DebugLevel:
		l = microlog.LevelDebug
	case logrus.InfoLevel:
		l = microlog.LevelInfo
	case logrus.WarnLevel:
		l = microlog.LevelWarn
	case logrus.ErrorLevel:
		l = microlog.LevelWarn
	case logrus.FatalLevel:
		l = microlog.LevelFatal
	case logrus.PanicLevel:
		l = microlog.LevelFatal
	}
	microlog.SetLevel(l)
}
