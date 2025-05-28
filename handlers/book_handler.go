//handlers/book_handler.go
package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"library/models"
)

// BookHandler содержит методы для обработки запросов, связанных с книгами
type BookHandler struct {
	DB *gorm.DB
}

// NewBookHandler создает новый экземпляр BookHandler с заданной базой данных
func NewBookHandler(db *gorm.DB) *BookHandler {
	return &BookHandler{DB: db}
}

// GetAllBooks godoc
// @Summary Получить все книги
// @Description Получить список всех книг в библиотеке
// @Tags books
// @Produce json
// @Success 200 {array} models.Book
// @Router /api/books [get]
func (h *BookHandler) GetAllBooks(c *gin.Context) {
	var books []models.Book

	// Получаем все книги из базы данных
	result := h.DB.Find(&books)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения книг"})
		return
	}

	c.JSON(http.StatusOK, books)
}

// GetBook godoc
// @Summary Получить книгу по ID
// @Description Получить детальную информацию о книге по её ID
// @Tags books
// @Produce json
// @Param id path int true "ID книги"
// @Success 200 {object} models.Book
// @Failure 404 {object} map[string]string
// @Router /api/books/{id} [get]
func (h *BookHandler) GetBook(c *gin.Context) {
	// Получаем ID из URL параметра
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}

	var book models.Book

	// Ищем книгу по ID
	result := h.DB.First(&book, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Книга не найдена"})
		return
	}

	c.JSON(http.StatusOK, book)
}

// CreateBook godoc
// @Summary Создать новую книгу
// @Description Добавить новую книгу в библиотеку
// @Tags books
// @Accept json
// @Produce json
// @Param book body models.Book true "Данные новой книги"
// @Success 201 {object} models.Book
// @Failure 400 {object} map[string]string
// @Router /api/books [post]
func (h *BookHandler) CreateBook(c *gin.Context) {
	var book models.Book

	// Привязываем JSON запроса к структуре книги
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Создаем книгу в базе данных
	result := h.DB.Create(&book)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка создания книги"})
		return
	}

	c.JSON(http.StatusCreated, book)
}

// UpdateBook godoc
// @Summary Обновить книгу
// @Description Обновить информацию о существующей книге
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "ID книги"
// @Param book body models.Book true "Обновленные данные книги"
// @Success 200 {object} models.Book
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/books/{id} [put]
func (h *BookHandler) UpdateBook(c *gin.Context) {
	// Получаем ID из URL параметра
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}

	var book models.Book

	// Проверяем существование книги
	result := h.DB.First(&book, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Книга не найдена"})
		return
	}

	// Привязываем JSON запроса к структуре книги
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Обновляем книгу в базе данных
	h.DB.Save(&book)

	c.JSON(http.StatusOK, book)
}

// DeleteBook godoc
// @Summary Удалить книгу
// @Description Удалить книгу из библиотеки
// @Tags books
// @Produce json
// @Param id path int true "ID книги"
// @Success 204 "No Content"
// @Failure 404 {object} map[string]string
// @Router /api/books/{id} [delete]
func (h *BookHandler) DeleteBook(c *gin.Context) {
	// Получаем ID из URL параметра
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}

	var book models.Book

	// Проверяем существование книги
	result := h.DB.First(&book, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Книга не найдена"})
		return
	}

	// Удаляем книгу из базы данных (soft delete)
	h.DB.Delete(&book)

	c.Status(http.StatusNoContent)
}
