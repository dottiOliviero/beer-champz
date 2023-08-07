package championship

import (
	championshipDB "beerchampz/pkg/championship/db"
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/exp/slices"
)

// SetRoundWinner :
func SetRoundWinner(championshipRepository Repository, championshipID int32, roundID string, winnerID int32) (*Championship, error) {

	c, err := championshipRepository.GetByID(context.Background(), championshipID)
	if err != nil {
		return nil, fmt.Errorf("unble to retrieve the champioship: %v", err)
	}

	if c.WinnerID != 0 {
		fmt.Printf("the championship %d already has a winner %d", c.ID, c.WinnerID)
		return &c, nil
	}

	for i, round := range c.Rounds {
		if round.ID == roundID {
			if round.WinnerID == -1 {
				if slices.Contains(round.getParticipantIds(), int(winnerID)) {
					c.Rounds[i].WinnerID = winnerID
					moveToNextRound(c, round, int(winnerID))
					if round.Kind == Final {
						c.WinnerID = winnerID
					}
					break
				} else {
					return nil, fmt.Errorf("the winer %d is not among the participants of round %s", winnerID, roundID)
				}
			} else {
				return nil, fmt.Errorf("the winner has already been set for the round %s, for championship %d", roundID, championshipID)
			}
		}
	}
	rounds, err := mapRoundsToDB(c.Rounds)
	if err != nil {
		return nil, fmt.Errorf("error while mapping rounds to db: %v", err)
	}
	args := championshipDB.UpdateChampionshipParams{
		ID:       championshipID,
		Winnerid: pgtype.Int4{Int32: int32(c.WinnerID), Valid: true},
		Rounds:   rounds,
	}
	cc, err := championshipRepository.UpdateChampionship(context.Background(), args)
	if err != nil {
		return nil, fmt.Errorf("unble to update the champioship: %v", err)
	}

	return &cc, nil
}

func moveToNextRound(championship Championship, round Round, winnerID int) {
	if round.Kind == Final {
		return
	}
	winner := round.getWinnerByID(winnerID)
	for i, r := range championship.Rounds {
		if r.Kind == round.Kind-1 && strings.Contains(r.ID, round.ID) {
			championship.Rounds[i].Beers = append(championship.Rounds[i].Beers, winner)
		}
	}
}
