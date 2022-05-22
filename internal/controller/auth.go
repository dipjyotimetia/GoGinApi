package controller

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// AuthMiddleware checks that token is valid, see https://godoc.org/github.com/dgrijalva/jwt-go#example-Parse--Hmac
func AuthMiddleware(ctx *gin.Context, jwtKey []byte) (jwt.MapClaims, bool) {
	// obtain session token from the requests cookies
	ck, err := ctx.Request.Cookie("token")
	if err != nil {
		log.Fatalf(err.Error())
		return nil, false
	}

	// Get the JWT string from the cookie
	tokenString := ck.Value

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) { //nolint:staticcheck,ineffassign
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	}
	return nil, false
}
