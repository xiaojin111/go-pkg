package rand

import (
	"crypto/rand"
	"io"

	"github.com/jinmukeji/go-pkg/v2/crypto/base64"
)

// RandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func RandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := io.ReadFull(rand.Reader, b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// Mask 定义
const (
	MaskSymbols            = "!@#$%^&*()"
	MaskDigits             = "0123456789"
	MaskLetters            = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	MaskLetterSymbols      = "!@#$%^&*()abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	MaskLetterDigits       = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	MaskLetterSymbolDigits = "!@#$%^&*()0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// RandomString returns a securely generated random digits string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func RandomString(n int) (string, error) {
	return RandomStringWithMask(MaskLetterSymbolDigits, n)
}

// RandomString returns a securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func RandomDigits(n int) (string, error) {
	return RandomStringWithMask(MaskDigits, n)
}

// RandomStringWithMask returns a securely generated random string with a mask string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func RandomStringWithMask(mask string, n int) (string, error) {
	bytes, err := RandomBytes(n)
	if err != nil {
		return "", err

	}
	for i, b := range bytes {
		bytes[i] = mask[b%byte(len(mask))]
	}
	return string(bytes), nil
}

// RandomStringURLSafe returns a URL-safe, base64 encoded
// securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func RandomStringURLSafe(n int) (string, error) {
	b, err := RandomBytes(n)
	if err != nil {
		return "", err
	}
	return base64.Base64EncodeRawURL(b), nil
}
