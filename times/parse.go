package times

import "time"

// 解析一般的date 格式
func ParseNormalDateTime(str string) (time.Time, error) {
	r, err := time.Parse(normalDateTime, str)
	return r, err
}
