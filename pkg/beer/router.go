package beer

import (
	"beerchampz/pkg/config"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// AddRouter :
func AddRouter(conf *config.Config, db *gorm.DB, r *gin.Engine) {

	logger := conf.GetLogger()
	route := r.Group("/beer")

	route.POST("/", func(ctx *gin.Context) {
		var requestBody BeerDTO

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
		res := CreateBeer(db, requestBody)
		ctx.JSON(http.StatusOK, res)
	})
}
