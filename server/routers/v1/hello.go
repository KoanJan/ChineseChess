package v1

import (
	"github.com/gin-gonic/gin"
)

// Hello echoes "hello, gin!"
func Hello(c *gin.Context) {

	c.String(200, "hello, gin!")
}
