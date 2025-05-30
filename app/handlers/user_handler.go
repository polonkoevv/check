package handlers

import (
	"library/app/models"
	"library/app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	DB *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{DB: db}
}

func (h *UserHandler) Login(c *gin.Context) {

	var userCreds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&userCreds); err != nil {
		return
	}
	var user models.User
	result := h.DB.Where("username = ?", userCreds.Username).First(&user)
	if result.Error == nil && utils.CheckPasswordHash(userCreds.Password, user.Password) {
		token, _ := utils.GenerateToken(userCreds.Username)
		c.JSON(http.StatusOK, gin.H{"token": token})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
	}

}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := utils.CheckPassword(user.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userPass, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user.Password = userPass

	result := h.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка добавления пользователя"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) GetUserByUsername(c *gin.Context) {
	username := c.Param("username")
	var user models.User

	result := h.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
		return
	}

	c.JSON(http.StatusOK, user)
}
