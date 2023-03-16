package main

import (
	"fmt"

	"github.com/spf13/viper"

	"Singleton_Pattern/database"
	"Singleton_Pattern/logging"
)

func main() {
	// Load the configuration file
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error loading config file:", err)
		return
	}

	// Use the logging and database modules
	loggingModule := logging.GetLoggingModule()
	loggingModule.Log("Hello, world!")
	databaseModule := database.GetDatabaseModule()
	result := databaseModule.Query("SELECT * FROM users")
	fmt.Println(result)
}
