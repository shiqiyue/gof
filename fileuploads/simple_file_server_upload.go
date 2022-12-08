package fileuploads

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type SimpleFileServerUploader struct {
	// 客户端
	Client *http.Client
	// 链接
	Url string
	// hook
	Hooks []FileUploadHook
}

func (s *SimpleFileServerUploader) UploadFile(ctx context.Context, file *os.File) (res *FileUploadResult, err error) {
	hookSize := len(s.Hooks)
	if hookSize > 0 {
		for hookI := range s.Hooks {
			hook := s.Hooks[hookI]
			ctx, err = hook.UploadBefore(ctx, file)
			if err != nil {
				return nil, err
			}
		}
	}
	// 执行
	res, err = s.doUpload(ctx, file)
	// 执行后
	if hookSize > 0 {
		for hookI := hookSize - 1; hookI >= 0; hookI-- {
			hook := s.Hooks[hookI]
			ctx, err = hook.UploadAfter(ctx, file, res, err)
			if err != nil {
				return nil, err
			}
		}
	}
	return res, err
}

func (s *SimpleFileServerUploader) doUpload(ctx context.Context, file *os.File) (res *FileUploadResult, err error) {
	payload := new(bytes.Buffer)
	writer := multipart.NewWriter(payload)
	part, err := writer.CreateFormFile("file", filepath.Base(file.Name()))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return nil, err
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", s.Url, payload)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body := &bytes.Buffer{}
	_, err = body.ReadFrom(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(body.String())
	var url = new(string)
	err = json.Unmarshal(body.Bytes(), url)
	if err != nil {
		return nil, err
	}
	return &FileUploadResult{Url: *url}, nil
}
