package mapper

import (
	"movies-service/internal/dto"
	"movies-service/internal/model"
	"movies-service/pkg/util"
	"time"
)

func MapToMovie(movieDto *dto.MovieDto, author string) *model.Movie {
	return &model.Movie{
		Title:       movieDto.Title,
		TypeCode:    movieDto.TypeCode,
		ReleaseDate: movieDto.ReleaseDate,
		Runtime:     movieDto.Runtime,
		MpaaRating:  movieDto.MpaaRating,
		Description: movieDto.Description,
		ImageUrl:    util.StringToSQLNullString(movieDto.ImageUrl),
		VideoPath:   util.StringToSQLNullString(movieDto.VideoPath),
		CreatedAt:   time.Now(),
		CreatedBy:   author,
		UpdatedAt:   time.Now(),
		UpdatedBy:   author,
	}
}

func MapToMovieUpdate(movieDto *dto.MovieDto, author string) *model.Movie {
	return &model.Movie{
		ID:          movieDto.ID,
		Title:       movieDto.Title,
		TypeCode:    movieDto.TypeCode,
		ReleaseDate: movieDto.ReleaseDate,
		Runtime:     movieDto.Runtime,
		MpaaRating:  movieDto.MpaaRating,
		Description: movieDto.Description,
		ImageUrl:    util.StringToSQLNullString(movieDto.ImageUrl),
		VideoPath:   util.StringToSQLNullString(movieDto.VideoPath),
		Price:       util.FloatToSQLNullFloat(float64(movieDto.Price)),
		UpdatedAt:   time.Now(),
		UpdatedBy:   author,
	}
}

func MapToGenre(genreDto *dto.GenreDto, author string) *model.Genre {
	return &model.Genre{
		ID:        genreDto.ID,
		Name:      genreDto.Name,
		TypeCode:  genreDto.TypeCode,
		CreatedBy: author,
		UpdatedBy: author,
	}
}
