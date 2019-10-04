package idgen

import (
	"github.com/google/uuid"
	"github.com/rs/xid"
)

// NewXid 生成一个新的 XID
func NewXid() string {
	guid := xid.New()
	return guid.String()
}

// NewSecureRandomUUID 生成一个随机的用于安全验证的 UUID (v4)
func NewSecureRandomUUID() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

// NewUUID 生成一个UUID (v1)
func NewUUID() (string, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}
