package championship

import (
	beer "beerchampz/pkg/beer"
)

// Championship model
type Championship struct {
	ID       int
	WinnerID int32
	Rounds   []Round
}

type ChampionshipEnhanced struct {
	ID         int
	WinnerID   int32
	WinnerName string
	Quarter1   RoundEnhanched
	Quarter2   RoundEnhanched
	Quarter3   RoundEnhanched
	Quarter4   RoundEnhanched
	Semi1      RoundEnhanched
	Semi2      RoundEnhanched
	Final      RoundEnhanched
}

type RoundEnhanched struct {
	ID       string
	WinnerID int32
	Left     beer.Beer
	Right    beer.Beer
	IsActive bool
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
	ID       string `gorm:"primaryKey"`
	WinnerID int32
	Kind     RoundKind
	Beers    []beer.Beer
}

func buildRound(kind RoundKind, id string) Round {
	return Round{
		ID:       id,
		WinnerID: -1,
		Kind:     kind,
		Beers:    []beer.Beer{},
	}
}

func (r Round) getParticipantIds() []int {
	var ids []int
	for _, beer := range r.Beers {
		ids = append(ids, beer.ID)
	}
	return ids
}

func (r Round) getWinnerByID(winnerID int) beer.Beer {
	for _, beer := range r.Beers {
		if beer.ID == winnerID {
			return beer
		}
	}
	return beer.Beer{}
}
