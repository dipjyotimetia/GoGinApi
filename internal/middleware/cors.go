package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Cors adding cors
func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost")
		ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		ctx.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if ctx.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			ctx.AbortWithStatus(200)
		} else {
			ctx.Next()
		}
	}
}
