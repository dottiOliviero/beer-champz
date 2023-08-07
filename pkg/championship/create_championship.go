package championship

import (
	"context"
)

// CreateChampionship :
func CreateChampionship(championshipRepository Repository, rounds []Round) (Championship, error) {
	r, err := mapRoundsToDB(rounds)
	if err != nil {
		return Championship{}, err
	}
	c, err := championshipRepository.InsertChampionship(context.Background(), r)
	if err != nil {
		return Championship{}, err
	}
	return c, nil
}
