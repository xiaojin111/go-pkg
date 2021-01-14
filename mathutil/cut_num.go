package mathutil

import (
	"errors"
)

// LastNumRange 对 num 进行截断，提取自 last 位置开始的 len 长度个数字。错误时返回 -1。
func LastNumRange(num, last, len int64) (int64, error) {
	if len <= 0 || len > last {
		return -1, errors.New("invalid range")
	}

	lastBegin := last
	lastEnd := last - len

	n, err := LastNum(num, lastBegin)
	if err != nil {
		return -1, err
	}

	b := pow10(lastEnd)

	return n / b, nil
}

// LastNum 对 num 进行截断，提取自 last 位置开始的所有数字。错误时返回 -1。
func LastNum(num, last int64) (int64, error) {
	if last <= 0 {
		return -1, errors.New("invalid count")
	}
	b := pow10(last)
	return num % b, nil
}

func pow10(p int64) int64 {
	if p < 0 {
		return -1
	}

	var b int64 = 1
	var i int64
	for i = 0; i < p; i++ {
		b = b * 10
	}
	return b
}
