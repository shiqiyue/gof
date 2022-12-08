package times

import (
	"fmt"
	"testing"
)

func TestGetPreMonthEnd(t *testing.T) {
	e := GetCurrentMonthEnd()
	fmt.Println(e)
}

func TestGetPreMonthStart(t *testing.T) {
	e := GetCurrentMonthStart()
	fmt.Println(e)
}
