package beer

import (
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
