package log

var (
	// std is a shared standard Logger
	std *Logger
)

func init() {
	std = NewLogger()
}

// StandardLogger returns a shared standard Logger.
func StandardLogger() *Logger {
	return std
}
