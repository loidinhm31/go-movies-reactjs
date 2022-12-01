package repository

import (
	"database/sql"
	"movies/backend/internal/models"
)

type DatabaseRepo interface {
	Connection() *sql.DB

	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id int) (*models.User, error)

	AllMovies(genre ...int) ([]*models.Movie, error)
	OneMovieForEdit(id int) (*models.Movie, []*models.Genre, error)
	OneMovie(id int) (*models.Movie, error)
	InsertMovie(movie models.Movie) (int, error)
	UpdateMovie(movie models.Movie) error
	DeleteMovie(id int) error

	AllGenres() ([]*models.Genre, error)
	UpdateMovieGenres(id int, genreIDs []int) error
}
