package fileuploads

import (
	"context"
	"os"
)

type FileUploader interface {
	UploadFile(ctx context.Context, file *os.File) (res *FileUploadResult, err error)
}
