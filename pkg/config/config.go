package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config struct holds configuration values
type Config struct {
	MySQLUsername string
	MySQLPassword string
	MySQLHost     string
	MySQLPort     string
	MySQLDBName   string
}

// Load loads configuration from environment variables or .env file
func Load() (*Config, error) {
	// Load environment variables from .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file, using default environment variables.")
	}

	// Initialize configuration struct
	config := &Config{
		MySQLUsername: getEnv("MYSQL_USERNAME", "username"),
		MySQLPassword: getEnv("MYSQL_PASSWORD", "password"),
		MySQLHost:     getEnv("MYSQL_HOST", "localhost"),
		MySQLPort:     getEnv("MYSQL_PORT", "3306"),
		MySQLDBName:   getEnv("MYSQL_DBNAME", "dbname"),
	}

	return config, nil
}

// getEnv retrieves the value of an environment variable, with a fallback default value
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
