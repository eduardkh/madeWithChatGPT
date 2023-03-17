package config

import (
	"sync"

	"github.com/spf13/viper"
)

// Config is a singleton object that holds configuration data
type Config struct {
	LogLevel       string
	DatabaseURL    string
	MaxConnections int
}

var configInstance *Config
var once sync.Once

// GetConfig returns the singleton instance of the Config object
func GetConfig() *Config {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
		if err := viper.ReadInConfig(); err != nil {
			panic("Error loading config file: " + err.Error())
		}
		configInstance = &Config{
			LogLevel:       viper.GetString("logLevel"),
			DatabaseURL:    viper.GetString("databaseUrl"),
			MaxConnections: viper.GetInt("maxConnections"),
		}
	})
	return configInstance
}
