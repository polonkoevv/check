// main.go
package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// Вместо абсолютных путей используем относительные
	"library/app/config"
	"library/app/db"
	"library/app/handlers"
	"library/app/middleware"
	"library/app/models"

	// Для Swagger
	_ "library/app/docs"
)

// @title Library API
// @version 1.0
// @description API для управления библиотекой
// @host localhost:8080
// @BasePath /api
func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
		PrettyPrint:     true,
	})
	var configPath string
	var configType string
	flag.StringVar(&configPath, "config", "", "Путь к конфигурационному файлу")
	flag.StringVar(&configType, "config-type", "env", "Тип конфигурации (yaml или env)")
	flag.Parse()

	var cfg *models.Config

	if configType == "" || configType == "env" || configPath == "" {
		cfg = config.GetEnvConfig(configPath)
	} else if configType != "yaml" && configType != "env" {
		log.Fatalf("Неверный тип конфигурации: %s", configType)
	} else {
		cfg = config.GetYamlConfig(configPath)
	}

	logLevel := logrus.InfoLevel
	if cfg.Env == "dev" {
		gin.SetMode(gin.DebugMode)
		logLevel = logrus.DebugLevel
	} else if cfg.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
		logLevel = logrus.InfoLevel
	}

	// Настройка логгера

	log.SetLevel(logLevel)

	// Открываем файл для логов
	file, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Warn("Не удалось открыть файл логов, логирование только в консоль")
	} else {
		// Используем io.MultiWriter для записи и в файл, и в консоль
		mw := io.MultiWriter(os.Stdout, file)
		log.SetOutput(mw)
	}

	// В продакшене используем logrus.InfoLevel
	if gin.Mode() == gin.ReleaseMode {
		log.SetLevel(logrus.InfoLevel)
	} else {
		log.SetLevel(logrus.DebugLevel)
	}

	// Получение конфигурации из файла
	// cfg := config.GetYamlConfig("./config.yaml")

	// Инициализация соединения с базой данных
	database, err := db.InitDB(cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort, log)
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	// Автомиграция моделей
	err = database.AutoMigrate(&models.Book{}, &models.User{}, &models.Lending{})
	if err != nil {
		log.Fatalf("Ошибка миграции моделей: %v", err)
	}

	// В продакшен режиме отключаем вывод в консоль Gin
	if gin.Mode() == gin.ReleaseMode {
		gin.DisableConsoleColor()
	}

	// Инициализация роутера Gin
	r := gin.New() // Используем gin.New() вместо gin.Default()

	// Подключаем middleware
	r.Use(middleware.LoggerMiddleware(log)) // Наш кастомный логгер
	r.Use(gin.Recovery())                   // Восстановление после паники

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	// Группа API маршрутов
	api := r.Group("/api")
	{
		// Инициализация обработчика книг
		bookHandler := handlers.NewBookHandler(database)
		userHandler := handlers.NewUserHandler(database)

		// Маршруты для книг
		api.GET("/books", bookHandler.GetAllBooks)
		api.GET("/books/:id", bookHandler.GetBook)
		api.POST("/books", bookHandler.CreateBook)
		api.PUT("/books/:id", bookHandler.UpdateBook)
		api.DELETE("/books/:id", bookHandler.DeleteBook)
		api.GET("/protected", middleware.AuthMiddleware(), func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Authorized access"})
		})

		api.POST("/user", userHandler.CreateUser)
		api.POST("/login", userHandler.Login)
		api.GET("/user/:username", userHandler.GetUserByUsername)

	}

	// Подключение Swagger документации
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Запуск сервера на порту 8080
	log.WithField("port", cfg.AppPort).Info("Запуск сервера")
	if err := r.Run(fmt.Sprintf(":%s", cfg.AppPort)); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
