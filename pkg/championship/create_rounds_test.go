package championship_test

import (
	"beerchampz/pkg/beer"
	"beerchampz/pkg/championship"
	"testing"

	"github.com/stretchr/testify/assert"
)

var beers = []beer.Beer{
	{ID: 1, Name: "foo1"},
	{ID: 2, Name: "foo2"},
	{ID: 3, Name: "foo3"},
	{ID: 4, Name: "foo4"},
	{ID: 5, Name: "foo5"},
	{ID: 6, Name: "foo6"},
	{ID: 7, Name: "foo7"},
	{ID: 8, Name: "foo8"},
}

func TestCreateRounds(t *testing.T) {
	t.Run("should correctly create rounds with beers", func(t *testing.T) {
		res := championship.CreateChampionshipRounds(beers)
		// Championship rounds should be of length 7:
		// A    -> 1
		// B    -> 2
		// C    -> 3
		// D    -> 4
		// AB   -> 5
		// CD   -> 6
		// ABCD -> 7
		assert.Equal(t, 7, len(res))

		for i := range res {
			round := res[i]
			// First rounds should have beers set
			if round.ID == "A" || round.ID == "B" || round.ID == "C" {
				assert.Equal(t, 2, len(round.Beers))
			} else {
				// Subsequent rounds should not have beers set
				assert.Zero(t, round.Beers)
			}
		}
	})
}
