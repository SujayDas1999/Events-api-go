package ReturnHelper

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Return(context *gin.Context, statusCode int, message map[string]any) {
	context.JSON(statusCode, message)
}

func ReturnBadRequest(context *gin.Context, message string) {
	context.JSON(http.StatusBadRequest, gin.H{
		"BadRequest": message,
	})
}

func ReturnSuccess(context *gin.Context, message string) {
	context.JSON(http.StatusOK, gin.H{
		"Success": message,
	})
}

func ReturnSuccessWithData(context *gin.Context, data interface{}) {
	context.JSON(http.StatusOK, data)
}

func ReturnWithUnAuthoError(context *gin.Context, data interface{}) {
	context.JSON(http.StatusUnauthorized, data)
}
