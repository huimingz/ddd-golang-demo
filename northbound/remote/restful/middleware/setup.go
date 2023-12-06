package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"demo/config"
	"demo/extension/metrics"
)

func Setup(engine *gin.Engine, conf *config.Schema) {
	engine.Use(Recovery())
	engine.Use(Logger(SkipWithPathPrefix("/healthz")))
	engine.Use(LogError())
	metrics.NewPrometheus(conf.Name).Use(engine)

	if conf.CORS.Enable {
		engine.Use(cors.New(cors.Config{
			AllowAllOrigins:  conf.CORS.AllowAllOrigins,
			AllowOrigins:     conf.CORS.AllowOrigins,
			AllowMethods:     conf.CORS.AllowMethods,
			AllowHeaders:     conf.CORS.AllowHeaders,
			AllowCredentials: conf.CORS.AllowCredentials,
			ExposeHeaders:    conf.CORS.ExposeHeaders,
			MaxAge:           time.Duration(conf.CORS.MaxAge) * time.Second,
			AllowWildcard:    conf.CORS.AllowWildcard,
		}))
	}
}
