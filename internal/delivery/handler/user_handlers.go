package handler

import (
	"movie-crud-app/internal/repository/models"
	"movie-crud-app/utils"
	// "movie-crud-app/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login endpoint
func (h *MovieHandler) Login(c *gin.Context) {
	var user models.User
	// Foydalanuvchi login ma'lumotlarini olish
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Foydalanuvchi nomi va parolini tekshirish
	if user.Username != "admin" || user.Password != "password" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Token yaratish
	token, err := utils.GenerateJWT(1) // 1 - foydalanuvchi ID (bu haqiqiy database'dan olinishi kerak)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	// Tokenni foydalanuvchiga yuborish
	c.JSON(http.StatusOK, gin.H{"token": token})
}
