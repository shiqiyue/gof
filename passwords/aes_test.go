package passwords

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAesEncryptBase64(t *testing.T) {
	password := "wuwenadasda"
	encPassword, err := AesEncryptBase64([]byte(password), privateKey)
	assert.Nil(t, err)
	fmt.Println(encPassword)

	passwordBs, err := AesDeCryptBase64(encPassword, privateKey)

	assert.Equal(t, password, string(passwordBs))
}
