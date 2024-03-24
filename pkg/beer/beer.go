package beer

import "beerchampz/pkg/common"

// Beer model of beer
type Beer struct {
	ID        int
	Name      string
	Style     string
	SubStyle  string
	ABV       string
	ShortDesc string
	Brewery   string
	Image     string
	Score     int
	Shop      string
	Family    common.Family
}

// IsInSlice : check item in slice
func (b Beer) IsInSlice(beers []Beer) bool {
	for _, elem := range beers {
		if elem.ID == b.ID {
			return true
		}
	}
	return false
}

func GetByWinnerID(beers []Beer, winnerID int) Beer {
	for _, elem := range beers {
		if elem.ID == winnerID {
			return elem
		}
	}
	return Beer{}
}
