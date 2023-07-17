package beer

import "gorm.io/gorm"

// BeerDTO model
type BeerDTO struct {
	ID        uint    `json:"id"`
	Name      string  `json:"name"`
	Style     string  `json:"style"`
	SubStyle  string  `json:"sub_style"`
	ABV       float32 `json:"abv"` // alcool
	ShortDesc string  `json:"short_desc"`
	Brewery   string  `json:"brewery"`
	Image     string  `json:"image"`
	Score     int     `json:"score"`
}

// Beer model of beer
type Beer struct {
	gorm.Model
	Name      string
	Style     string
	SubStyle  string
	ABV       float32
	ShortDesc string
	Brewery   string
	Image     string
	Score     int
}

// IsInSlice : check item in slice
func (b BeerDTO) IsInSlice(beers []BeerDTO) bool {
	for _, elem := range beers {
		if elem.ID == b.ID {
			return true
		}
	}
	return false
}

func (b *Beer) toBeerDTO() BeerDTO {
	return BeerDTO{
		ID:        b.ID,
		Name:      b.Name,
		Style:     b.Style,
		SubStyle:  b.SubStyle,
		ABV:       b.ABV,
		ShortDesc: b.ShortDesc,
		Brewery:   b.Brewery,
		Image:     b.Image,
		Score:     b.Score,
	}
}
