//main.go
package main

import (
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// Вместо абсолютных путей используем относительные
	"library/db"

	"library/handlers"
	"library/models"

	// Для Swagger
	_ "library/docs"
)

// @title Library API
// @version 1.0
// @description API для управления библиотекой
// @host localhost:8080
// @BasePath /api
func main() {
	// Инициализация соединения с базой данных
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	// Автомиграция моделей
	err = database.AutoMigrate(&models.Book{}, &models.User{}, &models.Lending{})
	if err != nil {
		log.Fatalf("Ошибка миграции моделей: %v", err)
	}

	// Инициализация роутера Gin
	r := gin.Default()

	// Группа API маршрутов
	api := r.Group("/api")
	{
		// Инициализация обработчика книг
		bookHandler := handlers.NewBookHandler(database)

		// Маршруты для книг
		api.GET("/books", bookHandler.GetAllBooks)
		api.GET("/books/:id", bookHandler.GetBook)
		api.POST("/books", bookHandler.CreateBook)
		api.PUT("/books/:id", bookHandler.UpdateBook)
		api.DELETE("/books/:id", bookHandler.DeleteBook)
	}

	// Подключение Swagger документации
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Запуск сервера на порту 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
