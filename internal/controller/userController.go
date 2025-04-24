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
		c.JSON(500, gin.H{"error": "تولید توکن ناموفق بود"})
		return
	}

	userResponse := models.UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Email:     user.Email,
		Phone:     user.Phone,
		Age:       user.Age,
		CreatedAt: user.CreatedAt,
	}

	// Send
	c.JSON(201, gin.H{
		"user":  userResponse,
		"token": token,
	})
}

func (h *UserControllerType) Login(c *gin.Context) {
	var req models.LoginRequest

	// Bind JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		validationErrors := validation.UserRegister(err)
		c.JSON(400, gin.H{"error": validationErrors})
		return
	}

	// Find user by username
	user, err := h.Service.LoginUser(req.Username)
	if err != nil {
		c.JSON(401, gin.H{"error": "نام کاربری یا رمز عبور نامعتبر است"})
		return
	}

	// Compare password
	if err := utils.CheckPassword(user.Password, req.Password); err != nil {
		c.JSON(401, gin.H{"error": "نام کاربری یا رمز عبور نامعتبر است"})
		return
	}

	// Generate JWT
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": "تولید توکن ناموفق بود"})
		return
	}

	// Create response
	userResponse := models.UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Email:     user.Email,
		Phone:     user.Phone,
		Age:       user.Age,
		CreatedAt: user.CreatedAt,
	}

	// Send
	c.JSON(200, gin.H{
		"token": token,
		"user":  userResponse,
	})
}
