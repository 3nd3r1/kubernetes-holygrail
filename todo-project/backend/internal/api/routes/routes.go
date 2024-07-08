package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Server is up and running")
	})
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})
}
