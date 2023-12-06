package middleware

import (
	"github.com/gin-gonic/gin"
)

// SkipperFunc 定义中间件跳过函数
type SkipperFunc func(*gin.Context) bool

// SkipWithPathPrefix 检查请求路径是否包含指定的前缀，如果包含则跳过
func SkipWithPathPrefix(prefixes ...string) SkipperFunc {
	return func(c *gin.Context) bool {
		path := c.Request.URL.Path
		pathLen := len(path)

		for _, p := range prefixes {
			if pl := len(p); pathLen >= pl && path[:pl] == p {
				return true
			}
		}
		return false
	}
}

// NoSkipWithPathPrefix 检查请求路径是否包含指定的前缀，如果包含则不跳过
func NoSkipWithPathPrefix(prefixes ...string) SkipperFunc {
	return func(c *gin.Context) bool {
		path := c.Request.URL.Path
		pathLen := len(path)

		for _, p := range prefixes {
			if pl := len(p); pathLen >= pl && path[:pl] == p {
				return false
			}
		}
		return true
	}
}

// skipHandler 统一处理跳过函数
func skipHandler(c *gin.Context, skippers ...SkipperFunc) bool {
	for _, skipper := range skippers {
		if skipper(c) {
			return true
		}
	}
	return false
}
