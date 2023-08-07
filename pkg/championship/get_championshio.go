package championship

import "context"

func getChampionship(championshipRepository Repository, championshipID int32) (Championship, error) {
	return championshipRepository.GetByID(context.Background(), championshipID)
}
