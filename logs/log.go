package logs

import (
	"errors"
	"io"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

var logFile *os.File

func Save(fileName string) error {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return errors.Join(errors.New("Failed to open file"), err)
	}
	logFile = file
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
	log.Println("This message is logged to a file.")
	return nil
}

func CloseLogFile() {
	if logFile != nil {
		logFile.Close()
	}
}

func CustomGinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Process request
		c.Next()

		// Stop timer
		duration := time.Since(start)

		// Log format
		log.Printf("[%s] %s %s %d %s\n",
			c.Request.Method,
			c.Request.RequestURI,
			c.ClientIP(),
			c.Writer.Status(),
			duration,
		)
	}
}
