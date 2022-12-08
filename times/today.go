package times

import "time"

func GetToday() time.Time {
	today, _ := time.ParseInLocation(normalDateTime, time.Now().Format(dateZeroTime), time.Local)
	return today
}
