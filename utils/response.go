package utils

import (
	"github.com/gin-gonic/gin"
)

func SendResponse(c *gin.Context, statusCode int, message string, data interface{}) {
    c.JSON(statusCode, gin.H{
        "status":  statusCode,
        "message": message,
        "data":    data,
    })
}

func SendError(c *gin.Context, statusCode int, errorMessage string) {
    c.JSON(statusCode, gin.H{
        "status": statusCode,
        "error":  errorMessage,
    })
}
