package times

import (
	"fmt"
	"github.com/jinzhu/now"
	"time"
)

// 获取上一月度开始时间
func GetPreMonthStartWithTime(n time.Time) time.Time {
	t := now.With(n).BeginningOfMonth().Add(-time.Minute)
	return now.With(t).BeginningOfQuarter()
}

// 获取上一个月份开始时间
func GetPreMonthStart() time.Time {
	n := time.Now()
	return GetPreMonthStartWithTime(n)
}

// 获取上一月度结束时间
func GetPreMonthEndWithTime(n time.Time) time.Time {
	t := now.With(n).BeginningOfMonth().Add(-time.Minute)
	return now.With(t).EndOfMonth()
}

// 获取上一个月份结束时间
func GetPreMonthEnd() time.Time {
	n := time.Now()
	return GetPreMonthEndWithTime(n)
}

// 获取当前月份的开始时间
func GetCurrentMonthStart() time.Time {
	n := time.Now()
	return now.With(n).BeginningOfMonth()
}

// 获取当前月份的结束时间
func GetCurrentMonthEnd() time.Time {
	n := time.Now()
	return now.With(n).EndOfMonth()
}

// 下一月份
func NextMonth(year, month int) (int, int) {
	if month == 12 {
		return year + 1, 1
	}
	return year, month + 1
}

// 获取月度开始时间
func GetMonthStart(year, month int) time.Time {
	str := fmt.Sprintf("%d-%02d-01 00:00:00", year, month)
	t, _ := ParseNormalDateTime(str)
	return t
}
