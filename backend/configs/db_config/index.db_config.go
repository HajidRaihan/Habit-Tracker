package db_config

import (
	"os"
	"strconv"
)

var DB_DRIVER = "postgres"
var DB_HOST = "localhost"
var DB_PORT = 5432
var DB_USER = "postgres"
var DB_PASSWORD = "postgres"
var DB_NAME = "habit_tracker_db"

func InitDatabaseConfig() {
	driverEnv := os.Getenv("DB_DRIVER")
	hostEnv := os.Getenv("DB_HOST")
	portEnv := os.Getenv("DB_PORT")
	userEnv := os.Getenv("DB_USER")
	passwordEnv := os.Getenv("DB_PASSWORD")
	nameEnv := os.Getenv("DB_NAME")

	if driverEnv != "" {
		DB_DRIVER = driverEnv
	}

	if hostEnv != "" {
		DB_HOST = hostEnv
	}

	if portEnv != "" {
		port, err := strconv.Atoi(portEnv)
		if err != nil {
			// handle error
		}
		DB_PORT = port
	}

	if userEnv != "" {
		DB_USER = userEnv
	}

	if passwordEnv != "" {
		DB_PASSWORD = passwordEnv
	}

	if nameEnv != "" {
		DB_NAME = nameEnv
	}

}
