package main

import (
	"beerchampz/pkg/beer"
	"beerchampz/pkg/championship"
	"beerchampz/pkg/config"
	"beerchampz/pkg/db"
	"beerchampz/pkg/router"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	conf := config.Get()
	database := db.GetDb(conf)

	r := router.New()

	r.Static("/assets", "./views/assets")
	r.Static("/dist", "./dist")
	r.LoadHTMLGlob("./views/*.html")

	apiRouter := r.Group("/api")
	beer.AddRouter(conf, database, apiRouter)
	championship.AddRouter(conf, database, apiRouter)
	viewApiRouter := r.Group("")
	viewApiRouter.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusFound, "/beer-champz")
	})
	viewApiRouter.GET("/beer-champz", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", nil)
	})
	viewApiRouter.GET("/about", func(ctx *gin.Context) {
		ctx.HTML(200, "about.html", nil)
	})
	championship.AddViewsRouter(conf, database, viewApiRouter)
	beer.AddViewsRouter(conf, database, viewApiRouter)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
