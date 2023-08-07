package championship

import (
	championshipDB "beerchampz/pkg/championship/db"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type repository struct {
	// querier is an interface containing all of the
	querier championshipDB.Querier
	db      *pgxpool.Pool
}

// Repository :
type Repository interface {
	Close()
	GetByID(ctx context.Context, id int32) (Championship, error)
	InsertChampionship(ctx context.Context, rounds []byte) (Championship, error)
	UpdateChampionship(ctx context.Context, arg championshipDB.UpdateChampionshipParams) (Championship, error)
}

// NewRepository creates a new users Repository
func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{db: db, querier: championshipDB.New(db)}
}

// Close closes Repository database connection
func (r repository) Close() {
	r.db.Close()
}

func (r *repository) GetByID(ctx context.Context, id int32) (Championship, error) {
	c, err := r.querier.GetByID(ctx, id)
	if err != nil {
		return Championship{}, err
	}
	championship, err := mapDBToChampioship(c)
	if err != nil {
		return Championship{}, err
	}
	return *championship, nil
}

func (r *repository) InsertChampionship(ctx context.Context, rounds []byte) (Championship, error) {
	c, err := r.querier.InsertChampionship(ctx, rounds)
	if err != nil {
		return Championship{}, err
	}
	championship, err := mapDBToChampioship(c)
	if err != nil {
		return Championship{}, err
	}
	return *championship, nil
}

func (r *repository) UpdateChampionship(ctx context.Context, arg championshipDB.UpdateChampionshipParams) (Championship, error) {
	c, err := r.querier.UpdateChampionship(ctx, arg)
	if err != nil {
		return Championship{}, err
	}
	championship, err := mapDBToChampioship(c)
	if err != nil {
		return Championship{}, err
	}
	return *championship, nil
}
