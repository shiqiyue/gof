package fileuploads

import (
	"context"
	"os"
)

// 文件上传结果
type FileUploadResult struct {
	// 内网地址
	//InUrl string
	// 外网地址
	//OutUrl string
	// 文件链接地址
	Url string
}

type FileUploadProvide func(ctx context.Context, f *os.File) (res *FileUploadResult, err error)
