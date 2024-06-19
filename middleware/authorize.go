package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sambasivareddy-ch/rest_api_go/utils"
)

func AuthorizeUser(ctx *gin.Context) {
	authToken := ctx.Request.Header.Get("Authorization")

	// If authorization key missed in the header
	if authToken == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Authorization token missed",
		})
		return
	}

	userId, err := utils.VerifyToken(authToken)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Authorization failed",
		})
		return
	}

	ctx.Set("userId", userId)

	// Token is valid
	ctx.Next() // Call next handler in the path
}
