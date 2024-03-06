package main

import (
	"beerchampz/pkg/beer"
	"beerchampz/pkg/callback"
	"beerchampz/pkg/championship"
	"beerchampz/pkg/config"
	"beerchampz/pkg/db"
	"beerchampz/pkg/login"
	"beerchampz/pkg/logout"
	"beerchampz/pkg/router"
	"beerchampz/platform/authenticator"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load the env vars: %v", err)
	}
	auth, err := authenticator.New()
	if err != nil {
		log.Fatalf("Failed to initialize the authenticator: %v", err)
	}
	conf := config.Get()
	database := db.GetDb(conf)

	r := router.New(auth)

	r.Static("/assets", "./views/assets")
	r.Static("/dist", "./dist")
	r.LoadHTMLGlob("./views/*.html")

	apiRouter := r.Group("/api")
	beer.AddRouter(conf, database, apiRouter)
	championship.AddRouter(conf, database, apiRouter)

	// AUTH
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "login.html", nil)
	})
	r.GET("/login", login.Handler(auth))
	r.GET("/callback", callback.Handler(auth))
	// r.GET("/user", user.Handler)
	r.GET("/logout", logout.Handler)

	viewApiRouter := r.Group("", router.IsAuthenticated)

	viewApiRouter.GET("/home", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", nil)
	})
	viewApiRouter.GET("/about", func(ctx *gin.Context) {
		ctx.HTML(200, "about.html", nil)
	})
	championship.AddViewsRouter(conf, database, viewApiRouter)
	beer.AddViewsRouter(conf, database, viewApiRouter)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
