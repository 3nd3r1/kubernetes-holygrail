package main

import (
	"fmt"
	"hash/fnv"
	"io"
	"net/http"
	"os"
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

	isReady := false

	pingPongUrl, ok := os.LookupEnv("PING_PONG_URL")
	if !ok {
		pingPongUrl = "http://ping-pong-svc:2345/pingpong"
	}

	go func() {
		for {
			res, err := http.Get(pingPongUrl)
			if err == nil && res.StatusCode == 200 {
				isReady = true
				break
			}
			fmt.Printf("Can't reach %v, retrying in 5 seconds...\n", pingPongUrl)
			time.Sleep(5 * time.Second)
		}
	}()

	router := gin.Default()
	router.GET("/healthz", func(ctx *gin.Context) {
		if !isReady {
			ctx.String(500, "not ready")
			return
		}
		ctx.String(200, "ok")
	})

	router.GET("/", func(ctx *gin.Context) {
		if !isReady {
			ctx.String(500, "not ready")
			return
		}

		res, err := http.Get(pingPongUrl)
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

		data, err := os.ReadFile("/usr/src/app/data/information.txt")
		if err != nil {
			ctx.String(500, err.Error())
			return
		}

		ctx.String(200, strings.Join([]string{
			"file content: " + string(data)[:len(data)-1],
			"env variable: MESSAGE=" + os.Getenv("MESSAGE"),
			timestamp + " " + hash(timestamp),
			"Ping / Pongs: " + requests}, "\n"))
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
