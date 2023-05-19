package mapper

import (
	"movies-service/internal/dto"
	"movies-service/internal/models"
)

func MapToMovieDto(movie *models.Movie) *dto.MovieDto {
	genreDtos := MaptoGenreDtoSlice(movie.Genres)
	return &dto.MovieDto{
		ID:          movie.ID,
		Title:       movie.Title,
		TypeCode:    movie.TypeCode,
		ReleaseDate: movie.ReleaseDate,
		Runtime:     movie.Runtime,
		MpaaRating:  movie.MpaaRating,
		Description: movie.Description,
		ImagePath:   movie.ImagePath.String,
		VideoPath:   movie.VideoPath.String,
		CreatedAt:   movie.CreatedAt,
		UpdatedAt:   movie.UpdatedAt,
		Genres:      genreDtos,
	}
}

func MapToMovieDtoSlice(movieSlice []*models.Movie) []*dto.MovieDto {
	var movieDtos []*dto.MovieDto
	for _, m := range movieSlice {
		movieDtos = append(movieDtos, MapToMovieDto(m))
	}
	return movieDtos
}

func MapToGenreDto(genre *models.Genre) *dto.GenreDto {
	return &dto.GenreDto{
		ID:       genre.ID,
		Name:     genre.Name,
		TypeCode: genre.TypeCode,
	}
}

func MaptoGenreDtoSlice(genreSlice []*models.Genre) []*dto.GenreDto {
	var genreDtos []*dto.GenreDto
	for _, g := range genreSlice {
		genreDtos = append(genreDtos, MapToGenreDto(g))
	}
	return genreDtos
}
