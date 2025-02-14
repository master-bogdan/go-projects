package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/master-bogdan/event-booking/models"
	"github.com/master-bogdan/event-booking/utils"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})

		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server error",
		})

		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
	})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})

		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})

		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not authenticate user",
		})

		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}
