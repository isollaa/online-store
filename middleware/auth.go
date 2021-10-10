package middleware

import (
	"fmt"
	"net/http"
	"online-store/config"
	"online-store/lib/conv"
	"strings"

	"github.com/gin-gonic/gin"

	jwt "github.com/dgrijalva/jwt-go"
)

func Auth(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")
	if token == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "jwt is missing",
		})
		ctx.Abort()
		return
	}

	if len(strings.Split(token, " ")) == 2 {
		token = strings.Split(token, " ")[1]
	}

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Token), nil
	})
	if err != nil {
		vE := err.(*jwt.ValidationError)
		if vE.Errors&jwt.ValidationErrorMalformed != 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "jwt is malformed",
			})
			ctx.Abort()
			return
		}
		if vE.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "jwt is either expired or not active yet",
			})
			ctx.Abort()
			return
		}
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": fmt.Sprintf("couldn't handle jwt '%+v'", err),
		})
		ctx.Abort()
		return
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok {
		ctx.Set("user_id", conv.StringToInt64(fmt.Sprintf("%v", claims["user_id"]), 0))
	}
}
