package controller

import (
	"go-lang-role-based-authentication/pkg/database"
	"go-lang-role-based-authentication/pkg/helpers"
	"go-lang-role-based-authentication/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user.Token, _ = helpers.GenerateToken(user.UserID, user.Email, user.UserType)
		database.DB.Create(&user)

		c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "token": user.Token})
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var existingUser models.User
		if err := database.DB.Where("email = ?", user.Email).First(&existingUser).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}

		token, _ := helpers.GenerateToken(existingUser.UserID, existingUser.Email, existingUser.UserType)
		c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
	}
}
