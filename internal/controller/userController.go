package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"userProfile/internal/models"
	"userProfile/internal/service"
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
	if err := c.ShouldBindJSON(&user); err != nil {
		validationErrors := validation.UserRegister(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErrors})
		return
	}

	if err := h.Service.CreateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}
