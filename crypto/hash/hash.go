package hash

import (
	"crypto/md5"    // nolint: gas
	"crypto/sha1"   // nolint: gas
	"crypto/sha256" // nolint: gas
	"fmt"
	"hash"
)

func hashString(h hash.Hash, in string) string {
	return hashBytes(h, []byte(in))
}

func hashBytes(h hash.Hash, in []byte) string {
	h.Write(in) // nolint: errcheck, gas
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

// SHA256 给定字符串，计返回SHA256十六进制字符串Hash值，其中字母部分为小写字符
func SHA256(in string) string {
	h := sha256.New()
	return hashString(h, in)
}

// SHA256Bytes 给定 []byte，计返回SHA256十六进制字符串Hash值，其中字母部分为小写字符
func SHA256Bytes(in []byte) string {
	h := sha256.New()
	return hashBytes(h, in)
}

// SHA1 给定字符串，计返回SHA1十六进制字符串Hash值，其中字母部分为小写字符
func SHA1(in string) string {
	h := sha1.New()
	return hashString(h, in)
}

// SHA1Bytes 给定 []byte，计返回SHA1十六进制字符串Hash值，其中字母部分为小写字符
func SHA1Bytes(in []byte) string {
	h := sha1.New()
	return hashBytes(h, in)
}

// MD5 给定字符串，计返回MD5十六进制字符串Hash值，其中字母部分为小写字符
func MD5(in string) string {
	h := md5.New() // nolint: gas
	return hashString(h, in)
}

// MD5Bytes 给定 []byte，计返回MD5十六进制字符串Hash值，其中字母部分为小写字符
func MD5Bytes(in []byte) string {
	h := md5.New() // nolint: gas
	return hashBytes(h, in)
}
