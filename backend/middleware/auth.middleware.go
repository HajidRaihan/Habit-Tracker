package middleware

import (
	"gin-gonic-gorm/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(ctx *gin.Context) {
	bearerToken := ctx.GetHeader("Authorization")
	if bearerToken == "" {
		ctx.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
		return
	}

	token := strings.Replace(bearerToken, "Bearer ", "", 1)

	if token == "" {
		ctx.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
		return
	}

	claimsData, err := utils.DecodeToken(token)

	if err != nil {
		ctx.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
		return
	}

	ctx.Set("user_id", claimsData["id"])

	ctx.Next()
}
