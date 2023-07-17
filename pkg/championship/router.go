package championship

import (
	beer "beerchampz/pkg/beer"
	"beerchampz/pkg/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// AddRouter :
func AddRouter(conf *config.Config, db *gorm.DB, r *gin.Engine) {

	logger := conf.GetLogger()
	route := r.Group("/championship")

	route.POST("/", func(ctx *gin.Context) {
		logger.Info("retrieving all available beers")
		beers := beer.GetAllBeers(db)

		logger.Info("creating champioship rounds")
		rounds := CreateChampionshipRounds(beers)

		res := CreateChampionship(db, rounds)

		ctx.JSON(http.StatusOK, res)
	})
}
