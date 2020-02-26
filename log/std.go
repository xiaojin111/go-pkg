package log

import (
	"github.com/sirupsen/logrus"
)

var (
	// std is a shared standard Logger
	std *Logger
)

// StandardLogger returns a shared standard Logger.
func StandardLogger() *Logger {
	return std
}

func GetLevel() logrus.Level {
	return std.GetLevel()
}

func init() {
	std = NewLogger()
	Hijack()
}
