package restful

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"

	"demo/config"
	"demo/extension/logz"
)

func run(lc fx.Lifecycle, conf *config.Schema, engine *gin.Engine) {
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", conf.HTTP.Host, conf.HTTP.Port),
		Handler: engine,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go listenAndServe(ctx, srv)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return shutdown(ctx, srv)
		},
	})
}

func listenAndServe(ctx context.Context, srv *http.Server) {
	logz.Info(ctx, fmt.Sprintf("[restful] listening and serving HTTP on %s", srv.Addr))
	if err := srv.ListenAndServe(); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			logz.Info(ctx, "[restful] service graceful shutdown")
			return
		}
		logz.Error(
			ctx,
			"[restful] service shutdown failure",
			logz.Any("listen_address", srv.Addr),
			logz.Any("err", err),
		)
	}
}

func shutdown(ctx context.Context, srv *http.Server) error {
	logz.Info(ctx, "[restful] received shutdown signal")
	if err := srv.Shutdown(ctx); err != nil {
		logz.Error(ctx, "[restful] an error occurred in server forced to shutdown", logz.Any("err", err))
		return err
	}

	logz.Info(ctx, "[restful] service shutdown successfully")
	return nil
}
