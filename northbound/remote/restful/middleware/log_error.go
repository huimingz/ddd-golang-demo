package middleware

import (
	"errors"

	"github.com/gin-gonic/gin"

	"demo/extension/errorx"
	"demo/extension/logz"
)

func LogError() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		for _, err := range c.Errors {
			if err.Type != gin.ErrorTypePrivate {
				continue
			}

			var ex errorx.Error
			if errors.As(err.Unwrap(), &ex) {
				switch ex.Level() {
				case errorx.LevelError:
					logz.Error(c.Request.Context(), "http error", logz.Err(ex))
				case errorx.LevelWarning:
					logz.Warn(c.Request.Context(), "http warn", logz.Err(ex))
				case errorx.LevelCritical:
					logz.Error(c.Request.Context(), "http critical", logz.Err(ex))
				case errorx.LevelInfo:
					logz.Info(c.Request.Context(), "http info", logz.Err(ex))
				case errorx.LevelDebug:
					logz.Debug(c.Request.Context(), "http debug", logz.Err(ex))
				}
			}
		}
	}
}
