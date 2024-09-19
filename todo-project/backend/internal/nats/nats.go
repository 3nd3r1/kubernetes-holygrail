package nats

import (
	"time"
	"todo-project-backend/internal/config"
	"todo-project-backend/internal/logger"

	"github.com/nats-io/nats.go"
)

var Connection *nats.Conn
var IsReady bool = false

func Init() error {
	go func() {
		var err error
		for {
			Connection, err = nats.Connect(config.Config.NatsUrl)
			if err == nil {
				break
			}
			logger.Logger.Info("Failed to connect to NATS, retrying in 5 seconds...")
			time.Sleep(5 * time.Second)
		}
		IsReady = true
	}()

	return nil
}
