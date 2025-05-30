// models/user.go
package models

import (
	"time"

	"gorm.io/gorm"
)

// User представляет собой модель пользователя библиотеки
type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"size:255;not null" json:"name" binding:"required"`
	Email     string         `gorm:"size:255;uniqueIndex;not null" json:"email" binding:"required"`
	Username  string         `gorm:"size:255;uniqueIndex;not null" json:"username" binding:"required"`
	Password  string         `gorm:"size:255;not null" json:"password" binding:"required"`
	Lendings  []Lending      `gorm:"foreignKey:UserID" json:"-"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // Поддержка soft delete
}
