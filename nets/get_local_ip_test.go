package nets

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetLocalIp(t *testing.T) {
	ip, err := GetLocalIp("192.168.50.1")
	assert.Nil(t, err)
	fmt.Println(ip)
}
