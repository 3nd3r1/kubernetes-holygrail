package database

import (
	"database/sql"
	"fmt"
	"time"

	"todo-project-backend/internal/config"
	"todo-project-backend/internal/logger"
)

var Connection *sql.DB
var IsReady bool = false

func Init() error {
	var err error

	Connection, err = sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.Config.PostgresUser, config.Config.PostgresPassword,
		config.Config.PostgresHost, config.Config.PostgresPort,
		config.Config.PostgresDatabase))

	if err != nil {
		return err
	}

	go func() {
		for {
			err = Connection.Ping()
			if err == nil {
				break
			}
			logger.Logger.Info("Failed connecting to database, retrying in 10 seconds...")
			time.Sleep(10 * time.Second)
		}

		logger.Logger.Info("Connected to database")
		logger.Logger.Info("Running migrations")

		_, err = Connection.Query("CREATE TABLE IF NOT EXISTS todos (id SERIAL PRIMARY KEY, title TEXT, completed BOOLEAN)")
		if err == nil {
			logger.Logger.Info("Migrations complete")
		}

		IsReady = true
	}()

	return err
}
