package fileuploads

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestSimpleFileServerUploader_UploadFile(t *testing.T) {
	uploader := &SimpleFileServerUploader{
		Client: http.DefaultClient,
		Url:    "http://120.37.177.122:61438/file/upload",
		Hooks:  nil,
	}
	file, err := os.Open("D:\\Documents\\WeChat Files\\wxid_3hbzkl27fzt712\\FileStorage\\File\\2021-07\\loc(1).md")
	assert.Nil(t, err)

	uploadResult, err := uploader.doUpload(context.Background(), file)
	assert.Nil(t, err)

	fmt.Println(uploadResult.Url)
}
