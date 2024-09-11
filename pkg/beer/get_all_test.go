package beer

import (
	beerDB "beerchampz/pkg/beer/db"
	"beerchampz/pkg/common"
	"beerchampz/pkg/config"
	"beerchampz/pkg/db"
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
)

var beers = []Beer{
	{ID: 1, Name: "foo1", Family: common.IPA},
	{ID: 2, Name: "foo2", Family: common.FRUMENTO},
	{ID: 3, Name: "foo3", Family: common.IPA},
	{ID: 4, Name: "foo4", Family: common.IPA},
	{ID: 5, Name: "foo5", Family: common.FRUMENTO},
	{ID: 6, Name: "foo6", Family: common.IPA},
	{ID: 7, Name: "foo7", Family: common.FRUMENTO},
	{ID: 8, Name: "foo8", Family: common.IPA},
}

var database = db.GetDb(config.GetTestConfig())
var beerRepository = NewRepository(database)

func setupBeersInDB(t *testing.T) func(t *testing.T) {
	for _, beer := range beers {
		beerRepository.InsertBeer(context.Background(), beerDB.InsertBeerParams{
			Name:   pgtype.Text{String: beer.Name, Valid: true},
			Family: pgtype.Text{String: string(beer.Family), Valid: true},
		})
	}
	return func(t *testing.T) {
		db.ExecuteStatement(context.Background(), database, "DELETE FROM beer where name like '%foo%'")
	}
}
func TestGetAllBeers(t *testing.T) {
	teardownBeersInDB := setupBeersInDB(t)
	defer teardownBeersInDB(t)

	resultBeers, err := beerRepository.GetAll(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, 8, len(resultBeers))
}
