package beer

import (
	beerDB "beerchampz/pkg/beer/db"
	"beerchampz/pkg/common"

	"github.com/jackc/pgx/v5/pgtype"
)

func mapBeerToBeerDTO(beer beerDB.Beer) Beer {
	return Beer{
		ID:        int(beer.ID),
		Name:      beer.Name.String,
		Style:     beer.Style.String,
		SubStyle:  beer.SubStyle.String,
		ABV:       beer.Abv.String,
		ShortDesc: beer.ShortDesc.String,
		Brewery:   beer.Brewery.String,
		Image:     beer.Image.String,
		Score:     int(beer.Score.Int32),
		Shop:      beer.Shop.String,
		Family:    common.Family(beer.Family.String),
	}
}

// MapRequestToInsertParams :
func MapRequestToInsertParams(requestBody Beer) beerDB.InsertBeerParams {
	return beerDB.InsertBeerParams{
		Name:      pgtype.Text{String: requestBody.Name, Valid: true},
		Style:     pgtype.Text{String: requestBody.Style, Valid: true},
		SubStyle:  pgtype.Text{String: requestBody.SubStyle, Valid: true},
		Abv:       pgtype.Text{String: requestBody.ABV, Valid: true},
		ShortDesc: pgtype.Text{String: requestBody.ShortDesc, Valid: true},
		Brewery:   pgtype.Text{String: requestBody.Brewery, Valid: true},
		Image:     pgtype.Text{String: requestBody.Image, Valid: true},
		Score:     pgtype.Int4{Int32: int32(requestBody.Score), Valid: true},
		Shop:      pgtype.Text{String: requestBody.Shop, Valid: true},
		Family:    pgtype.Text{String: string(requestBody.Family), Valid: true},
	}
}

func MapFamily(family common.Family) pgtype.Text {
	return pgtype.Text{String: string(family), Valid: true}
}

func MapToGetAllByFamilyParams(family common.Family, limit int32) beerDB.GetAllByFamilyLimitParams {
	return beerDB.GetAllByFamilyLimitParams{
		Family: pgtype.Text{String: string(family), Valid: true},
		Limit:  limit,
	}
}
