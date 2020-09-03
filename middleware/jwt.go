package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(ctx *gin.Context) bool {
	ck, err := ctx.Request.Cookie("token")
	if err != nil {
		fmt.Print("Verify auth token")
	}
	tokenString := ck.Value

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) { //nolint:staticcheck,ineffassign
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true
	}
	return false
}
