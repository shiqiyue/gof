package loggers

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"io"
	"log"
	"time"
)

// 获取文件writer
func NetFileWriter(filename string) io.Writer {
	// 生成rotatelogs的Logger 实际生成的文件名 demo.log.YYmmddHH
	// demo.log是指向最新日志的链接
	// 保存7天内的日志，每1小时(整点)分割一次日志
	hook, err := rotatelogs.New(
		filename+".%Y%m%d", // 没有使用go风格反人类的format格式
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*24*7),
		rotatelogs.WithRotationTime(time.Hour*24),
	)

	if err != nil {
		log.Fatal(err)
	}
	return hook
}
