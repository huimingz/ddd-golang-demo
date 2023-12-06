package middleware

import (
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"

	"github.com/gin-gonic/gin"

	"demo/extension/datetime"
	"demo/extension/logz"
)

func Recovery() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			err := recover()
			if err == nil {
				return
			}

			httpRequest, _ := httputil.DumpRequest(ctx.Request, false)

			if isBrokenPipe(err) {
				logz.Error(ctx, ctx.Request.URL.Path, logz.Any("error", err), logz.Any("request", string(httpRequest)))
				_ = ctx.Error(err.(error))
				return
			}

			logz.Error(ctx, "[Recovery from panic]",
				logz.Any("time", datetime.Now()),
				logz.Any("error", err),
				logz.Any("request", string(httpRequest)),
				logz.Any("stack", string(debug.Stack())),
			)
			ctx.AbortWithStatus(http.StatusInternalServerError)
		}()
		ctx.Next()
	}
}

func isBrokenPipe(err any) bool {
	if ne, ok := err.(*net.OpError); ok {
		if se, ok := ne.Err.(*os.SyscallError); ok {
			if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
				return true
			}
		}
	}
	return false
}
