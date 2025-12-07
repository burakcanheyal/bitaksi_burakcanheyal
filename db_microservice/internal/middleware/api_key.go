package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const InternalAPIKey = "BITAKSI-DB-ACCESS-KEY-5555"

func InternalApiKeyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		key := c.GetHeader("X-INTERNAL-KEY")
		if key == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing X-INTERNAL-KEY"})
			c.Abort()
			return
		}

		if key != InternalAPIKey {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid internal api key"})
			c.Abort()
			return
		}

		c.Next()
	}
}
