package championship

import (
	beerModel "beerchampz/pkg/beer"
	"math/rand"
)

func extractParticipants(beers []beerModel.BeerDTO) []beerModel.BeerDTO {
	// extract 8 items randomly from the list
	var selectedItems []beerModel.BeerDTO
	if len(beers) < 8 {
		panic("not enough items!!")
	}
	for i := 0; i < 8; {
		idx := rand.Intn(len(beers))
		selectedItem := beers[idx]
		if !selectedItem.IsInSlice(selectedItems) {
			selectedItems = append(selectedItems, selectedItem)
			i++
		}
	}
	return selectedItems
}

func createPairings(beers []beerModel.BeerDTO) []Round {
	// builds the initial rounds randomly using a list of participants
	rounds := []Round{buildRound(Final, "ADBC"), buildRound(Semi, "AD"), buildRound(Semi, "BC")}

	for i, letter := range "ABCD" {
		var round = buildRound(Quarter, string(letter))
		round.Players = []beerModel.BeerDTO{beers[2*i], beers[2*i+1]}
		rounds = append(rounds, round)
	}
	return rounds
}

// CreateChampionshipRounds :
func CreateChampionshipRounds(beers []beerModel.BeerDTO) []Round {
	return createPairings(extractParticipants(beers))
}
