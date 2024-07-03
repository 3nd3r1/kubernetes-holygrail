package main

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

func main() {
	randomString := uuid.New().String()
	for {
		fmt.Println(time.Now().Format(time.RFC3339) + ": " + randomString)
		time.Sleep(5 * time.Second)
	}
}
