package main

import (
	"beerchampz/pkg/beer"
	"beerchampz/pkg/championship"
	"beerchampz/pkg/config"
	"beerchampz/pkg/db"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	database, err := db.GetDb()
	if err != nil {
		fmt.Printf("a problem occurred with db %s", err)
		panic("failed to connect database")
	}
	database.AutoMigrate(&beer.Beer{}, &championship.Championship{})
	conf := config.Get()
	r := gin.Default()
	beer.AddRouter(conf, database, r)
	championship.AddRouter(conf, database, r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
