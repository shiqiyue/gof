package times

import (
	"fmt"
	"testing"
)

func TestIsTimeIntersect(t *testing.T) {
	t1, _ := ParseNormalDateTime("2011-01-01 20:07:02")
	t2, _ := ParseNormalDateTime("2011-01-01 21:07:02")
	t3, _ := ParseNormalDateTime("2011-01-01 11:19:55")
	t4, _ := ParseNormalDateTime("2011-01-01 12:19:55")

	r := IsTimeIntersect(t1, t2, t3, t4)
	fmt.Println(r)
}
