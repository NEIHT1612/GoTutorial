package middlewares

import (
	"net/http"

	"example.com/main/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context){
	//Authorize before add event
	token := context.Request.Header.Get("Authorization")
	if token == ""{
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}
	
	//Verify token
	userId, err := utils.VerifyToken(token)
	if err != nil{
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	context.Set("userId", userId)
	context.Next()
}