package util

import (
	"github.com/google/uuid"
	"strings"
)

// GetUUID 生成UUID
// GetUUID()
func GetUUID() (res string) {
	res = uuid.NewString()
	res = strings.ReplaceAll(res, "-", "")
	return
}
