package championship

import (
	beer "beerchampz/pkg/beer"
	common "beerchampz/pkg/common"
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

type CreateChampionshipRequestBody struct {
	Family common.Family `json:"family" form:"family"`
}

func AddViewsRouter(conf *config.Config, db *pgxpool.Pool, r *gin.RouterGroup) {
	logger := conf.GetLogger()
	route := r.Group("/championship")
	beerRepository := beer.NewRepository(db)
	championshipRepository := NewRepository(db)

	route.POST("/new", func(ctx *gin.Context) {
		var requestBody CreateChampionshipRequestBody

		if err := ctx.ShouldBind(&requestBody); err != nil {
			logger.Error("error unmarshaling JSON req body", zap.Error(err))
			ctx.String(http.StatusBadRequest, "Bad Request")
			return
		}
		beers, err := beer.GetAllBeersByStyle(beerRepository, requestBody.Family)

		if err != nil {
			ctx.String(http.StatusInternalServerError, "Unable to retrive beers by kind")
			return
		}

		logger.Info("creating champioship rounds")
		rounds := CreateChampionshipRounds(beers)
		params, err := mapToChampionshipInsertParams(rounds, requestBody.Family)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "Unable to map rounds")
			return
		}
		championship, _ := CreateChampionship(championshipRepository, params)

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
		ctx.HTML(200, "championship.html", mapChampionshipToEnhanced(championship))
	})

	route.PUT("/:id", func(ctx *gin.Context) {

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
			logger.Info("updating score of winner", zap.Int64("winnerId", winnerId))
			beer.UpdateBeerScore(beerRepository, championship.WinnerID)
		}
		ctx.HTML(200, "championship.html", mapChampionshipToEnhanced(*championship))
	})

}
