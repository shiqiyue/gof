package https

import (
	"net/http"
	"strings"
)

var IP_HEADERS = []string{"x-forwarded-for", "x-real-ip", "Proxy-Client-IP", "WL-Proxy-Client-IP", "HTTP_CLIENT_IP", "HTTP_X_FORWARDED_FOR"}

// 获取客户端IP
func GetClientIp(r *http.Request) string {
	result := ""
	for _, header := range IP_HEADERS {
		headValue := r.Header.Get(header)
		if headValue == "" {
			continue
		} else {
			if headValue == "[" {
				result = "127.0.0.1"
			} else if strings.ToLower(headValue) == "unknow" {
				result = strings.Split(r.RemoteAddr, ":")[0]
			} else {
				result = headValue

			}
			break
		}
	}
	if result == "" {
		result = strings.Split(r.RemoteAddr, ":")[0]
	}
	if result != "" && len(result) > 15 {
		result = strings.Split(result, ",")[0]
	}
	if result == "0:0:0:0:0:0:0:1" {
		result = "127.0.0.1"
	}
	return result
}
