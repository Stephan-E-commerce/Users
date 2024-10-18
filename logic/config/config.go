package config

import (
	"log"
	"os"
)

// GetDBConnection возвращает строку подключения к базе данных
func GetDBConnection() string {
	connection := os.Getenv("DB_CONNECTION")
	if connection == "" {
		log.Fatal("DB_CONNECTION environment variable is not set")
	}
	return connection
}
