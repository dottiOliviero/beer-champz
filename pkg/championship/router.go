package championship

import (
	beer "beerchampz/pkg/beer"
	"beerchampz/pkg/config"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type SetRoundWinnerRequestBody struct {
	roundID  string
	winnerID int
}

// AddRouter :
func AddRouter(conf *config.Config, db *pgxpool.Pool, r *gin.RouterGroup) {

	logger := conf.GetLogger()
	route := r.Group("/championship")
	beerRepository := beer.NewRepository(db)
	championshipRepository := NewRepository(db)

	route.POST("/", func(ctx *gin.Context) {
		logger.Info("retrieving all available beers")
		beers, err := beer.GetAllBeers(beerRepository)

		if err != nil {
			ctx.String(http.StatusInternalServerError, "Unable to retrive beers")
			return
		}

		logger.Info("creating champioship rounds")
		rounds := CreateChampionshipRounds(beers)

		res, err := CreateChampionship(championshipRepository, rounds)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "Unable to create championship")
			return
		}

		ctx.JSON(http.StatusCreated, res)
	})

	route.POST("/:id/round/:roundId/winner/:winnerId", func(ctx *gin.Context) {
		var championshipID, _ = strconv.ParseInt(ctx.Param("id"), 10, 32)
		var roundID = ctx.Param("roundId")
		var winnerID, _ = strconv.ParseInt(ctx.Param("winnerId"), 10, 32)

		logger.Info("set winner for round", zap.String("roundId", roundID))

		championship, err := SetRoundWinner(championshipRepository, int32(championshipID), roundID, int32(winnerID))
		if err != nil {
			logger.Error("error setting round winner", zap.Error(err))
			ctx.String(http.StatusInternalServerError, "Unable to set winner for the round")
			return
		}
		ctx.JSON(http.StatusCreated, championship)
	})

	route.GET("/:id", func(ctx *gin.Context) {
		var championshipID, _ = strconv.ParseInt(ctx.Param("id"), 10, 32)

		championship, err := getChampionship(championshipRepository, int32(championshipID))
		if err != nil {
			logger.Error("error getting championship", zap.Error(err))
			ctx.String(http.StatusNotFound, "Unable to set get championship")
			return
		}
		ctx.JSON(http.StatusOK, championship)
	})
}

func AddViewsRouter(conf *config.Config, db *pgxpool.Pool, r *gin.RouterGroup) {
	logger := conf.GetLogger()
	route := r.Group("/championship")
	beerRepository := beer.NewRepository(db)
	championshipRepository := NewRepository(db)

	route.GET("/home", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", nil)
	})

	route.GET("/new", func(ctx *gin.Context) {
		beers, err := beer.GetAllBeers(beerRepository)

		if err != nil {
			ctx.String(http.StatusInternalServerError, "Unable to retrive beers")
			return
		}

		logger.Info("creating champioship rounds")
		rounds := CreateChampionshipRounds(beers)

		championship, _ := CreateChampionship(championshipRepository, rounds)

		ctx.HTML(200, "championship.html", mapChampionshipToEnhanced(championship))
	})

	route.GET("/:id", func(ctx *gin.Context) {
		var championshipID, _ = strconv.ParseInt(ctx.Param("id"), 10, 32)

		championship, err := getChampionship(championshipRepository, int32(championshipID))
		if err != nil {
			logger.Error("error getting championship", zap.Error(err))
			ctx.String(http.StatusNotFound, "Unable to set get championship")
			return
		}
		ctx.JSON(http.StatusOK, championship)
	})

	route.POST("/:id", func(ctx *gin.Context) {

		var championshipID, _ = strconv.ParseInt(ctx.Param("id"), 10, 32)
		var requestBody SetRoundWinnerRequestBody

		b, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			logger.Error("error reading req body", zap.Error(err))
			ctx.String(http.StatusBadRequest, "Bad Request")
			return
		}
		if err := json.Unmarshal(b, &requestBody); err != nil {
			logger.Error("error unmarshaling JSON req body", zap.Error(err))
			ctx.String(http.StatusBadRequest, "Bad Request")
			return
		}

		logger.Info("set winner for round", zap.String("roundId", requestBody.roundID))

		championship, err := SetRoundWinner(championshipRepository, int32(championshipID), requestBody.roundID, int32(requestBody.winnerID))
		if err != nil {
			logger.Error("error setting round winner", zap.Error(err))
			ctx.String(http.StatusInternalServerError, "Unable to set winner for the round")
			return
		}
		ctx.HTML(200, "championship.html", mapChampionshipToEnhanced(*championship))
	})

}
