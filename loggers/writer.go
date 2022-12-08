package loggers

import (
	"github.com/shiqiyue/gof/loggers/redis_zap"
	"io"
	"os"
)

type Writer []io.Writer

func (w *Writer) AddSysOut() {
	*w = append(*w, os.Stdout)
}

func (w *Writer) AddFileOut(filepath string) {
	file := NetFileWriter(filepath)
	*w = append(*w, file)
}

func (w *Writer) AddRedisOut(client redis.UniversalClient, key string) {
	*w = append(*w, redis_zap.NewRedisWriter(key, client))
}
