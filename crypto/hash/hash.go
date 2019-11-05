package hash

import (
	"crypto/md5"    // nolint: gas
	"crypto/sha1"   // nolint: gas
	"crypto/sha256" // nolint: gas
	"fmt"
	"hash"
)

func hashString(h hash.Hash, in string) []byte {
	return hashBytes(h, []byte(in))
}

func hashBytes(h hash.Hash, in []byte) []byte {
	h.Write(in) // nolint: errcheck, gas
	bs := h.Sum(nil)
	return bs
}

// HexString 返回字节数组的十六进制字符串，其中字母部分为小写字符
func HexString(in []byte) string {
	return fmt.Sprintf("%x", in)
}

// SHA256String 给定字符串，计返回SHA256 Hash值
func SHA256String(in string) []byte {
	h := sha256.New()
	return hashString(h, in)
}

// SHA256 给定 []byte，计返回SHA256 Hash值
func SHA256(in []byte) []byte {
	h := sha256.New()
	return hashBytes(h, in)
}

// SHA1String 给定字符串，计返回SHA1 Hash值
func SHA1String(in string) []byte {
	h := sha1.New()
	return hashString(h, in)
}

// SHA1 给定 []byte，计返回SHA1 Hash值
func SHA1(in []byte) []byte {
	h := sha1.New()
	return hashBytes(h, in)
}

// MD5 给定字符串，计返回MD5 Hash值
func MD5String(in string) []byte {
	h := md5.New() // nolint: gas
	return hashString(h, in)
}

// MD5Bytes 给定 []byte，计返回MD5 Hash值
func MD5(in []byte) []byte {
	h := md5.New() // nolint: gas
	return hashBytes(h, in)
}
