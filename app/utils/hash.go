package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	// Проверяем минимальную длину
	if len(password) < 8 {
		return "", errors.New("пароль должен быть не менее 8 символов")
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	if len(password) == 0 || len(password) < 8 || len(hash) == 0 {
		return false
	}

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false
	}
	return true
}

var jwtSecret = []byte("supersecretkey")

// Claims представляет собой структуру данных JWT токена
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(username string) (string, error) {
	// Создаем новый токен с кастомными claims
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "library-api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ValidateToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Проверяем, что используется правильный метод подписи
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("неверный метод подписи токена")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, nil, err
	}

	// Проверяем валидность токена
	if !token.Valid {
		return nil, nil, errors.New("невалидный токен")
	}

	// Проверяем срок действия
	if claims.ExpiresAt < time.Now().Unix() {
		return nil, nil, errors.New("токен истек")
	}

	return token, claims, nil
}

// Константы для проверки пароля
const (
	minLength = 8
	maxLength = 72
)

// CheckPassword проверяет надежность пароля
func CheckPassword(password string) error {
	if len(password) < minLength {
		return errors.New("пароль должен содержать не менее 8 символов")
	}

	if len(password) > maxLength {
		return errors.New("пароль не должен превышать 72 символа")
	}

	var (
		hasUpper   bool
		hasLower   bool
		hasNumber  bool
		hasSpecial bool
	)

	for _, char := range password {
		switch {
		case char >= 'A' && char <= 'Z':
			hasUpper = true
		case char >= 'a' && char <= 'z':
			hasLower = true
		case char >= '0' && char <= '9':
			hasNumber = true
		case char >= '!' && char <= '/' ||
			char >= ':' && char <= '@' ||
			char >= '[' && char <= '`' ||
			char >= '{' && char <= '~':
			hasSpecial = true
		}
	}

	if !hasUpper {
		return errors.New("пароль должен содержать хотя бы одну заглавную букву")
	}
	if !hasLower {
		return errors.New("пароль должен содержать хотя бы одну строчную букву")
	}
	if !hasNumber {
		return errors.New("пароль должен содержать хотя бы одну цифру")
	}
	if !hasSpecial {
		return errors.New("пароль должен содержать хотя бы один специальный символ")
	}

	return nil
}
