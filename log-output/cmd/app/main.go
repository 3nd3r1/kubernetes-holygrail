package main

import (
	"fmt"
	"hash/fnv"
	"io"
	"net/http"
	"strings"
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
		res, err := http.Get("http://ping-pong-svc:2345/pingpong")
		if err != nil {
			ctx.String(500, err.Error())
			return
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			ctx.String(res.StatusCode, res.Status)
			return
		}

		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			ctx.String(500, err.Error())
			return
		}

		bodyString := string(bodyBytes)
		requests := strings.Split(bodyString, " ")[1]

		ctx.String(200, timestamp+" "+hash(timestamp)+"\n"+"Ping / Pongs: "+requests)
	})

	go func() {
		for {
			timestamp = time.Now().Format(time.RFC3339)
			fmt.Println(timestamp + " " + hash(timestamp))
			time.Sleep(5 * time.Second)
		}
	}()

	router.Run("0.0.0.0:8080")
}
