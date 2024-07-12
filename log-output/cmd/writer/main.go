package main

import (
	"time"
	"os"
)

func main() {
	for {
		timestamp := time.Now().Format(time.RFC3339)
		if err := os.WriteFile("/usr/src/app/data/data.txt", []byte(string(timestamp)), 0644); err != nil {
			panic(err)
		}
		time.Sleep(5 * time.Second)
	}

}
