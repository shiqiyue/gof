package gorms

import (
	"bytes"
	"strconv"
	"strings"
)

func IntArrayStr(a ...int) string {
	if len(a) == 0 {
		return "{}"
	}
	strs := make([]string, 0)
	for _, i := range a {
		strs = append(strs, strconv.Itoa(i))
	}
	bs := bytes.Buffer{}
	bs.WriteString("{")
	bs.WriteString(strings.Join(strs, ","))
	bs.WriteString("}")
	return bs.String()
}

func IntArrayStrP(a ...*int) string {
	if len(a) == 0 {
		return "{}"
	}
	intArray := make([]int, 0)
	for _, i := range a {
		intArray = append(intArray, *i)
	}
	return IntArrayStr(intArray...)
}

func StrToIntArray(str string) []int {
	re := make([]int, 0)
	if str == "" {
		return re
	}
	newStr := strings.ReplaceAll(
		strings.ReplaceAll(str, "{", ""),
		"}", "")
	if newStr == "" {
		return re
	}
	strArray := strings.Split(newStr, ",")
	for _, s := range strArray {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		re = append(re, i)
	}
	return re
}
