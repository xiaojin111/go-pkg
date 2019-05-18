package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

const keyLength = 32

type DefaultPasswordCipherHelper struct {
}

type PasswordCipherHelper interface {
	Encrypt(str, seed, key string) string
	Decrypt(encryptedStr, seed, key string) string
}

func NewPasswordCipherHelper() PasswordCipherHelper {
	return DefaultPasswordCipherHelper{}
}

// Encrypt 加密算法
func (client DefaultPasswordCipherHelper) Encrypt(str, seed, key string) string {
	return base64.StdEncoding.EncodeToString(AESEncrypt([]byte(str), []byte(key+seed)))
}

// Decrypt 解密算法
func (client DefaultPasswordCipherHelper) Decrypt(encrypted, seed, key string) string {
	encryptedPassword, _ := base64.StdEncoding.DecodeString(encrypted)
	return string(AESDecrypt([]byte(encryptedPassword), []byte(key+seed)))
}

// AESEncrypt aes中的CFB加密
func AESEncrypt(src []byte, key []byte) (ciphertext []byte) {
	block, err := aes.NewCipher(generateKey(key))
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext = make([]byte, aes.BlockSize+len(src))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], src)
	return ciphertext
}

// AESDecrypt aes中的CFB解密
func AESDecrypt(ciphertext []byte, key []byte) (plaintext []byte) {
	block, err := aes.NewCipher(generateKey(key))
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)
	return ciphertext
}

// generateKey key扩充到32位
func generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, keyLength)
	copy(genKey, key)
	for i := keyLength; i < len(key); {
		for j := 0; j < keyLength && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}
