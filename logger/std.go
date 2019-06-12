package logger

var (
	// std is a shared standard Logger
	std *Logger
)

// StandardLogger returns a shared standard Logger.
func StandardLogger() *Logger {
	return std
}

func init() {
	std = NewLogger()
	Hijack()
}
