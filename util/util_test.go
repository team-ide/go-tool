package util

import (
	"fmt"
	"testing"
	"time"
)

func TestUtil(t *testing.T) {
	var err error
	fmt.Println("IsEmpty", IsEmpty(""))
	fmt.Println("IsNotEmpty", IsNotEmpty(""))
	fmt.Println("IsTrue", IsTrue("1"))
	fmt.Println("IsFalse", IsFalse("1"))

	fmt.Println("RandomInt", RandomInt(1, 100))
	fmt.Println("RandomInt64", RandomInt64(1, 100))

	fmt.Println("FirstToUpper", FirstToUpper("aaa"))
	fmt.Println("FirstToUpper", FirstToUpper("a张三"))
	fmt.Println("FirstToLower", FirstToLower("AAA"))
	fmt.Println("FirstToLower", FirstToLower("AAA张三"))
	fmt.Println("RandomString", RandomString(1, 100))
	fmt.Println("RandomUserName", RandomUserName(2))
	fmt.Println("GetStringValue", GetStringValue(1))
	py, err := ToPinYin("惠波琬")
	if err != nil {
		panic("ToPinYin error:" + err.Error())
	}
	fmt.Println("ToPinYin", py)

	fmt.Println("GetNow", GetNow())
	fmt.Println("GetNowTime", GetNowTime())
	fmt.Println("GetNowSecond", GetNowSecond())
	fmt.Println("GetTimeByTime", GetTimeByTime(GetNow()))
	fmt.Println("GetSecondByTime", GetSecondByTime(GetNow()))
	fmt.Println("GetNowFormat", GetNowFormat())
	fmt.Println("GetFormatByTime", GetFormatByTime(GetNow()))
	fmt.Println("GetFormatByTime", GetFormatByTime(time.Time{}))

	fmt.Println("UUID", UUID())

	fmt.Println("MD5", MD5("xxx"))

	fmt.Println("ArrayIndexOf", ArrayIndexOf([]string{"2", "3", "1"}, "1"))
	fmt.Println("ArrayIndexOf", ArrayIndexOf([]int64{2, 3, 1}, int64(1)))
	fmt.Println("ArrayIndexOf", Int64IndexOf([]int64{2, 3, 1}, 1))

}
