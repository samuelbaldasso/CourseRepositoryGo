package middleware

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
	"fmt"
)

func AuthMiddleware() gin.HandlerFunc {
	const rawToken = "secrettoken"
	hash := sha256.Sum256([]byte(rawToken))
	hashedToken := hex.EncodeToString(hash[:])
	fmt.Println("Hashed token for validation:", hashedToken)

	return func(c *gin.Context) {
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
