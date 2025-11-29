package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


type BaseResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}


func Success(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, BaseResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}


func Error(c *gin.Context, statusCode int, message string, err interface{}) {
	c.JSON(statusCode, BaseResponse{
		Success: false,
		Message: message,
		Error:   err,
	})
}
