package beer

import (
	"beerchampz/pkg/common"
	"context"
	"fmt"
)

// GetAllBeers :
func GetAllBeers(beerRepository Repository) ([]Beer, error) {
	beers, err := beerRepository.GetAll(context.Background())
	if err != nil {
		return nil, fmt.Errorf("unable to retrive beers: %v", err)
	}
	return beers, nil
}

// GetAllBeersByStyle
func GetAllBeersByStyle(beerRepository Repository, family common.Family) ([]Beer, error) {

	beers, err := beerRepository.GetAllBeersByFamily(context.Background(), MapFamily(family))
	if err != nil {
		return nil, fmt.Errorf("unable to retrive beers: %v", err)
	}
	return beers, nil
}
