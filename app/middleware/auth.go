package middleware

import (
	"library/app/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// AuthMiddleware возвращает middleware для проверки JWT токена
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Получаем токен из заголовка Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "отсутствует токен авторизации"})
			return
		}

		// Проверяем формат токена
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(401, gin.H{"error": "неверный формат токена"})
			return
		}

		tokenString := parts[1]

		// Валидируем токен
		token, claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			logrus.WithError(err).Warn("Ошибка валидации токена")
			c.AbortWithStatusJSON(401, gin.H{"error": "неверный токен"})
			return
		}

		// Сохраняем данные пользователя в контексте
		c.Set("user", claims.Username)
		c.Set("token", token)

		c.Next()
	}
}
