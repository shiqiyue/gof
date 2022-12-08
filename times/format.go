package times

import "time"

// 转换成yyyyMMdd
func FormatToyyyyMMdd(d time.Time) string {
	return d.Format(yyyyMMdd)
}

// 转换成一般的date 格式
func FormatToDate(d time.Time) string {
	return d.Format(normalDate)
}

// 转换成一般的time 格式
func FormatToTime(d time.Time) string {
	return d.Format(normalTime)
}

// 转换成一般的datetime 格式
func FormatToDateTime(d time.Time) string {
	return d.Format(normalDateTime)
}
