package championship

import (
	championshipDB "beerchampz/pkg/championship/db"
	"encoding/json"
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

func getRoundById(rounds []Round, ID string) Round {
	for _, round := range rounds {
		if round.ID == ID {
			return round
		}
	}
	return Round{}
}

func mapChampionshipToEnhanced(c Championship) ChampionshipEnhanced {
	enhanced := ChampionshipEnhanced{
		ID:       c.ID,
		WinnerID: c.WinnerID,
		Quarter1: getRoundById(c.Rounds, "A"),
		Quarter2: getRoundById(c.Rounds, "B"),
		Quarter3: getRoundById(c.Rounds, "C"),
		Quarter4: getRoundById(c.Rounds, "D"),
		Semi1:    getRoundById(c.Rounds, "AD"),
		Semi2:    getRoundById(c.Rounds, "BC"),
		Final:    getRoundById(c.Rounds, "ADBC"),
	}
	return enhanced
}
