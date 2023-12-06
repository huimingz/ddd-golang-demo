package router

import (
	"github.com/gin-gonic/gin"

	"demo/northbound/remote/restful/handler"
)

func RegisterHello(router *gin.RouterGroup, hello *handler.Hello) {
	router.GET("/hello", hello.SayHello)
}
