package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// パニック発生時にエラーレスポンスを返す
func CustomRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("パニック発生: %v", err)
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error": "内部サーバーエラーが発生しました",
				})
			}
		}()
		c.Next()
	}
}
