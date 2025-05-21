package time_util

import "time"

// TimeToString 时间指针转字符串
func TimeToString(t *time.Time) string {
	if t != nil {
		return t.Format("2006-01-02 15:04:05")
	}
	return ""
}

func TimeToString1(t *time.Time) string {
	if t != nil {
		return t.Format("2006-01-02")
	}
	return ""
}

// TimeToStr 时间转字符串
func TimeToStr(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
