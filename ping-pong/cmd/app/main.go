package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	requests := 0

	router := gin.Default()
	router.GET("/pingpong", func(ctx *gin.Context) {
		ctx.String(200, "pong "+fmt.Sprint(requests))
		requests++
	})
	router.Run("0.0.0.0:8080")
}
