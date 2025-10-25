package helpers

import (
	"github.com/gin-gonic/gin"
)

func Respond(c *gin.Context, status int, data any) {
	c.JSON(status, data)
}

func NoContent(c *gin.Context, status int) {
	c.Status(status)
}

func Error(c *gin.Context, status int, message string) {
	Respond(c, status, gin.H{"error": message})
}
