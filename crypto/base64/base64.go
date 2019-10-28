package base64

import b64 "encoding/base64"

// Base64EncodeRawUrl Base64 编码
func Base64EncodeRawURL(in []byte) string {
	return b64.RawURLEncoding.EncodeToString(in)
}

// Base64DecodeRawUrl Base64 解码
func Base64DecodeRawURL(in string) ([]byte, error) {
	return b64.RawURLEncoding.DecodeString(in)
}
