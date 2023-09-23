package beer

import (
	beerDB "beerchampz/pkg/beer/db"
	"context"
	"fmt"
)

// CreateBeer :
func CreateBeer(beerRepository Repository, params beerDB.InsertBeerParams) (int32, error) {
	newBeerID, err := beerRepository.InsertBeer(context.Background(), params)
	if err != nil {
		return 0, fmt.Errorf("An error occurred while creting the beer %v", err)
	}
	return newBeerID, nil
}

func UpdateBeerScore(beerRepository Repository, id int32) (Beer, error) {
	beer, err := beerRepository.UpdateBeerScore(context.Background(), id)
	if err != nil {
		return Beer{}, fmt.Errorf("An error occurred while creting the beer %v", err)
	}
	return beer, nil
}
