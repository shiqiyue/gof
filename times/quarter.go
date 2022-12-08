package times

import (
	"fmt"
	"github.com/jinzhu/now"
	"time"
)

// 通过月份获取季度
func GetQuarterByMonth(month int) int {
	return (month-1)/3 + 1
}

// 获取上一个季度开始时间
func GetPreQuarterStart() time.Time {
	return GetPreQuarterStartWithTime(time.Now())
}

// 获取上一季度开始时间
func GetPreQuarterStartWithTime(n time.Time) time.Time {
	t := now.With(n).BeginningOfQuarter().Add(-time.Minute)
	return now.With(t).BeginningOfQuarter()
}

// 获取上一个季度结束时间
func GetPreQuarterEnd() time.Time {
	return GetPreQuarterEndWithTime(time.Now())
}

// 获取上一个季度结束时间
func GetPreQuarterEndWithTime(n time.Time) time.Time {
	t := now.With(n).BeginningOfQuarter().Add(-time.Minute)
	return now.With(t).EndOfQuarter()
}

// 获取当前季度开始时间
func GetCurrentQuarterStart() time.Time {
	n := time.Now()
	return now.With(n).BeginningOfQuarter()
}

// 获取当前季度结束时间
func GetCurrentQuarterEnd() time.Time {
	n := time.Now()
	return now.With(n).EndOfQuarter()
}

// 下一季度
func NextQuarter(year, quarter int) (int, int) {
	if quarter == 4 {
		return year + 1, 1
	}
	return year, quarter + 1
}

// 获取季度开始时间
func GetQuarterStart(year, quarter int) time.Time {
	month := (quarter-1)*3 + 1
	str := fmt.Sprintf("%d-%02d-01 00:00:00", year, month)
	t, _ := ParseNormalDateTime(str)
	return t
}
