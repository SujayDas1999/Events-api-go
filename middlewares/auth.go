package middlewares

import (
	"events-api/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorised",
		})
		return
	}

	userid, errTok := helpers.VerifyToken(token)
	if errTok != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorised",
		})
		return
	}

	context.Set("user_id", userid)
	context.Next()
}
