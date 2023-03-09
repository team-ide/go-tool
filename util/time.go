package util

import "time"

// GetNow 获取当前时间
func GetNow() time.Time {
	return time.Now()
}

// GetNowTime 获取当前时间戳  到毫秒
func GetNowTime() int64 {
	return GetTimeByTime(time.Now())
}

// GetNowSecond 获取当前时间戳 到秒
func GetNowSecond() int64 {
	return GetSecondByTime(time.Now())
}

// GetTimeByTime 获取时间戳  到毫秒
// @param v time.Time "时间"
func GetTimeByTime(v time.Time) int64 {
	return v.UnixNano() / 1e6
}

// GetSecondByTime 获取时间戳 到秒
// @param v time.Time "时间"
func GetSecondByTime(v time.Time) int64 {
	return v.Unix()
}

var (
	// DefaultTimeFormatLayout 默认时间格式化
	DefaultTimeFormatLayout = `2006-01-02 15:04:05`
)

// GetNowFormat 获取当前格式化时间 `2006-01-02 15:04:05`
func GetNowFormat() string {
	now := time.Now()
	return TimeFormat(now, DefaultTimeFormatLayout)
}

// GetFormatByTime 获取格式化时间 `2006-01-02 15:04:05`
// @param v time.Time "时间"
func GetFormatByTime(v time.Time) string {
	return TimeFormat(v, DefaultTimeFormatLayout)
}

// TimeFormat 时间格式化 默认 `2006-01-02 15:04:05`
// @param v time.Time "时间"
// @param layout string "格式化字符串，默认使用`2006-01-02 15:04:05`"
func TimeFormat(v time.Time, layout string) string {
	if layout == "" {
		layout = DefaultTimeFormatLayout
	}
	return v.Format(layout)
}
