package crypto

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
	assert.Equal(t, "+vd8pIrFyfTmpSfYdt9vRKHEgjjv4Ig==", encryptedText)

}

//TestDecrypt测试解密算法
func (suite *EncrypltTestTestSuite) TestDecrypt() {
	t := suite.T()
	h := NewPasswordCipherHelper()
	plainText := h.Decrypt("GQv9wl3NP5CPDr1qbiZ+/WwInNSVDRE7yMM=", "3LaO", "jinmuid")
	suite.T().Log(plainText)
	assert.Equal(t, "a1234567", plainText)
}

//TestDecrypt测试解密算法
func (suite *EncrypltTestTestSuite) TestDecryptT() {
	t := suite.T()
	h := NewPasswordCipherHelper()
	plainText := h.Decrypt("8jp4+6HR+ktnp0vV5/kcZ6C3XlZhBmjLoCA=", "Keg2", "aLGGjPzg7My5Kxv0s9hC")
	suite.T().Log(plainText)
	assert.Equal(t, "a1234567", plainText)
}
func TestEncrypltTestTestSuite(t *testing.T) {
	suite.Run(t, new(EncrypltTestTestSuite))
}
