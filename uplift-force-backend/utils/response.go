package utils

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code" example:"200"`
	Message string      `json:"message" example:"操作成功"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty" example:"错误详情"`
}

func SuccessResponse(c *gin.Context, message string, data interface{}) {
	c.JSON(200, Response{
		Code:    200,
		Message: message,
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, code int, message string, error string) {
	response := Response{
		Code:    code,
		Message: message,
	}
	if error != "" {
		response.Error = error
	}
	c.JSON(code, response)
}
