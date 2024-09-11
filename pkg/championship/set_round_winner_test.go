package championship

import (
	beer "beerchampz/pkg/beer"
	championshipDB "beerchampz/pkg/championship/db"
	"beerchampz/pkg/common"
	"context"
	"testing"
)

var beers = []beer.Beer{
	{ID: 1, Name: "foo1"},
	{ID: 2, Name: "foo2"},
	{ID: 3, Name: "foo3"},
	{ID: 4, Name: "foo4"},
	{ID: 5, Name: "foo5"},
	{ID: 6, Name: "foo6"},
	{ID: 7, Name: "foo7"},
	{ID: 8, Name: "foo8"},
}

type RepositoryMock struct {
	championship Championship
}

func (rm *RepositoryMock) GetByID(ctx context.Context, id int32) (Championship, error) {
	return rm.championship, nil
}

func (rm *RepositoryMock) Close() {
	return
}

func (rm *RepositoryMock) InsertChampionship(ctx context.Context, params championshipDB.InsertChampionshipParams) (Championship, error) {
	return rm.championship, nil
}

func (rm *RepositoryMock) UpdateChampionship(ctx context.Context, arg championshipDB.UpdateChampionshipParams) (Championship, error) {
	return rm.championship, nil
}

func TestSetRoundWinnerForQuarter(t *testing.T) {
	championshipRepoMock := &RepositoryMock{championship: championshipFactory()}
	championshipID := 1

	roundID := "A"
	winnerID := 2 // the rounds are built sequentially

	resultChampionship, err := setRoundWinner(championshipRepoMock, int32(championshipID), roundID, int32(winnerID))

	expected := Championship{
		ID:       1,
		WinnerID: 2,
		Rounds: []Round{
			{ID: roundID,
				WinnerID: int32(winnerID),
				Kind:     Quarter,
				Beers:    []beer.Beer{}},
		},
		Family: common.IPA,
	}

	assertNoError(t, err)

	if resultChampionship.WinnerID != 0 {
		t.Error("the winner of the championship should not be set")
	}

	assertRoundWinner(t, *resultChampionship, expected, roundID)

}

func TestSetRoundWinnerForFinal(t *testing.T) {
	championshipRepoMock := &RepositoryMock{championship: championshipFactory()}

	StoreRoundWinner(championshipRepoMock, 1, "A", 2)
	StoreRoundWinner(championshipRepoMock, 1, "B", 4)
	StoreRoundWinner(championshipRepoMock, 1, "C", 6)
	StoreRoundWinner(championshipRepoMock, 1, "D", 8)
	// Semi
	StoreRoundWinner(championshipRepoMock, 1, "AD", 2)
	StoreRoundWinner(championshipRepoMock, 1, "BC", 6)
	// Final
	resultChampionship, err := setRoundWinner(championshipRepoMock, 1, "ADBC", 2)

	assertNoError(t, err)

	if resultChampionship.WinnerID != 2 {
		t.Errorf("winner should be 2, instead is %d", resultChampionship.WinnerID)
	}

}

func assertRoundWinner(t testing.TB, got, expected Championship, roundID string) {
	t.Helper()
	var gotRound Round
	var expectedRound Round
	for _, round := range got.Rounds {
		if round.ID == roundID {
			gotRound = round
		}
	}
	for _, round := range expected.Rounds {
		if round.ID == roundID {
			expectedRound = round
		}
	}

	if gotRound.WinnerID != expectedRound.WinnerID {
		t.Errorf("error with assertRoundWinner: got %d, expected %d", gotRound.WinnerID, expectedRound.WinnerID)
	}

}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("got error %q", err)
	}
}

func championshipFactory() Championship {
	return Championship{
		ID:       1,
		WinnerID: 0,
		Rounds:   createPairings(beers),
		Family:   common.IPA,
	}
}
