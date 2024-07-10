package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	randomString := uuid.New().String()
	var lastTimestamp string

	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.String(200, lastTimestamp+": "+randomString)
	})

	go func() {
		for {
			lastTimestamp = time.Now().Format(time.RFC3339)
			fmt.Println(lastTimestamp + ": " + randomString)
			time.Sleep(5 * time.Second)
		}
	}()

	router.Run("0.0.0.0:8080")
}
