package gorms

import (
	"bytes"
	"strings"
)

func TextArrayStr(a ...string) string {
	if len(a) == 0 {
		return "{}"
	}

	bs := bytes.Buffer{}
	bs.WriteString("{")
	bs.WriteString(strings.Join(a, ","))
	bs.WriteString("}")
	return bs.String()
}

func TextArrayStrP(a ...*string) string {
	if len(a) == 0 {
		return "{}"
	}

	strs := make([]string, 0)
	for _, i := range a {
		strs = append(strs, *i)
	}
	return TextArrayStr(strs...)
}

func StrToTextArray(str string) []string {
	if str == "" || str == "{}" {
		return []string{}
	}
	newStr := strings.ReplaceAll(
		strings.ReplaceAll(str, "{", ""),
		"}", "")
	if newStr == "" {
		return []string{}
	}
	strArray := strings.Split(newStr, ",")

	return strArray
}
