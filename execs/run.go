package execs

import (
	"fmt"
	"github.com/gookit/goutil/sysutil"
)

// RunInDir ...
func RunInDir(cmd, dir string) ([]byte, error) {
	fmt.Println(cmd, dir)
	s, err := sysutil.QuickExec(cmd, dir)
	if err != nil {
		fmt.Println("执行命令出错", cmd, dir, err)
	}
	return []byte(s), err
}

// RunInteractiveInDir ...
func RunInteractiveInDir(cmd, dir string) error {
	_, err := RunInDir(cmd, dir)
	if err != nil {
		fmt.Println("执行命令出错", cmd, dir, err)
	}
	return err
}

// RunInteractive ...
func RunInteractive(cmd string) error {
	return RunInteractiveInDir(cmd, "")
}
