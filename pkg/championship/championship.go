package championship

import (
	beer "beerchampz/pkg/beer"

	"gorm.io/gorm"
)

// Championship model
type Championship struct {
	gorm.Model
	Winner beer.BeerDTO `json:"winner"`
	Rounds []Round      `json:"rounds"`
}

// RoundKind type
type RoundKind int64

// Round kind enum
const (
	Final   RoundKind = 1
	Semi    RoundKind = 2
	Quarter RoundKind = 3
)

// Round model
type Round struct {
	ID      string
	Winner  beer.BeerDTO
	Kind    RoundKind
	Players []beer.BeerDTO
}

func buildRound(kind RoundKind, id string) Round {
	return Round{
		ID:      id,
		Winner:  beer.BeerDTO{},
		Kind:    kind,
		Players: []beer.BeerDTO{},
	}
}
