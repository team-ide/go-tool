package util

import (
	"crypto/md5"
	"fmt"
	"go.uber.org/zap"
	"io"
)

// MD5 获取MD5字符串
// @param str string "需要MD5的字符串"
func MD5(str string) string {
	m := md5.New()
	_, err := io.WriteString(m, str)
	if err != nil {
		Logger.Error("MD5 error", zap.Error(err))
	}
	arr := m.Sum(nil)
	return fmt.Sprintf("%x", arr)
}
