package times

import (
	"time"
)

// 设置date(年月日)
func SetDate(t time.Time, year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
}
