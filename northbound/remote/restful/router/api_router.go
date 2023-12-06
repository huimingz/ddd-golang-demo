package router

import (
	"github.com/gin-gonic/gin"

	"demo/config"
	"demo/northbound/remote/restful/handler"
)

func NewAPIRouter(conf *config.Schema, engine *gin.Engine) *gin.RouterGroup {
	engine.GET("/healthz", handler.HealthyHandler)

	return engine.Group(conf.HTTP.APIPrefix)
}
