package handler

import "github.com/gin-gonic/gin"

type Hello struct{}

func NewHello() *Hello {
	return &Hello{}
}

func (h *Hello) SayHello(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "hello",
	})
}
