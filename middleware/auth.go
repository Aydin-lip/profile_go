package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"userProfile/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "سرصفحه مجوز وجود ندارد یا نامعتبر است"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "رمز نامعتبر یا منقضی شده است"})
			c.Abort()
			return
		}

		// Set user ID in context
		c.Set("userID", claims.UserID)
		c.Next()
	}
}
