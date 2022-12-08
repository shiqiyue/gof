package times

import (
	"fmt"
	"testing"
)

func TestGetPreQuarterEnd(t *testing.T) {
	e := GetPreQuarterEnd()
	fmt.Println(e)
}

func TestGetPreQuarterStart(t *testing.T) {
	e := GetPreQuarterStart()
	fmt.Println(e)
}

func TestGetQuarterByMonth(t *testing.T) {

}

func TestGetQuarterStart(t *testing.T) {
	tt := GetQuarterStart(2020, 4)
	fmt.Println(tt)
}
