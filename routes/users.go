package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sambasivareddy-ch/rest_api_go/models"
	"github.com/sambasivareddy-ch/rest_api_go/utils"
)

func createUserRoute(ctx *gin.Context) {
	var usr models.User

	if err := ctx.ShouldBindJSON(&usr); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "unable to parse the request",
		})
		return
	}

	if err := usr.Save(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to create the user",
			"user":    usr,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "A new user created successfully",
	})
}

func userLoginHandler(ctx *gin.Context) {
	var usrCredentials models.User

	if err := ctx.ShouldBindJSON(&usrCredentials); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "unable to parse the request",
		})
		return
	}

	userId, userHashedPassword, err := models.GetUserIDAndPasswordByEmail(usrCredentials.EMAIL)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to find the user with given email",
		})
		return
	}

	// Compare hash password & user entered password, if not matched err will be returned
	err = utils.ComparePasswordWithHashPassword(userHashedPassword, usrCredentials.PASSWORD)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "bad credentials",
		})
		return
	}

	// Credential matched, now generate JWT token for restricted access / authorization
	userToken, err1 := utils.CreateToken(usrCredentials.EMAIL, userId)

	if err1 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to generate the token",
		})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"message": "LoggedIn successfully",
		"token":   userToken,
	})
}
