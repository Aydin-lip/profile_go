package controller

import (
	"github.com/gin-gonic/gin"

	"userProfile/internal/models"
	"userProfile/internal/service"
	"userProfile/utils"
	"userProfile/validation"

)

type UserControllerType struct {
	Service *service.UserServiceType
}

func UserController(s *service.UserServiceType) *UserControllerType {
	return &UserControllerType{Service: s}
}

func (h *UserControllerType) Register(c *gin.Context) {
	var user models.User
	// Validate json values
	if err := c.ShouldBindJSON(&user); err != nil {
		validationErrors := validation.UserRegister(err)
		c.JSON(400, gin.H{"error": validationErrors})
		return
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	// Set hash password in obj
	user.Password = hashedPassword

	// Create user
	if err := h.Service.CreateUser(user); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Generate JWT token using user.ID
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not generate token"})
		return
	}

	// Send
	c.JSON(201, gin.H{
		"token":   token,
		"user":    user,
	})
}
