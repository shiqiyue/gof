package crypts

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckHmac(t *testing.T) {
	input := `{"api_url":"w.ushirts.cn","access_token":"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJjbGllbnRfaWQiOiI0MDFjMTZjYTEzZGIzYzUyZTJjYTA1NWIyZDczNmQ1OSIsInN0YXR1c19jb2RlIjoiOHZMYy05V0ZRS2F6dHRtRnJHZVJjIn0.ZGRiZmMwMDVhY2MyMzAyYTk1ZGNkMzM2MjdmMmUxYWM4YzU0ZDhkNjRhMjc2ZGRlZWQ0MjcwNzJiN2MwYTRhMw","open_id":"86eb202355adff984d9925100870306e05173522","state":"8vLc-9WFQKazttmFrGeRc"}`
	output := `OTA3OWNjNDZkYjRmMGViNjQ3ZDVhZmQ3ZjBiNDZlYWU1Y2ZmMDgyNGZlODAyY2QyNWI5YzFmOWY2NTVlM2UzZg==`
	clientSecret := `00d2181be1945723ed480aa83f4abd101ed8f544`
	err := CheckHmac(clientSecret, []byte(input), output)
	assert.Nil(t, err)
}

func TestEncodeHmac(t *testing.T) {
	input := "{\"business_no\":\"login:captcha\",\"phones\":[\"15359600585\"],\"params\":[\"362748\"],\"timestamp\":1665653020}"

	r := EncodeHmac("71295aa6-cec2-49d9-87d0-b26d5265552c", []byte(input))
	fmt.Println(r)
}
