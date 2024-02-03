package storage

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/samber/lo"

	"github.com/triple0zero/lets-movie/internal/model"
)

type MoviePostgresStorage struct {
	db *sqlx.DB
}

func NewMovieStorage(db *sqlx.DB) *MoviePostgresStorage {
	return &MoviePostgresStorage{db: db}
}

func (s *MoviePostgresStorage) Movies(ctx context.Context) ([]model.Movie, error) {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var movies []dbMovie
	if err := conn.SelectContext(ctx, &movies, `SELECT * FROM movies`); err != nil {
		return nil, err
	}

	return lo.Map(movies, func(movie dbMovie, _ int) model.Movie { return model.Movie(movie) }), nil
}

//func (s *MoviePostgresStorage) MovieByID(ctx context.Context, id int64) (*model.Movie, error) {
//	conn, err := s.db.Connx(ctx)
//	if err != nil {
//		return nil, err
//	}
//	defer conn.Close()
//
//	var movie dbMovie
//	if err := conn.GetContext(ctx, &movie, `SELECT * FROM movies WHERE id = $1`, id); err != nil {
//		return nil, err
//	}
//
//	return (*model.Movie)(&movie), nil
//}

func (s *MoviePostgresStorage) Add(ctx context.Context, movie model.Movie) (int64, error) {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	var id int64

	row := conn.QueryRowxContext(
		ctx,
		`INSERT INTO movies (name, url, description, kprating, imdbrating)
					VALUES ($1, $2, $3, $4, $5) RETURNING id;`,
		movie.Name, movie.Url, movie.Description, movie.KpRating, movie.ImdbRating,
	)

	if err := row.Err(); err != nil {
		return 0, err
	}

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (s *MoviePostgresStorage) Delete(ctx context.Context, id int64) error {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	if _, err := conn.ExecContext(ctx, `DELETE FROM movies WHERE id = $1`, id); err != nil {
		return err
	}
	return nil
}

type dbMovie struct {
	ID          int64     `db:"id"`
	Name        string    `db:"name"`
	Url         string    `db:"url"`
	Description string    `db:"description"`
	KpRating    float64   `db:"kprating"`
	ImdbRating  float64   `db:"imdbrating"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
