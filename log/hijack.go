package log

import (
	microlog "github.com/micro/go-micro/util/log"
)

func Hijack() {
	// 将 std 设定到 go-micro 的 logger 之中
	microlog.SetLogger(std)
}
