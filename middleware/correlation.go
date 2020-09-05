package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/twinj/uuid"
)

//RequestIDMiddleware requestId middleware
func RequestIDMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uuid := uuid.NewV4()
		ctx.Writer.Header().Set("X-Request-Id", uuid.String())
		ctx.Next()
	}
}
