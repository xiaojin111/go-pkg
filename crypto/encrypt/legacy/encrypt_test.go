package legacy

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type EncrypltTestTestSuite struct {
	suite.Suite
}

// TestEncrypt测试加密算法
func (suite *EncrypltTestTestSuite) TestEncrypt() {
	t := suite.T()
	h := NewPasswordCipherHelper()

	encryptedText := h.Encrypt("a1234567", "a123", "jinmuid")
	suite.T().Log(encryptedText)

	decryptedTest := h.Decrypt(encryptedText, "a123", "jinmuid")
	suite.T().Log(decryptedTest)
	assert.Equal(t, "a1234567", decryptedTest)
}

//TestDecrypt测试解密算法
func (suite *EncrypltTestTestSuite) TestDecrypt() {
	t := suite.T()
	h := NewPasswordCipherHelper()
	plainText := h.Decrypt("GQv9wl3NP5CPDr1qbiZ+/WwInNSVDRE7yMM=", "3LaO", "jinmuid")
	suite.T().Log(plainText)
	assert.Equal(t, "123456789a", plainText)
}

//TestDecrypt测试解密算法
func (suite *EncrypltTestTestSuite) TestDecrypt2() {
	t := suite.T()
	h := NewPasswordCipherHelper()
	plainText := h.Decrypt("8jp4+6HR+ktnp0vV5/kcZ6C3XlZhBmjLoCA=", "Keg2", "aLGGjPzg7My5Kxv0s9hC")
	suite.T().Log(plainText)
	assert.Equal(t, "hengyang73", plainText)
}
func TestEncrypltTestTestSuite(t *testing.T) {
	suite.Run(t, new(EncrypltTestTestSuite))
}
