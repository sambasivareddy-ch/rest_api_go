package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sambasivareddy-ch/rest_api_go/models"
)

func registerTheUserHandler(ctx *gin.Context) {
	var registrationInfo models.Register

	if err := ctx.ShouldBindJSON(&registrationInfo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Unable to parse the request data",
		})
		return // Return to prevent from further processing
	}

	userId := ctx.GetInt64("userId") // Get userId from context
	registrationInfo.ID = userId

	if err1 := registrationInfo.Save(); err1 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to insert the data into registration table",
		})
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "You have been registered successfully!!",
	})
}

func unregisterUserHandler(ctx *gin.Context) {
	var registrationInfo models.Register

	if err := ctx.ShouldBindJSON(&registrationInfo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Unable to parse the request data",
		})
		return // Return to prevent from further processing
	}

	userId := ctx.GetInt64("userId") // Get userId from context
	registrationInfo.ID = userId

	if err1 := models.UnregisterToEvent(userId, registrationInfo.EventId); err1 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to unregister",
		})
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Your registration removed successfully",
	})
}
