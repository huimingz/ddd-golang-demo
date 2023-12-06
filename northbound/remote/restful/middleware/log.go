package middleware

import (
	"bytes"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"

	"demo/extension/logz"
)

const MAX_LOG_CONTENT_LENGTH = 1024

func Logger(skippers ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if skipHandler(c, skippers...) {
			c.Next()
			return
		}

		start := time.Now()
		logRequest(c)
		c.Next()
		logResponse(c, start)
	}
}

func logRequest(c *gin.Context) {
	logz.Info(c.Request.Context(), "[middleware] http request",
		logz.Any("method", c.Request.Method),
		logz.Any("uri", c.Request.URL.Path),
		logz.Any("content_type", c.Request.Header.Get("Content-Type")),
		logz.Any("content_length", c.Request.ContentLength),
		logz.Any("query", c.Request.URL.RawQuery),
		logz.Any("headers", c.Request.Header),
		logz.Any("ip", c.ClientIP()),
		logz.Any("ua", c.Request.Header.Get("User-Agent")),
		logz.Any("body", getPrintableBody(c)),
	)
}

func logResponse(c *gin.Context, start time.Time) {
	logz.Info(c.Request.Context(), "[middleware] http response",
		logz.Any("status", c.Writer.Status()),
		logz.Any("method", c.Request.Method),
		logz.Any("uri", c.Request.URL.Path),
		logz.Any("latency", time.Since(start).Nanoseconds()),
	)
}

func getPrintableBody(c *gin.Context) string {
	if c.Request.ContentLength < MAX_LOG_CONTENT_LENGTH {
		return string(readBody(c))
	}
	return "Exceeded the maximum data limit"
}

func readBody(c *gin.Context) []byte {
	body, _ := c.GetRawData()
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	return body
}
