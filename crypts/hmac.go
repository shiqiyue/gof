package crypts

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"github.com/shiqiyue/gof/ferror"
)

func CheckHmac(key string, input []byte, output string) error {
	csign := EncodeHmac(key, input)
	if csign == output {
		return nil
	}
	return ferror.New("hmac不一致")
}

func EncodeHmac(key string, input []byte) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write(input)
	sha := hex.EncodeToString(h.Sum(nil))
	return base64.StdEncoding.EncodeToString([]byte(sha))
}
