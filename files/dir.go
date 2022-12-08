package files

import (
	"os"
	"path/filepath"
)

// 写个函数，确保文件夹存在，省的重复写
func EnsureDirExist(path string) error {
	dir := filepath.Dir(path)
	exists := IsPathExists(dir)
	if !exists {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

// 判断路径是否存在
func IsPathExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}
	return true
}
