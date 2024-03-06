package router

import (
	"beerchampz/platform/authenticator"
	"encoding/gob"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

	"github.com/gin-gonic/gin"
)

// New returns a new router with defaults such as CORS and 404 route
func New(auth *authenticator.Authenticator) *gin.Engine {
	r := gin.Default()
	gob.Register(map[string]interface{}{})

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("auth-session", store))

	// r.Use(corsMiddleware())
	// r.NoRoute(func(ctx *gin.Context) {
	// 	ctx.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found: " + ctx.Request.URL.Path})
	// })
	return r
}

// func corsMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
// 		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

// 		if c.Request.Method == "OPTIONS" {
// 			c.AbortWithStatus(204)
// 			return
// 		}

// 		c.Next()
// 	}
// }

func IsAuthenticated(ctx *gin.Context) {
	profile := sessions.Default(ctx).Get("profile")
	if profile == nil {
		ctx.Redirect(http.StatusSeeOther, "/")
	} else {
		log.Printf("profile: %v", profile)
		ctx.Next()
	}
}
