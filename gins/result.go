package gins

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shiqiyue/gof/loggers"
	"github.com/shiqiyue/gof/results"
	"go.uber.org/zap"
	"net/http"
)

// 操作成功
func Suc(ctx *gin.Context, data interface{}) {
	ctx.JSON(200, results.Suc(data))
}

// 操作失败
func Err(ctx *gin.Context, err error, statuscode int) {
	loggers.Error(ctx.Request.Context(), fmt.Sprintf("%+v\n", err), zap.Int("StatusCode", statuscode))
	ctx.JSON(statuscode, results.Err(err))
}

func BadRequest(ctx *gin.Context, err error) {
	Err(ctx, err, http.StatusBadRequest)
}
