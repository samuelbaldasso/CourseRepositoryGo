package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ExceptionHandler captura panics e erros não tratados, retornando erro 500 e logando detalhes
func ExceptionHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if rec := recover(); rec != nil {
				log.Printf("[PANIC] %v", rec)
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error": "Erro interno do servidor",
				})
			}
		}()
		c.Next()
	}
}
