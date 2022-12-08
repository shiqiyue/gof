package fileuploads

import (
	"context"
	"os"
)

type FileUploadHook interface {
	// 上传前
	UploadBefore(ctx context.Context, file *os.File) (newCtx context.Context, err error)

	// 上传后
	UploadAfter(ctx context.Context, file *os.File, res *FileUploadResult, uploadErr error) (newCtx context.Context, err error)
}
