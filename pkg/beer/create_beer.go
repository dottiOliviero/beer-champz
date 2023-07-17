package beer

import "gorm.io/gorm"

// CreateBeer :
func CreateBeer(db *gorm.DB, beer BeerDTO) BeerDTO {
	beerDB := Beer{
		Name:      beer.Name,
		Style:     beer.Style,
		SubStyle:  beer.SubStyle,
		ABV:       beer.ABV,
		ShortDesc: beer.ShortDesc,
		Brewery:   beer.Brewery,
		Image:     beer.Image,
		Score:     beer.Score,
	}
	db.Create(&beerDB)
	beer.ID = beerDB.ID
	return beer
}
