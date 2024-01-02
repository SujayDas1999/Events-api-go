package routes

import (
	"events-api/helpers"
	"events-api/helpers/ReturnHelper"
	"events-api/helpers/db"
	userModel "events-api/models"
	"events-api/models/dtos/userDto"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func CreateUser(context *gin.Context) {
	var userDto userDto.UserDto

	err := context.ShouldBindJSON(&userDto)

	emailExists := userModel.CheckIfEmailExists(userDto, db.DB)

	if emailExists == true {
		context.JSON(http.StatusOK, gin.H{
			"error": "User already exists",
		})
		return
	}

	if err != nil {
		return
	}

	savedUser, err := userModel.Save(userDto, db.DB)
	if err != nil {
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"result": "User created Successfully",
		"user":   savedUser,
	})
}

func Login(context *gin.Context) {
	var user userDto.LoginDto
	err := context.ShouldBindJSON(&user)

	if err != nil {
		ReturnHelper.ReturnBadRequest(context, "Unable to bind JSON")
		return
	}

	err = userModel.ValidateCredentials(&user, db.DB)

	if err != nil {
		ReturnHelper.ReturnBadRequest(context, "Invalid credentials")
		return
	}

	jwt, err := helpers.GenerateToken(user.Email, user.Id)

	if err != nil {
		ReturnHelper.ReturnBadRequest(context, "Bad Request")
	}

	context.JSON(http.StatusOK, gin.H{
		"result": "Login Success",
		"token":  jwt,
	})

}

func GetAll(context *gin.Context) {

	userid, _ := context.Get("user_id")

	userId := uuid.MustParse(userid.(string))

	user, _ := userModel.GetAllUserEvents(userId, db.DB)

	context.JSON(http.StatusOK, user)
}
