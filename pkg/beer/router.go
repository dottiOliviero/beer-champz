package beer

import (
	"beerchampz/pkg/config"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

// AddRouter :
func AddRouter(conf *config.Config, db *pgxpool.Pool, r *gin.RouterGroup) {

	logger := conf.GetLogger()
	route := r.Group("/beer")

	beerRepository := NewRepository(db)

	route.POST("/", func(ctx *gin.Context) {
		var requestBody Beer

		b, err := io.ReadAll(ctx.Request.Body)
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

		params := MapRequestToInsertParams(requestBody)
		res, err := CreateBeer(beerRepository, params)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "Error while creating beer")
		}
		ctx.JSON(http.StatusOK, res)
	})
}

// AddRouter :
func AddViewsRouter(conf *config.Config, db *pgxpool.Pool, r *gin.RouterGroup) {

	route := r.Group("/beer")

	beerRepository := NewRepository(db)

	route.GET("/ranking", func(ctx *gin.Context) {
		beers, err := beerRepository.GetAll(ctx)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "Error while getting beers")
			return
		}
		ctx.HTML(200, "ranking.html", beers)
	})
}
