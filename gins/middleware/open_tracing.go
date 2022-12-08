package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

// sf sampling frequency
var sf = 100

func init() {
	rand.Seed(time.Now().Unix())
}

// SetSamplingFrequency 设置采样频率
// 0 <= n <= 100
func SetSamplingFrequency(n int) {
	sf = n
}

var staticFileSuffix = []string{".css", ".js", ".woff", ".icon", ".map"}

// 判断是否是静态文件
func isStaticFile(c *gin.Context) bool {
	path := c.Request.URL.Path
	for _, fileSuffix := range staticFileSuffix {
		if strings.HasSuffix(path, fileSuffix) {
			return true
		}
	}

	return false
}

func OpenTracingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 静态文件和option请求不处理
		if isStaticFile(c) || c.Request.Method == "OPTIONS" {
			c.Next()
			return
		}
		// 提取
		spanCtx, _ := opentracing.GlobalTracer().Extract(
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(c.Request.Header),
		)
		var sp opentracing.Span
		var ctx context.Context
		operateName := "REST: " + c.Request.Method + " " + c.FullPath()
		if spanCtx == nil {
			sp, ctx = opentracing.StartSpanFromContext(c.Request.Context(), operateName)

		} else {
			sp, ctx = opentracing.StartSpanFromContext(c.Request.Context(), operateName, opentracing.ChildOf(spanCtx))

		}
		c.Request = c.Request.WithContext(ctx)
		sp = sp.SetBaggageItem("op", operateName)
		// 将链路信息写到response的头部
		_ = opentracing.GlobalTracer().Inject(sp.Context(), opentracing.HTTPHeaders, c.Writer.Header())
		defer sp.Finish()

		c.Next()
		// 收集
		statusCode := c.Writer.Status()
		// Tag
		ext.PeerHostname.Set(sp, c.Request.Host)
		ext.PeerAddress.Set(sp, c.Request.RemoteAddr)
		ext.PeerService.Set(sp, c.ClientIP())
		ext.HTTPStatusCode.Set(sp, uint16(statusCode))
		ext.HTTPMethod.Set(sp, c.Request.Method)
		ext.HTTPUrl.Set(sp, c.Request.URL.Path)
		// log
		sp.LogKV("Header:", c.Request.Header)
		// 异常取样
		if statusCode >= http.StatusInternalServerError {
			ext.Error.Set(sp, true)
			sp.LogKV("RequestURI:", c.Request.RequestURI)
			sp.LogKV("RemoteAddr:", c.Request.RemoteAddr)
		} else if rand.Intn(100) > sf {
			ext.SamplingPriority.Set(sp, 0)
		}

	}
}
