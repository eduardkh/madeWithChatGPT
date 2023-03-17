package main

import (
	"fmt"
	"log"
	"time"

	"Singleton_Pattern/config"
	"Singleton_Pattern/database"
	"Singleton_Pattern/logging"
)

func main() {
	// Get the configuration data
	config := config.GetConfig()

	// Use the logging and database modules
	loggingModule := logging.GetLoggingModule(config.LogLevel)
	databaseModule := database.GetDatabaseMock()

	// Update the status every second
	for {
		status := fmt.Sprintf("[LOG] Logging level: %s\n[DB] Database URL: %s\n[DB] Max connections: %d\n", config.LogLevel, config.DatabaseURL, config.MaxConnections)
		loggingModule.Log(status)

		result := databaseModule.Query("SELECT * FROM users")
		log.Println(status, result)

		time.Sleep(time.Second)
	}
}
