package helper

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func JSONResponse(c *gin.Context, statusCode int, message string, data interface{}) error {
	resp := Response{
		Code:    statusCode,
		Message: message,
		Data:    data,
	}
	Log.Info(resp)
	return c.ShouldBindJSON(resp)
}
