package championship

import (
	beer "beerchampz/pkg/beer"
	"beerchampz/pkg/config"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type SetRoundWinnerRequestBody struct {
	RoundID  string      `json:"roundID" form:"round"`
	WinnerID json.Number `json:"winnerID" form:"winnerID"`
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

		if err := ctx.ShouldBind(&requestBody); err != nil {
			logger.Error("error unmarshaling JSON req body", zap.Error(err))
			ctx.String(http.StatusBadRequest, "Bad Request")
			return
		}

		logger.Info("set winner for round", zap.String("roundId", requestBody.RoundID))
		winnerId, _ := requestBody.WinnerID.Int64()
		championship, err := SetRoundWinner(championshipRepository, int32(championshipID), requestBody.RoundID, int32(winnerId))
		if err != nil {
			logger.Error("error setting round winner", zap.Error(err))
			ctx.String(http.StatusInternalServerError, "Unable to set winner for the round")
			return
		}
		if championship.WinnerID != 0 {
			beer.UpdateBeerScore(beerRepository, championship.WinnerID)
		}
		ctx.HTML(200, "championship.html", mapChampionshipToEnhanced(*championship))
	})

}
