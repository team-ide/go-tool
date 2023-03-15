package util

import (
	"errors"
	"os"
	"reflect"
)

// IsEmpty 是否为nil或空字符串
// @param v interface{} "传入任意值"
// @return bool
func IsEmpty(v interface{}) bool {
	if v == nil || v == "" {
		return true
	}
	return false
}

// IsNotEmpty 是否不为nil或空字符串
// @param v interface{} "传入任意值"
// @return bool
func IsNotEmpty(v interface{}) bool {
	return !IsEmpty(v)
}

// IsTrue 是否为真 判断是true、"true"、1、"1"
// @param v interface{} "传入任意值"
// @return bool
func IsTrue(v interface{}) (res bool) {
	if v == true || v == "true" || v == 1 || v == "1" {
		res = true
	}
	return
}

// IsFalse 是否为否 判断不是true、"true"、1、"1"
// @param v interface{} "传入任意值"
// @return bool
func IsFalse(v interface{}) (res bool) {
	return !IsTrue(v)
}

// IntIndexOf 返回 某个值 在数组中的索引位置，未找到返回 -1
func IntIndexOf(array []int, v int) (index int) {
	index = -1
	size := len(array)
	for i := 0; i < size; i++ {
		if array[i] == v {
			index = i
			return
		}
	}
	return
}

// Int64IndexOf 返回 某个值 在数组中的索引位置，未找到返回 -1
func Int64IndexOf(array []int64, v int64) (index int) {
	index = -1
	size := len(array)
	for i := 0; i < size; i++ {
		if array[i] == v {
			index = i
			return
		}
	}
	return
}

// StringIndexOf 返回 某个值 在数组中的索引位置，未找到返回 -1
func StringIndexOf(array []string, v string) (index int) {
	index = -1
	size := len(array)
	for i := 0; i < size; i++ {
		if array[i] == v {
			index = i
			return
		}
	}
	return
}

// ArrayIndexOf 返回 某个值 在数组中的索引位置，未找到返回 -1
func ArrayIndexOf(array interface{}, v interface{}) (index int) {
	index = -1
	vOf := reflect.ValueOf(array)
	if vOf.IsNil() {
		return
	}
	kind := vOf.Kind()
	switch kind {
	case reflect.Array, reflect.Slice:
		size := vOf.Len()
		for i := 0; i < size; i++ {
			iV := vOf.Index(i)
			if iV.Interface() == v {
				index = i
				return
			}
		}
	}
	return
}

var (
	tempDir = ""
)

// SetTempDir 设置临时目录
func SetTempDir(dir string) {
	tempDir = dir
}

// GetTempDir 获取临时目录
func GetTempDir() (dir string, err error) {
	if tempDir != "" {
		dir = tempDir
		return
	}
	tempDir = os.TempDir()
	tempDir = tempDir + "/go-tool-temp"
	tempDir = FormatPath(tempDir)

	exists, err := PathExists(tempDir)
	if err != nil {
		err = errors.New("PathExists error,path:" + tempDir + ",error:" + err.Error())
		return
	}
	if !exists {
		err = os.MkdirAll(tempDir, os.ModePerm)
		if err != nil {
			err = errors.New("MkdirAll error,path:" + tempDir + ",error:" + err.Error())
			return
		}
	}
	dir = tempDir
	return
}
