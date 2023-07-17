package championship

import (
	"beerchampz/pkg/beer"

	"gorm.io/gorm"
)

// CreateChampionship :
func CreateChampionship(db *gorm.DB, rounds []Round) Championship {
	champ := Championship{
		Winner: beer.BeerDTO{},
		Rounds: rounds,
	}
	db.Create(&champ)
	return champ
}
