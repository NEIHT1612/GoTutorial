package routes

import (
	"net/http"

	"example.com/main/models"
	"example.com/main/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context){
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message": "Can't parse JSON user"})
		return
	}
	err = user.CreateUser()
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Can't add user"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Sign up successfully"})
}

func login(context *gin.Context){
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message": "Can't parse JSON user"})
		return
	}

	err = user.ValidateCredentials()
	if err != nil{
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Can't authenticate user"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Can't create token"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successfully", "token": token})
}