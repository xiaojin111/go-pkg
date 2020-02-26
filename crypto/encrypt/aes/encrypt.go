package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"

	"github.com/jinmukeji/go-pkg/v2/crypto/rand"
)

// AESGCMEncrypt 根据 key 将明文 plainText 进行 AES-GCM 加密，根据 additionalData 进行验证，返回密文
// The key argument should be the AES key,
// either 16, 24, or 32 bytes to select
// AES-128, AES-192, or AES-256.
func AESGCMEncrypt(key, plainText, additionalData []byte) ([]byte, error) {
	// generate a new aes cipher using our 32 byte long key
	// NewCipher creates and returns a new cipher.Block.
	// The key argument should be the AES key,
	// either 16, 24, or 32 bytes to select
	// AES-128, AES-192, or AES-256.
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// gcm or Galois/Counter Mode, is a mode of operation
	// for symmetric key cryptographic block ciphers
	// - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	aesgcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce, err := rand.RandomBytes(aesgcm.NonceSize())
	if err != nil {
		return nil, err
	}

	// out <= nonce + encrypted data
	out := aesgcm.Seal(nonce, nonce, plainText, additionalData)
	return out, nil
}

// AESGCMDecrypt 根据 key 将密文 cipherMessage 进行 AES-GCM 解密，根据 additionalData 进行验证，返回明文
// The key argument should be the AES key,
// either 16, 24, or 32 bytes to select
// AES-128, AES-192, or AES-256.
func AESGCMDecrypt(key, cipherMessage, additionalData []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(cipherMessage) < nonceSize {
		return nil, errors.New("incorrect cipher message length")
	}

	nonce, cipherText := cipherMessage[:nonceSize], cipherMessage[nonceSize:]
	return gcm.Open(nil, nonce, cipherText, additionalData)
}
