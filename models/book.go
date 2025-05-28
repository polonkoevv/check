package models

import (
	"time"
	
	"gorm.io/gorm"
)

// Book представляет собой модель книги в библиотеке
type Book struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Title     string         `gorm:"size:255;not null;index" json:"title" binding:"required"`
	Author    string         `gorm:"size:255;not null;index" json:"author" binding:"required"`
	ISBN      string         `gorm:"size:20;uniqueIndex;not null" json:"isbn" binding:"required"`
	Available bool           `gorm:"default:true" json:"available"`
	Lendings  []Lending      `gorm:"foreignKey:BookID" json:"-"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // Поддержка soft delete
}

