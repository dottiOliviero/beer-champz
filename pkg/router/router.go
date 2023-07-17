package router

import (
	"github.com/gin-gonic/gin"
)

// New returns a new router with defaults such as CORS and 404 route
func New() *gin.Engine {
	r := gin.Default()
	r.Use(corsMiddleware())
	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found: " + ctx.Request.URL.Path})
	})
	return r
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
