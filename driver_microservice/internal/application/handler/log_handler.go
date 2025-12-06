package handler

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type LogHandler struct{}

func NewLogHandler() *LogHandler {
	return &LogHandler{}
}

func (h *LogHandler) GetLogs(c *gin.Context) {

	data, err := os.ReadFile("logs/requests.log")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Data(http.StatusOK, "text/plain", data)
}
