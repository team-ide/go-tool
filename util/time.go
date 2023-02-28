package util

import "time"

// GetNowTime 获取当前时间戳
func GetNowTime() int64 {
	return GetTimeTime(time.Now())
}

// GetNowTimeSecond 获取当前时间戳 秒
func GetNowTimeSecond() int64 {
	return GetTimeSecond(time.Now())
}

// GetTimeTime 获取当前时间戳
func GetTimeTime(time time.Time) int64 {
	return time.UnixNano() / 1e6
}

// GetTimeSecond 获取当前时间秒
func GetTimeSecond(time time.Time) int64 {
	return time.Unix()
}

// Now 获取当前时间
func Now() time.Time {
	return time.Now()
}

func NowStr() string {
	now := time.Now()
	return Format(now)
}

func Format(date time.Time) string {
	return date.Format("2006-01-02 15:04:05.000")
}
