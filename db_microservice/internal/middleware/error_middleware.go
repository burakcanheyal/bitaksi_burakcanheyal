package middleware

import (
	"bitaksi_burakcanheyal/db_microservice/internal/application"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorMapper() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next() // handler çalışır

		// Hata yok ise çık
		if len(c.Errors) == 0 {
			return
		}

		err := c.Errors.Last().Err
		errCode := err.Error()

		// ErrorMap içinde varsa map'le
		if val, ok := application.CustomErrorMap[errCode]; ok {
			c.JSON(val.Status, gin.H{
				"error":   errCode,
				"message": val.Message,
			})
			return
		}

		// Map’te yoksa INTERNAL SERVER ERROR olarak dön
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "ERR_INTERNAL",
			"message": err.Error(),
		})
	}
}
