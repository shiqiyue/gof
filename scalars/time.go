package scalars

import "time"

func TimeIfNilReturnInt64(t *time.Time) int64 {
	if t == nil {
		return 0
	}
	return (*t).Unix()
}
