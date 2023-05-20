package util

import "time"

// GetNow 获取当前时间
// GetNow()
func GetNow() time.Time {
	return time.Now()
}

// GetNowNano 获取当前时间戳  到纳秒
// GetNowNano()
func GetNowNano() int64 {
	return GetNanoByTime(time.Now())
}

// GetNowMilli 获取当前时间戳  到毫秒
// GetNowMilli()
func GetNowMilli() int64 {
	return GetMilliByTime(time.Now())
}

// GetNowSecond 获取当前时间戳 到秒
// GetNowSecond()
func GetNowSecond() int {
	return GetSecondByTime(time.Now())
}

// GetNanoByTime 获取时间戳  到纳秒
// @param v time.Time "时间"
// GetNanoByTime(time)
func GetNanoByTime(v time.Time) int64 {
	return v.UnixNano()
}

// GetMilliByTime 获取时间戳  到毫秒
// @param v time.Time "时间"
// GetMilliByTime(time)
func GetMilliByTime(v time.Time) int64 {
	return v.UnixMilli()
}

// GetSecondByTime 获取时间戳 到秒
// @param v time.Time "时间"
// GetSecondByTime(time)
func GetSecondByTime(v time.Time) int {
	return v.Second()
}

var (
	// DefaultTimeFormatLayout 默认时间格式化
	DefaultTimeFormatLayout = "2006-01-02 15:04:05"
)

// GetNowFormat 获取当前格式化时间 "2006-01-02 15:04:05"
// GetNowFormat()
func GetNowFormat() string {
	now := time.Now()
	return TimeFormat(now, DefaultTimeFormatLayout)
}

// GetFormatByTime 获取格式化时间 "2006-01-02 15:04:05"
// @param v time.Time "时间"
// GetFormatByTime(time)
func GetFormatByTime(v time.Time) string {
	return TimeFormat(v, DefaultTimeFormatLayout)
}

// TimeFormat 时间格式化 默认 "2006-01-02 15:04:05"
// @param v time.Time "时间"
// @param layout string "格式化字符串，默认使用"2006-01-02 15:04:05""
// TimeFormat(time, "2006-01-02 15:04:05")
func TimeFormat(v time.Time, layout string) string {
	if layout == "" {
		layout = DefaultTimeFormatLayout
	}
	return v.Format(layout)
}
