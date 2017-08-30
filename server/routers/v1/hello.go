package v1

import (
	"github.com/gin-gonic/gin"

	"ChineseChess/server/logger"
)

// Hello echoes "hello, gin!"
func Hello(c *gin.Context) {

	logger.Debug("Hello, logger!")

	c.Header("Content-Type", "text/plain")
	c.String(200, "hello, gin!")
}
