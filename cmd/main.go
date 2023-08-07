package main

import (
	"beerchampz/pkg/beer"
	"beerchampz/pkg/championship"
	"beerchampz/pkg/config"
	"beerchampz/pkg/db"
	"beerchampz/pkg/router"
)

func main() {

	database := db.GetDb()
	conf := config.Get()
	r := router.New()
	r.LoadHTMLGlob("./views/*.html")
	apiRouter := r.Group("/api")
	beer.AddRouter(conf, database, apiRouter)
	championship.AddRouter(conf, database, apiRouter)
	viewApiRouter := r.Group("")
	championship.AddViewsRouter(conf, database, viewApiRouter)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
