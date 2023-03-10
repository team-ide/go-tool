package util

import (
	"fmt"
	"math"
	"testing"
	"time"
)

func TestGetStringValue(t *testing.T) {
	fmt.Println("GetStringValue 11111", GetStringValue(11111))
	fmt.Println("GetStringValue 111", GetStringValue(int8(111)))
	fmt.Println("GetStringValue 9999", GetStringValue(int16(9999)))
	fmt.Println("GetStringValue 2147483647", GetStringValue(int32(2147483647)))
	fmt.Println("GetStringValue 9999999999999999", GetStringValue(9999999999999999))
	fmt.Println("GetStringValue 2147482000", GetStringValue(float32(2147482000)))
	fmt.Println("GetStringValue 2147420.99", GetStringValue(float32(2147420.99)))
	fmt.Println("GetStringValue 99999999999999", GetStringValue(float64(99999999999999)))
	fmt.Println("GetStringValue 999999999999.9999", GetStringValue(999999999999.9999))
	fmt.Println("GetStringValue math.MaxInt", GetStringValue(math.MaxInt))
	fmt.Println("GetStringValue math.MaxInt8", GetStringValue(math.MaxInt8))
	fmt.Println("GetStringValue math.MaxInt16", GetStringValue(math.MaxInt16))
	fmt.Println("GetStringValue math.MaxInt32", GetStringValue(math.MaxInt32))
	fmt.Println("GetStringValue math.MaxInt64", GetStringValue(math.MaxInt64))
	fmt.Println("GetStringValue math.MaxFloat32", GetStringValue(math.MaxFloat32))
	fmt.Println("GetStringValue math.MaxFloat64", GetStringValue(math.MaxFloat64))

}
func TestBase(t *testing.T) {
	fmt.Println("IsEmpty", IsEmpty(""))
	fmt.Println("IsEmpty", IsEmpty("1"))
	fmt.Println("IsNotEmpty", IsNotEmpty(""))
	fmt.Println("IsNotEmpty", IsNotEmpty("1"))
	fmt.Println("IsTrue", IsTrue("0"))
	fmt.Println("IsTrue", IsTrue("1"))
	fmt.Println("IsFalse", IsFalse("0"))
	fmt.Println("IsFalse", IsFalse("1"))

	fmt.Println("ArrayIndexOf", ArrayIndexOf([]string{"2", "3", "1"}, "1"))
	fmt.Println("ArrayIndexOf", ArrayIndexOf([]string{"2", "3", "1"}, "22"))
	fmt.Println("ArrayIndexOf", ArrayIndexOf([]int64{2, 3, 1}, int64(1)))
	fmt.Println("ArrayIndexOf", ArrayIndexOf([]int64{2, 3, 1}, int64(0)))
	fmt.Println("IntIndexOf", IntIndexOf([]int{2, 3, 1}, 1))
	fmt.Println("IntIndexOf", IntIndexOf([]int{2, 3, 1}, 0))
	fmt.Println("Int64IndexOf", Int64IndexOf([]int64{2, 3, 1}, 1))
	fmt.Println("Int64IndexOf", Int64IndexOf([]int64{2, 3, 1}, 7))
	fmt.Println("StringIndexOf", StringIndexOf([]string{"2", "3", "1"}, "1"))
	fmt.Println("StringIndexOf", StringIndexOf([]string{"2", "3", "1"}, "12"))
}

func TestFile(t *testing.T) {
	fmt.Println("GetRootDir", GetRootDir())
	fmt.Println("FormatPath", FormatPath("/sss\\asd\\dads/dd"))
	fmt.Println("GetAbsolutePath", GetAbsolutePath("/sss\\asd\\dads/dd"))
	Exists, err := PathExists("/sss\\asd\\dads/dd")
	if err != nil {
		panic("PathExists error:" + err.Error())
	}
	fmt.Println("PathExists", Exists)
	fileMap, err := LoadDirFiles("/sss\\asd\\dads/dd")
	if err != nil {
		panic("LoadDirFiles error:" + err.Error())
	}
	fmt.Println("LoadDirFiles", len(fileMap))
	filenames, err := LoadDirFilenames("/xx\\ss")
	if err != nil {
		panic("LoadDirFilenames error:" + err.Error())
	}
	fmt.Println("LoadDirFilenames", len(filenames))
}

func TestAes(t *testing.T) {
	key := "xSsdAssAXssDDsDs"
	data := "xxx"
	res, err := AesEncryptCBCByKey(data, key)
	if err != nil {
		panic("AesEncryptCBCByKey error:" + err.Error())
	}
	fmt.Println("AesEncryptCBCByKey", res)
	res, err = AesDecryptCBCByKey(res, key)
	if err != nil {
		panic("AesDecryptCBCByKey error:" + err.Error())
	}
	fmt.Println("AesDecryptCBCByKey", res)

	res, err = AesEncryptECBByKey(data, key)
	if err != nil {
		panic("AesEncryptECBByKey error:" + err.Error())
	}
	fmt.Println("AesEncryptECBByKey", res)
	res, err = AesDecryptECBByKey(res, key)
	if err != nil {
		panic("AesDecryptECBByKey error:" + err.Error())
	}
	fmt.Println("AesDecryptECBByKey", res)

}

func TestLock(t *testing.T) {
	LockByKey("xxx")
	UnlockByKey("xxx")
}

func TestIp(t *testing.T) {
	fmt.Println("GetLocalIPList", GetLocalIPList())
}

func TestUtil(t *testing.T) {
	var err error

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

	fmt.Println("GetUUID", GetUUID())

	fmt.Println("GetMD5", GetMD5("xxx"))

}
