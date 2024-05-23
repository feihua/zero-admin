package common

import "time"

// TimeToString 时间指针转字符串
func TimeToString(t *time.Time) string {
	if t != nil {
		return t.Format("2006-01-02 15:04:05")
	}
	return ""
}
