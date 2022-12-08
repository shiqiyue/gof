package times

import "time"

// 判断时间是否相交
func IsTimeIntersect(startTime1 time.Time, endTime1 time.Time, startTime2 time.Time, endTime2 time.Time) bool {
	return IsTimeBetween(startTime1, startTime2, endTime2) || IsTimeBetween(startTime2, startTime1, endTime1)

}

// 判断是否处于时间段
func IsTimeBetween(t time.Time, startTime time.Time, endTime time.Time) bool {
	if (t.Equal(startTime) || t.After(startTime)) && (t.Equal(endTime) || t.Before(endTime)) {
		return true
	}
	return false
}
