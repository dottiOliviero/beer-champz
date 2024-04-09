package main

import (
	"beerchampz/pkg/beer"
	"beerchampz/pkg/common"
	"beerchampz/pkg/config"
	"beerchampz/pkg/db"
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func createBeerList(data [][]string, family common.Family) []beer.Beer {
	var beers []beer.Beer

	for i, line := range data {
		if i > 0 { // omit header line
			var offset = 25
			var rec beer.Beer
			for j, field := range line {
				if j == 0 {
					rec.Name = field
				} else if j == 1 {
					rec.Style = field
				} else if j == 2 {
					rec.ABV = field
				} else if j == 3 {
					rec.ShortDesc = field
				} else if j == 4 {
					rec.Brewery = field
				} else if j == 5 {
					rec.Shop = field
				}
			}
			rec.Image = fmt.Sprintf("assets/%d.png", i+offset)
			rec.Score = 0
			rec.SubStyle = ""
			rec.Family = family
			beers = append(beers, rec)
		}
	}
	return beers
}

func readData(family common.Family) []beer.Beer {
	// open file
	f, err := os.Open("./scripts/beerz-frumento.csv")
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	csvReader.Comma = ';'
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// convert records to array of structs
	return createBeerList(data, family)
}

func main() {
	conf := config.Get()
	conn := db.GetDb(conf)
	beers := readData(common.FRUMENTO)
	beerRepo := beer.NewRepository(conn)
	for _, b := range beers {
		id, err := beerRepo.InsertBeer(context.Background(), beer.MapRequestToInsertParams(b))
		if err != nil {
			fmt.Printf("error %s \n", err)
		} else {
			fmt.Printf("inserted %d\n", id)
		}
	}

	fmt.Printf("Inserted %d Beers", len(beers))
}
