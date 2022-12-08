package fileuploads

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"os"
)

type OpenTracingHook struct {
}

func (o *OpenTracingHook) UploadBefore(ctx context.Context, file *os.File) (newCtx context.Context, err error) {
	_, newCtx = opentracing.StartSpanFromContext(ctx, "upload file")
	return newCtx, nil
}

func (o *OpenTracingHook) UploadAfter(ctx context.Context, file *os.File, res *FileUploadResult, uploadErr error) (newCtx context.Context, err error) {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		span.Finish()
	}
	return ctx, nil
}
