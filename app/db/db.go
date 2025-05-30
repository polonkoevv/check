// db/db.go
package db

import (
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitDB устанавливает соединение с базой данных PostgreSQL
func InitDB(db_host, db_user, db_password, db_name, db_port string, log *logrus.Logger) (*gorm.DB, error) {
	// Получение параметров подключения из переменных окружения

	// Формирование строки подключения
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		db_host, db_user, db_password, db_name, db_port)

	// Ожидание готовности БД с повторами
	var db *gorm.DB
	var err error

	retryCount := 5
	retryDelay := 5 * time.Second

	for i := 0; i < retryCount; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err == nil {
			break
		}

		fmt.Printf("Не удалось подключиться к базе данных (попытка %d/%d): %v\n",
			i+1, retryCount, err)

		if i < retryCount-1 {
			fmt.Printf("Повторная попытка через %v...\n", retryDelay)
			time.Sleep(retryDelay)
		}
	}

	if err != nil {
		return nil, fmt.Errorf("после %d попыток не удалось подключиться к базе данных: %w",
			retryCount, err)
	}

	// Настройка пула соединений
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}

// getEnv получает значение переменной окружения или возвращает значение по умолчанию
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
