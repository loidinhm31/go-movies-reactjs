package mapper

import (
	"movies-service/internal/dto"
	"movies-service/internal/model"
)

func MapToMovieDto(movie *model.Movie) *dto.MovieDto {
	genreDtos := MapToGenreDtoSlice(movie.Genres)
	return &dto.MovieDto{
		ID:          movie.ID,
		Title:       movie.Title,
		TypeCode:    movie.TypeCode,
		ReleaseDate: movie.ReleaseDate,
		Runtime:     movie.Runtime,
		MpaaRating:  movie.MpaaRating,
		Description: movie.Description,
		ImageUrl:    movie.ImageUrl.String,
		VideoPath:   movie.VideoPath.String,
		CreatedAt:   movie.CreatedAt,
		UpdatedAt:   movie.UpdatedAt,
		Genres:      genreDtos,
	}
}

func MapToMovieDtoSlice(movieSlice []*model.Movie) []*dto.MovieDto {
	var movieDtos []*dto.MovieDto
	for _, m := range movieSlice {
		movieDtos = append(movieDtos, MapToMovieDto(m))
	}
	return movieDtos
}

func MapToGenreDto(genre *model.Genre) *dto.GenreDto {
	return &dto.GenreDto{
		ID:       genre.ID,
		Name:     genre.Name,
		TypeCode: genre.TypeCode,
	}
}

func MapToGenreDtoSlice(genreSlice []*model.Genre) []*dto.GenreDto {
	var genreDtos []*dto.GenreDto
	for _, g := range genreSlice {
		genreDtos = append(genreDtos, MapToGenreDto(g))
	}
	return genreDtos
}

func MapToSeasonDto(season *model.Season) *dto.SeasonDto {
	return &dto.SeasonDto{
		ID:          season.ID,
		Name:        season.Name,
		AirDate:     season.AirDate,
		Description: season.Description,
		MovieID:     season.MovieID,
		EpisodeDtos: nil,
	}
}

func MapToSeasonDtoSlice(seasonSlice []*model.Season) []*dto.SeasonDto {
	var seasonDtos []*dto.SeasonDto
	for _, s := range seasonSlice {
		seasonDtos = append(seasonDtos, MapToSeasonDto(s))
	}
	return seasonDtos
}

func MapToEpisodeDto(episode *model.Episode) *dto.EpisodeDto {
	return &dto.EpisodeDto{
		ID:        episode.ID,
		Name:      episode.Name,
		AirDate:   episode.AirDate,
		Runtime:   episode.Runtime,
		VideoPath: episode.VideoPath,
		SeasonID:  episode.SeasonID,
	}
}

func MapToEpisodeDtoSlice(seasonSlice []*model.Episode) []*dto.EpisodeDto {
	var episodeDtos []*dto.EpisodeDto
	for _, e := range seasonSlice {
		episodeDtos = append(episodeDtos, MapToEpisodeDto(e))
	}
	return episodeDtos
}

func MapToRatingDto(rating *model.Rating) *dto.RatingDto {
	return &dto.RatingDto{
		ID:   rating.ID,
		Code: rating.Code,
		Name: rating.Name,
	}
}

func MapToRatingDtoSlice(ratingSlice []*model.Rating) []*dto.RatingDto {
	var ratingDtos []*dto.RatingDto
	for _, r := range ratingSlice {
		ratingDtos = append(ratingDtos, MapToRatingDto(r))
	}
	return ratingDtos
}
