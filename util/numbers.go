package util

import (
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type RandForRandom struct {
	rand *rand.Rand
	lock sync.Locker
}

func NewRandForRandom() *RandForRandom {

	return &RandForRandom{
		//  设置随机数种子
		rand: rand.New(rand.NewSource(time.Now().UnixNano())),
		lock: &sync.Mutex{},
	}
}

func (this_ *RandForRandom) RandomInt(min int, max int) (res int) {
	this_.lock.Lock()
	defer this_.lock.Unlock()

	res = min + this_.rand.Intn(max-min+1)
	return
}

func (this_ *RandForRandom) RandomInt64(min int64, max int64) (res int64) {
	this_.lock.Lock()
	defer this_.lock.Unlock()

	res = min + this_.rand.Int63n(max-min+1)
	return
}

var (
	randForRandomIntList []*RandForRandom
	randForRandomIntSize int
	randForRandomIntC    int
)

func init() {
	InitRandForRandomIntList(20)
}
func InitRandForRandomIntList(size int) {
	if size < 1 {
		return
	}
	var list []*RandForRandom
	for i := 0; i < size; i++ {
		list = append(list, NewRandForRandom())
	}
	randForRandomIntList = list
	randForRandomIntSize = size
}

// RandomInt 获取随机数
// @param min int "最小值"
// @param max int "最大值"
// @return int "随机数"
// RandomInt(1, 10)
func RandomInt(min int, max int) (res int) {
	randForRandomIntC++
	if randForRandomIntC > 10000 {
		randForRandomIntC = 0
	}
	randForRandomInt := randForRandomIntList[randForRandomIntC%randForRandomIntSize]
	return randForRandomInt.RandomInt(min, max)
}

// RandomInt64 获取随机数
// @param min int64 "最小值"
// @param max int64 "最大值"
// @return int64 "随机数"
// RandomInt64(1, 10)
func RandomInt64(min int64, max int64) (res int64) {
	randForRandomIntC++
	if randForRandomIntC > 10000 {
		randForRandomIntC = 0
	}
	randForRandomInt := randForRandomIntList[randForRandomIntC%randForRandomIntSize]
	return randForRandomInt.RandomInt64(min, max)
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

// StringToUint64 字符串转 uint64
// StringToUint64("11")
func StringToUint64(str string) uint64 {
	if str == "" {
		return 0
	}
	res, _ := strconv.ParseUint(str, 10, 64)
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

// SumToString 将任意数相加，为防止精度丢失，可以传入数字字符串
// SumToString("4611686027042965191", 11)
func SumToString(nums ...interface{}) string {
	var res int64
	for _, num := range nums {
		if num == nil || num == "" || num == 0 {
			continue
		}
		res += StringToInt64(GetStringValue(num))
	}
	return strconv.FormatInt(res, 10)
}

// ValueToInt64 值 转 int64
// ValueToInt64("11")
func ValueToInt64(value any) (res int64, err error) {
	if value == nil {
		return
	}
	switch tV := value.(type) {
	case int:
		res = int64(tV)
		return
	case uint:
		res = int64(tV)
		return
	case int8:
		res = int64(tV)
		return
	case uint8:
		res = int64(tV)
		return
	case int16:
		res = int64(tV)
		return
	case uint16:
		res = int64(tV)
		return
	case int32:
		res = int64(tV)
		return
	case uint32:
		res = int64(tV)
		return
	case int64:
		res = tV
		return
	case uint64:
		res = int64(tV)
		return
	case float32:
		res = int64(tV)
		return
	case float64:
		res = int64(tV)
		return
	case bool:
		if tV {
			res = 1
		}
		return
	case time.Time:
		res = tV.UnixMilli()
		return
	default:
		str := GetStringValue(value)
		if str != "" {
			res, err = strconv.ParseInt(str, 10, 64)
		}
	}
	return
}

// ValueToUint64 值 转 uint64
// ValueToUint64("11")
func ValueToUint64(value any) (res uint64, err error) {
	if value == nil {
		return
	}
	switch tV := value.(type) {
	case int:
		res = uint64(tV)
		return
	case uint:
		res = uint64(tV)
		return
	case int8:
		res = uint64(tV)
		return
	case uint8:
		res = uint64(tV)
		return
	case int16:
		res = uint64(tV)
		return
	case uint16:
		res = uint64(tV)
		return
	case int32:
		res = uint64(tV)
		return
	case uint32:
		res = uint64(tV)
		return
	case int64:
		res = uint64(tV)
		return
	case uint64:
		res = tV
		return
	case float32:
		res = uint64(tV)
		return
	case float64:
		res = uint64(tV)
		return
	case bool:
		if tV {
			res = 1
		}
		return
	case time.Time:
		res = uint64(tV.UnixMilli())
		return
	default:
		str := GetStringValue(value)
		if str != "" {
			res, err = strconv.ParseUint(str, 10, 64)
		}
	}
	return
}

// ValueToFloat64 值 转 float64
// ValueToFloat64("11")
func ValueToFloat64(value any) (res float64, err error) {
	if value == nil {
		return
	}
	switch tV := value.(type) {
	case int:
		res = float64(tV)
		return
	case uint:
		res = float64(tV)
		return
	case int8:
		res = float64(tV)
		return
	case uint8:
		res = float64(tV)
		return
	case int16:
		res = float64(tV)
		return
	case uint16:
		res = float64(tV)
		return
	case int32:
		res = float64(tV)
		return
	case uint32:
		res = float64(tV)
		return
	case int64:
		res = float64(tV)
		return
	case uint64:
		res = float64(tV)
		return
	case float32:
		res = float64(tV)
		return
	case float64:
		res = tV
		return
	case bool:
		if tV {
			res = 1
		}
		return
	case time.Time:
		res = float64(tV.UnixMilli())
		return
	default:
		str := GetStringValue(value)
		if str != "" {
			res, err = strconv.ParseFloat(str, 64)
		}
	}
	return
}
