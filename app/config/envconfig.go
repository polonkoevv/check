package config

import (
	"fmt"
	"os"

	"library/app/models"

	"github.com/joho/godotenv"
)

func LoadEnvConfig(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		fmt.Println("Ошибка загрузки конфигурации: ", err)
	}
	return nil
}

func GetEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func GetEnvConfig(path string) *models.Config {
	if path == "" {
		LoadEnvConfig(".env")
	} else {
		LoadEnvConfig(path)
	}
	return &models.Config{
		AppPort:    GetEnv("APP_PORT", "8080"),
		DBHost:     GetEnv("DB_HOST", "db"),
		DBUser:     GetEnv("DB_USER", "postgres"),
		DBPassword: GetEnv("DB_PASSWORD", "efimka48"),
		DBName:     GetEnv("DB_NAME", "library"),
		DBPort:     GetEnv("DB_PORT", "5432"),
	}
}
