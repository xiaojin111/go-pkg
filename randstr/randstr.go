// Package randstr 提供生成随机字符串的工具方法
package randstr

import "crypto/rand"

// Mask 定义
const (
	// Symbols 符号
	Symbols = "!@#$%^&*()"

	// Digits 数字
	Digits = "0123456789"

	// Letters 字母
	Letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// LetterSymbols 字母与符号
	LetterSymbols = "!@#$%^&*()abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// LetterDigits 字母与数字
	LetterDigits = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// LetterSymbolDigits 字母、符号与数字
	LetterSymbolDigits = "!@#$%^&*()0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	Email = "^[A-Za-z0-9!#$%&'+/=?^_{|}~-]+(.[A-Za-z0-9!#$%&'+/=?^_{|}~-]+)*@([A-Za-z0-9]+(?:-[A-Za-z0-9]+)?.)+[A-Za-z0-9]+(-[A-Za-z0-9]+)?$"
)

// RandASCIIString 生成随机字符串
func RandASCIIString(mask string, n int) string {
	output := make([]byte, n)
	// We will take n bytes, one byte for each character of output.
	randomness := make([]byte, n)
	// read all random
	_, err := rand.Read(randomness)
	if err != nil {
		panic(err)
	}
	l := len(mask)
	// fill output
	for pos := range output {
		// get random item
		random := uint8(randomness[pos])
		// random % 64
		randomPos := random % uint8(l)
		// put into output
		output[pos] = mask[randomPos]
	}
	return string(output)
}
