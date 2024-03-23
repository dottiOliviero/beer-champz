package championship

import (
	championshipDB "beerchampz/pkg/championship/db"
	"context"
)

// CreateChampionship :
func CreateChampionship(championshipRepository Repository, params championshipDB.InsertChampionshipParams) (Championship, error) {
	c, err := championshipRepository.InsertChampionship(context.Background(), params)
	if err != nil {
		return Championship{}, err
	}
	return c, nil
}
