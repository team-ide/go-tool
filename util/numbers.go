package util

import (
	"math/rand"
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
func RandomInt(min int, max int) (res int) {
	res = min + RandForRandomInt.Intn(max-min+1)
	return
}

// RandomInt64 获取随机数
// @param min int64 "最小值"
// @param max int64 "最大值"
// @return int64 "随机数"
func RandomInt64(min int64, max int64) (res int64) {
	res = min + RandForRandomInt.Int63n(max-min+1)
	return
}
