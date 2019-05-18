package idgen

import (
	"github.com/rs/xid"
)

// NewXid 生成一个新的 XID
func NewXid() string {
	guid := xid.New()
	return guid.String()
}
