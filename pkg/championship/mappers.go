package championship

import (
	"beerchampz/pkg/beer"
	championshipDB "beerchampz/pkg/championship/db"
	"beerchampz/pkg/common"
	"encoding/json"

	"github.com/jackc/pgx/v5/pgtype"
)

func mapDBToChampioship(championship championshipDB.Championship) (*Championship, error) {
	c := &Championship{
		ID:       int(championship.ID),
		WinnerID: championship.Winnerid.Int32,
	}
	rounds := []Round{}
	if err := json.Unmarshal(championship.Rounds, &rounds); err != nil {
		return nil, err
	}
	c.Rounds = rounds
	return c, nil
}

func mapRoundsToDB(rounds []Round) ([]byte, error) {
	r, err := json.Marshal(rounds)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func mapToChampionshipInsertParams(rounds []Round, family common.Family) (championshipDB.InsertChampionshipParams, error) {
	mappedRound, err := mapRoundsToDB(rounds)
	if err != nil {
		return championshipDB.InsertChampionshipParams{}, err
	}
	return championshipDB.InsertChampionshipParams{
		Rounds: mappedRound,
		Family: pgtype.Text{String: string(family), Valid: true},
	}, nil
}

func getRoundById(rounds []Round, ID string) Round {
	for _, round := range rounds {
		if round.ID == ID {
			return round
		}
	}
	return Round{}
}

func mapRoundToRoundEnhanced(r Round) RoundEndanched {

	if len(r.Beers) == 2 {
		isActive := r.WinnerID == -1
		return RoundEndanched{
			ID:       r.ID,
			WinnerID: r.WinnerID,
			Left:     r.Beers[0],
			Right:    r.Beers[1],
			IsActive: isActive,
		}
	} else if len(r.Beers) == 1 {
		return RoundEndanched{
			ID:       r.ID,
			WinnerID: r.WinnerID,
			Left:     r.Beers[0],
			Right:    beer.Beer{},
			IsActive: false,
		}
	}
	return RoundEndanched{
		ID:       r.ID,
		WinnerID: r.WinnerID,
		Left:     beer.Beer{},
		Right:    beer.Beer{},
		IsActive: false,
	}
}

func mapChildRoundToRoundEnhanced(r Round, parentLeft Round, parentRight Round) RoundEndanched {
	leftBeer := beer.GetByWinnerID(parentLeft.Beers, int(parentLeft.WinnerID))
	rightBeer := beer.GetByWinnerID(parentRight.Beers, int(parentRight.WinnerID))

	var isActive = false

	if r.WinnerID == -1 && len(r.Beers) == 2 {
		isActive = true
	}

	return RoundEndanched{
		ID:       r.ID,
		WinnerID: r.WinnerID,
		Left:     leftBeer,
		Right:    rightBeer,
		IsActive: isActive,
	}
}

func mapChampionshipToEnhanced(c Championship) ChampionshipEnhanced {
	quarter1 := getRoundById(c.Rounds, "A")
	quarter2 := getRoundById(c.Rounds, "B")
	quarter3 := getRoundById(c.Rounds, "C")
	quarter4 := getRoundById(c.Rounds, "D")
	semi1 := getRoundById(c.Rounds, "AD")
	semi2 := getRoundById(c.Rounds, "BC")
	final := getRoundById(c.Rounds, "ADBC")

	enhanced := ChampionshipEnhanced{
		ID:         c.ID,
		WinnerID:   c.WinnerID,
		WinnerName: final.getWinnerByID(int(c.WinnerID)).Name,
		Quarter1:   mapRoundToRoundEnhanced(quarter1),
		Quarter2:   mapRoundToRoundEnhanced(quarter2),
		Quarter3:   mapRoundToRoundEnhanced(quarter3),
		Quarter4:   mapRoundToRoundEnhanced(quarter4),
		Semi1:      mapChildRoundToRoundEnhanced(semi1, quarter1, quarter4),
		Semi2:      mapChildRoundToRoundEnhanced(semi2, quarter2, quarter3),
		Final:      mapChildRoundToRoundEnhanced(final, semi1, semi2),
	}
	return enhanced
}
