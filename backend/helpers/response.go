package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Respond(c *gin.Context, status int, data any) {
	c.JSON(status, data)
}

func Error(c *gin.Context, status int, message string) {
	Respond(c, status, gin.H{"error": message})
}

func ValidationError(c *gin.Context, err error) {
	Error(c, http.StatusBadRequest, err.Error())
}
