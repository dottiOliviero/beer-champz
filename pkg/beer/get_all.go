package beer

import (
	"gorm.io/gorm"
)

// GetAllBeers :
func GetAllBeers(db *gorm.DB) []BeerDTO {
	var beersDB []Beer
	db.Find(&beersDB)
	var result []BeerDTO
	for _, beer := range beersDB {
		result = append(result, beer.toBeerDTO())
	}
	return result
}
