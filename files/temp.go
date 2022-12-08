package files

import (
	"github.com/hashicorp/go-uuid"
	"io/ioutil"
	"os"
	"path"
)

func SaveTempFile(content []byte, ext string) string {
	tempDir := os.TempDir()
	u, _ := uuid.GenerateUUID()
	fileName := u + ext
	filepath := path.Join(tempDir, fileName)
	_ = ioutil.WriteFile(filepath, content, os.ModePerm)
	return filepath
}
