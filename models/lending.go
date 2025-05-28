package models

import (
	"time"
	
	"gorm.io/gorm"
)

// Lending представляет собой модель выдачи книги пользователю
type Lending struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	BookID     uint           `gorm:"not null;index" json:"book_id"`
	Book       Book           `gorm:"foreignKey:BookID" json:"book,omitempty"`
	UserID     uint           `gorm:"not null;index" json:"user_id"`
	User       User           `gorm:"foreignKey:UserID" json:"user,omitempty"`
	BorrowDate time.Time      `json:"borrow_date"`
	ReturnDate *time.Time     `json:"return_date,omitempty"` // Может быть null, если книга не возвращена
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"` // Поддержка soft delete
}

