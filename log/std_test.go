package log

import (
	"log"
	"testing"
)

func TestAgeAt(t *testing.T) {
	std.Errorln("some error")
	std.Infoln("some info")
	std.Debugln("no debug output")

	log.Println("hijack stdlib log")
}
