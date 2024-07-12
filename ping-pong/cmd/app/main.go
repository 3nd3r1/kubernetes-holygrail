package main

import (
	"fmt"
	"os"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	requests := 0
	router.GET("/pingpong", func(ctx *gin.Context) {
		ctx.String(200, "pong "+fmt.Sprint(requests))

		requests++
		if err := os.WriteFile("/usr/src/app/data/pingpong.txt", []byte(fmt.Sprint(requests)), 0644); err != nil {
			panic(err)
		}
	})
	router.Run("0.0.0.0:8080")
}
