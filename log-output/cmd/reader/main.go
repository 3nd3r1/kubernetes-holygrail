package main

import (
	"fmt"
	"hash/fnv"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func hash(s string) string {
	hasher := fnv.New32a()
	hasher.Write([]byte(s))
	return fmt.Sprint(hasher.Sum32())
}

func main() {
	var timestamp string
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.String(200, timestamp+" "+hash(timestamp))
	})

	go func() {
		for {
			data, err := os.ReadFile("/usr/src/app/data/data.txt")
			if err != nil {
				panic(err)
			}
			if timestamp != string(data) {
				timestamp = string(data)
				fmt.Println(timestamp + " " + hash(timestamp))
			}
			time.Sleep(1 * time.Second)
		}
	}()

	router.Run("0.0.0.0:8080")
}
