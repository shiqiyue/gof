package rediss

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRedisCodeGenerator_Gen(t *testing.T) {
	client := redis.NewClient(&redis.Options{Network: "tcp", Addr: "127.0.0.1:6379", DB: 3})
	defer client.Close()
	generator, err := NewRedisCodeGenerator(client, "", "test", "codeGeneate:", DATE_AND_REDIS_INC)
	assert.Nil(t, err)
	for i := 0; i < 10000; i++ {
		ctx := context.Background()
		val, err := generator.Gen(ctx)
		assert.Nil(t, err)
		fmt.Println(val)
	}
}
