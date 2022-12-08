package sets

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewInt64SortSet(t *testing.T) {
	set := NewInt64SortSet()
	assert.NotNil(t, set)
}

func TestInt64SortSet_Add(t *testing.T) {
	set := NewInt64SortSet()
	set.Add(1)
	set.Add(2)
	set.Add(3)
	set.Add(3)
	fmt.Println(len(set.Data()))
}
