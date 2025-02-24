package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

// Config holds the database configuration
type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

var DBConfig *Config

// LoadConfig loads the configuration from environment variables
func LoadConfig() {
	// Automatically read from environment variables
	viper.AutomaticEnv()

	// Load configuration values from environment
	DBConfig = &Config{
		DBHost:     viper.GetString("DB_HOST"),
		DBPort:     viper.GetString("DB_PORT"),
		DBUser:     viper.GetString("DB_USER"),
		DBPassword: viper.GetString("DB_PASSWORD"),
		DBName:     viper.GetString("DB_NAME"),
	}

	// Check if any required fields are missing
	if DBConfig.DBHost == "" || DBConfig.DBPort == "" || DBConfig.DBUser == "" ||
		DBConfig.DBPassword == "" || DBConfig.DBName == "" {
		log.Fatal("Missing one or more required environment variables")
	}
}

// DBConfigString returns the formatted database connection string
func DBConfigString() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		DBConfig.DBHost, DBConfig.DBPort, DBConfig.DBUser, DBConfig.DBName, DBConfig.DBPassword)
}
