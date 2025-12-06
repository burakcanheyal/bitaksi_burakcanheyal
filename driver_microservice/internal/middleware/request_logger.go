package middleware

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

var logFile *os.File

func init() {
	os.MkdirAll("logs", 0755)

	var err error
	logFile, err = os.OpenFile("logs/requests.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("cannot open log file:", err)
	}
}

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {

		start := time.Now()

		c.Next()

		entry := fmt.Sprintf(
			"%s | METHOD=%s | PATH=%s | STATUS=%d | IP=%s | LATENCY=%s\n",
			time.Now().Format("2006-01-02 15:04:05"),
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			c.ClientIP(),
			time.Since(start),
		)

		// Console
		log.Print(entry)

		// File
		logFile.WriteString(entry)
	}
}
