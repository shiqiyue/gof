package ratelimter

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func NewGinMiddleware(limitNum int) func(c *gin.Context) {
	bucket := NewFullBucket(time.Second, int64(limitNum))

	return func(c *gin.Context) {
		available := bucket.TakeAvailable(1)
		if available == 0 {
			_ = c.AbortWithError(http.StatusTooManyRequests, errors.New("too many request"))
			return
		}

		c.Next()
	}
}
