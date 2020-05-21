package ip

import (
	"encoding/binary"
	"fmt"
	"net"
)

// IPV4ToInt IPv4 地址字符串到 uint32 转换
func IPV4ToInt(ipAddr string) (uint32, error) {
	ip := net.ParseIP(ipAddr)
	if ip == nil {
		return 0, fmt.Errorf("wrong ip format: %s", ip)
	}

	ip = ip.To4()
	return binary.BigEndian.Uint32(ip), nil
}

// IntToIPv4 uint32 转换为 IPv4 字符串
func IntToIPv4(i uint32) string {
	ipByte := make([]byte, 4)
	binary.BigEndian.PutUint32(ipByte, i)
	ip := net.IP(ipByte)
	return ip.String()
}
