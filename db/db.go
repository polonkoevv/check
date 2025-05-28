package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDB устанавливает соединение с базой данных PostgreSQL
func InitDB() (*gorm.DB, error) {
	// Обновите пароль в строке подключения
	dsn := "host=localhost user=postgres password=efimka48 dbname=library port=5432 sslmode=disable"

	// Подключение к базе данных
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
