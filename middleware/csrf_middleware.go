package middleware

import "github.com/gin-gonic/gin"

func CSRFMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
