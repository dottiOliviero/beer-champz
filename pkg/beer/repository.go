package beer

import (
	beerDB "beerchampz/pkg/beer/db"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type repository struct {
	// querier is an interface containing all of the
	querier beerDB.Querier
	db      *pgxpool.Pool
}

// Repository :
type Repository interface {
	Close()
	GetAll(ctx context.Context) ([]Beer, error)
	InsertBeer(ctx context.Context, params beerDB.InsertBeerParams) (int32, error)
}

// NewRepository creates a new users Repository
func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{db: db, querier: beerDB.New(db)}
}

// Close closes Repository database connection
func (r repository) Close() {
	r.db.Close()
}

func (r *repository) GetAll(ctx context.Context) ([]Beer, error) {
	beers, err := r.querier.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	var res []Beer
	for _, beer := range beers {
		res = append(res, mapBeerToBeerDTO(beer))
	}
	return res, nil
}

func (r *repository) InsertBeer(ctx context.Context, params beerDB.InsertBeerParams) (int32, error) {
	beerID, err := r.querier.InsertBeer(ctx, params)
	if err != nil {
		return 0, err
	}
	return beerID, nil
}
