package mapper

import (
	"movies-service/internal/dto"
	"movies-service/internal/models"
	"movies-service/pkg/utils"
	"time"
)

func MapToMovie(movieDto *dto.MovieDto, author string) *models.Movie {
	return &models.Movie{
		Title:       movieDto.Title,
		TypeCode:    movieDto.TypeCode,
		ReleaseDate: movieDto.ReleaseDate,
		Runtime:     movieDto.Runtime,
		MpaaRating:  movieDto.MpaaRating,
		Description: movieDto.Description,
		ImagePath:   utils.StringToSQLNullString(movieDto.ImagePath),
		VideoPath:   utils.StringToSQLNullString(movieDto.VideoPath),
		CreatedAt:   time.Now(),
		CreatedBy:   author,
		UpdatedAt:   time.Now(),
		UpdatedBy:   author,
	}
}

func MapToMovieUpdate(movieDto *dto.MovieDto, author string) *models.Movie {
	return &models.Movie{
		Title:       movieDto.Title,
		TypeCode:    movieDto.TypeCode,
		ReleaseDate: movieDto.ReleaseDate,
		Runtime:     movieDto.Runtime,
		MpaaRating:  movieDto.MpaaRating,
		Description: movieDto.Description,
		ImagePath:   utils.StringToSQLNullString(movieDto.ImagePath),
		VideoPath:   utils.StringToSQLNullString(movieDto.VideoPath),
		UpdatedAt:   time.Now(),
		UpdatedBy:   author,
	}
}

func MapToGenre(genreDto *dto.GenreDto, author string) *models.Genre {
	return &models.Genre{
		ID:        genreDto.ID,
		Name:      genreDto.Name,
		TypeCode:  genreDto.TypeCode,
		CreatedBy: author,
		UpdatedBy: author,
	}
}
