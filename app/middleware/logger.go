package middleware

import (
	"fmt"
	"math"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// LoggerMiddleware возвращает middleware для логирования запросов с помощью logrus
func LoggerMiddleware(logger *logrus.Logger) gin.HandlerFunc {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	return func(c *gin.Context) {
		// Время начала запроса
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Обработка запроса
		c.Next()

		// Время окончания запроса
		stop := time.Now()
		latency := stop.Sub(start)
		if latency > time.Minute {
			latency = latency.Round(time.Second)
		}

		// Статус ответа
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		method := c.Request.Method
		errorMessage := c.Errors.ByType(gin.ErrorTypePrivate).String()

		if raw != "" {
			path = path + "?" + raw
		}

		// Подготовка полей для логирования
		fields := logrus.Fields{
			"hostname":  hostname,
			"status":    statusCode,
			"latency":   math.Round(float64(latency.Nanoseconds()) / 1e6), // в миллисекундах
			"client_ip": clientIP,
			"method":    method,
			"path":      path,
			"errors":    errorMessage,
		}

		entry := logger.WithFields(fields)

		if len(c.Errors) > 0 {
			// Логируем как ошибку
			entry.Error(c.Errors.String())
		} else {
			msg := fmt.Sprintf("%s %s %d %s", method, path, statusCode, latency)
			if statusCode >= 500 {
				entry.Error(msg)
			} else if statusCode >= 400 {
				entry.Warn(msg)
			} else {
				entry.Info(msg)
			}
		}
	}
}
