package sets

import (
	"fmt"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestNewStringSortSet(t *testing.T) {
	set := NewStringSortSet()
	assert.NotNil(t, set)
}

func TestStringSortSet_Add(t *testing.T) {
	set := NewStringSortSet()
	set.Add("dasda")
	set.Add("caca")
	set.Add("caca")
	set.Add("caca")
	fmt.Println(len(set.Data()))
}
