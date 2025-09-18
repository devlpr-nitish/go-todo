package middlewares

import (
	"net/http"

	"github.com/devlpr-nitish/todo/internal/utils"
	"github.com/gin-gonic/gin"
)


func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == ""{
			c.JSON(http.StatusUnauthorized, gin.H{"error":"Unauthorized user"})
			c.Abort()
			return 
		}

		userID, err := utils.ParseTokenAndGetUserID(token)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error":"Unauthorized user"})
			c.Abort()
			return 
		}

		c.Set("user_id", userID)
		c.Next()
	}
}