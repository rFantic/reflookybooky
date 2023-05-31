package middleware

import (
	"context"
	"flookybooky/internal/util"

	"github.com/gin-gonic/gin"
)

func CookieMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), util.ContextKey{}, c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
