package helper

import (
	"app/app/util/jwt"

	"github.com/gin-gonic/gin"
)

func GetUserByToken(ctx *gin.Context) (*jwt.Claims, error) {
	claims, exist := ctx.Get("claims")
	if !exist {
		return nil, nil
	}
	if claim, ok := claims.(*jwt.Claims); ok {
		return claim, nil
	}

	return nil, nil
}

func SetUserInClaims(ctx *gin.Context, user *jwt.Claims) {
	if user != nil {
		ctx.Set("claims", user)
	} else {
		ctx.Set("claims", nil)
	}
}
