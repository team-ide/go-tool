package util

import (
	"math/rand"
	"strconv"
	"time"
)

var (
	// RandForRandomInt 设置随机数种子
	RandForRandomInt = rand.New(rand.NewSource(time.Now().UnixNano()))
)

// RandomInt 获取随机数
// @param min int "最小值"
// @param max int "最大值"
// @return int "随机数"
// RandomInt(1, 10)
func RandomInt(min int, max int) (res int) {
	res = min + RandForRandomInt.Intn(max-min+1)
	return
}

// RandomInt64 获取随机数
// @param min int64 "最小值"
// @param max int64 "最大值"
// @return int64 "随机数"
// RandomInt64(1, 10)
func RandomInt64(min int64, max int64) (res int64) {
	res = min + RandForRandomInt.Int63n(max-min+1)
	return
}

// StringToInt 字符串转 int
// StringToInt("11")
func StringToInt(str string) int {
	if str == "" {
		return 0
	}
	res, _ := strconv.Atoi(str)
	return res
}

// StringToInt64 字符串转 int64
// StringToInt64("11")
func StringToInt64(str string) int64 {
	if str == "" {
		return 0
	}
	res, _ := strconv.ParseInt(str, 10, 64)
	return res
}

// StringToFloat64 字符串转 float64
// StringToFloat64("11.2")
func StringToFloat64(str string) float64 {
	if str == "" {
		return 0
	}
	res, _ := strconv.ParseFloat(str, 64)
	return res
}
