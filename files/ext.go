package files

import "strings"

// 返回文件后缀,带.
// 例如1.jpg,则返回.jpg
func Ext(fileName string) string {
	extI := strings.LastIndex(fileName, ".")
	ext := fileName[extI:]
	return ext
}
