package mapper

import (
	"movies-service/internal/dto"
	"movies-service/internal/models"
)

func MapToMovieDto(movie *models.Movie) *dto.MovieDto {
	genreDtos := MapToGenreDtoSlice(movie.Genres)
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

func MapToGenreDtoSlice(genreSlice []*models.Genre) []*dto.GenreDto {
	var genreDtos []*dto.GenreDto
	for _, g := range genreSlice {
		genreDtos = append(genreDtos, MapToGenreDto(g))
	}
	return genreDtos
}

func MapToSeasonDto(season *models.Season) *dto.SeasonDto {
	return &dto.SeasonDto{
		ID:          season.ID,
		Name:        season.Name,
		AirDate:     season.AirDate,
		Description: season.Description,
		MovieID:     season.MovieID,
		EpisodeDtos: nil,
	}
}

func MapToSeasonDtoSlice(seasonSlice []*models.Season) []*dto.SeasonDto {
	var seasonDtos []*dto.SeasonDto
	for _, s := range seasonSlice {
		seasonDtos = append(seasonDtos, MapToSeasonDto(s))
	}
	return seasonDtos
}

func MapToEpisodeDto(episode *models.Episode) *dto.EpisodeDto {
	return &dto.EpisodeDto{
		ID:        episode.ID,
		Name:      episode.Name,
		AirDate:   episode.AirDate,
		Runtime:   episode.Runtime,
		VideoPath: episode.VideoPath,
		SeasonID:  episode.SeasonID,
	}
}

func MapToEpisodeDtoSlice(seasonSlice []*models.Episode) []*dto.EpisodeDto {
	var episodeDtos []*dto.EpisodeDto
	for _, e := range seasonSlice {
		episodeDtos = append(episodeDtos, MapToEpisodeDto(e))
	}
	return episodeDtos
}
