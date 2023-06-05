package mapper

import (
	"movies-service/internal/common/dto"
	model2 "movies-service/internal/common/model"
	"movies-service/pkg/util"
	"time"
)

func MapToMovie(movieDto *dto.MovieDto, author string) *model2.Movie {
	return &model2.Movie{
		Title:       movieDto.Title,
		TypeCode:    movieDto.TypeCode,
		ReleaseDate: movieDto.ReleaseDate,
		Runtime:     movieDto.Runtime,
		MpaaRating:  movieDto.MpaaRating,
		Description: movieDto.Description,
		ImageUrl:    util.StringToSQLNullString(movieDto.ImageUrl),
		VideoPath:   util.StringToSQLNullString(movieDto.VideoPath),
		Price:       util.FloatToSQLNullFloat(movieDto.Price),
		CreatedAt:   time.Now(),
		CreatedBy:   author,
		UpdatedAt:   time.Now(),
		UpdatedBy:   author,
	}
}

func MapToMovieUpdate(movieDto *dto.MovieDto, author string) *model2.Movie {
	return &model2.Movie{
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

func MapToGenre(genreDto *dto.GenreDto, author string) *model2.Genre {
	return &model2.Genre{
		ID:        genreDto.ID,
		Name:      genreDto.Name,
		TypeCode:  genreDto.TypeCode,
		CreatedBy: author,
		UpdatedBy: author,
	}
}
