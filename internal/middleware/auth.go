package middleware

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
	"fmt"
	"os"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		rawToken := os.Getenv("SECRET_KEY")
		hash := sha256.Sum256([]byte(rawToken))
		hashedToken := hex.EncodeToString(hash[:])
		fmt.Println("Hashed token for validation:", hashedToken)

		token := c.GetHeader("Authorization")
		if token == "" || !strings.HasPrefix(token, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		provided := strings.TrimPrefix(token, "Bearer ")
		if provided != hashedToken {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}
		c.Next()
	}
}
