// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Beer struct {
	ID        int32
	Name      pgtype.Text
	Style     pgtype.Text
	SubStyle  pgtype.Text
	Abv       pgtype.Text
	ShortDesc pgtype.Text
	Brewery   pgtype.Text
	Image     pgtype.Text
	Score     pgtype.Int4
}
