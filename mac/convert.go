package mac

import "strings"

// NormalizeMac 返回规范化的 MAC 地址。所有字母大写，去除冒号分隔符与 0x 前缀。
func NormalizeMac(mac string) string {
	ret := strings.ToUpper(mac)

	// 去除 0x 开头标记
	ret = strings.TrimPrefix(ret, "0X")

	// 去掉前后空格
	ret = strings.TrimSpace(ret)

	// 去除分号
	ret = strings.Replace(ret, ":", "", -1)

	// mac 地址长度为奇数时，前面补一个 0
	if len(ret)%2 == 1 {
		ret = "0" + ret
	}

	return ret
}
