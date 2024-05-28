package utils

import (
	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func SendErrorResponse(c *gin.Context, status int, message string) {
	c.JSON(status, ErrorResponse{Status: status, Message: message})
}

func SendSuccessResponse(c *gin.Context, status int, data interface{}) {
	c.JSON(status, data)
}
