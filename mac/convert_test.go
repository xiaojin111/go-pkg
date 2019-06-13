package mac

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/suite"
)

type NormalizeMacTestTestSuite struct {
	suite.Suite
}

// TestNormalizeMac 返回规范化的 MAC 地址
func (suite *NormalizeMacTestTestSuite) TestNormalizeMac() {
	t := suite.T()
	mac := "304511440CCF"
	ret := NormalizeMac(mac)
	hexMac, _ := strconv.ParseUint(ret, 16, 64)
	fmt.Println(hexMac)
	assert.Equal(t, "0CB2B717B18D", ret)

}

func TestNormalizeMacTestTestSuite(t *testing.T) {
	suite.Run(t, new(NormalizeMacTestTestSuite))
}
