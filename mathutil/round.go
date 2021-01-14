package mathutil

import (
	"math"
)

func RoundToInt(f float64) int {
	r := math.Round(f)
	return int(r)
}

func RoundToInt32(f float64) int32 {
	r := math.Round(f)
	return int32(r)
}

func RoundToInt64(f float64) int64 {
	r := math.Round(f)
	return int64(r)
}
