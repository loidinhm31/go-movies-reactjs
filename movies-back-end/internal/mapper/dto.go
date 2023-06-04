package mapper

import (
	"database/sql"
	"movies-service/internal/dto"
	"movies-service/internal/model"
	"time"
)

func MapToMovieDto(movie *model.Movie, isRestrictResource bool, isPrivilege bool) *dto.MovieDto {
	genreDtos := MapToGenreDtoSlice(movie.Genres)

	// Filter video path for release date
	if isRestrictResource ||
		(!isPrivilege && time.Now().Before(movie.ReleaseDate)) {
		movie.VideoPath = sql.NullString{}
	}

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
		Price:       movie.Price.Float64,
		CreatedAt:   movie.CreatedAt,
		UpdatedAt:   movie.UpdatedAt,
		Genres:      genreDtos,
	}
}

func MapToMovieDtoSlice(movieSlice []*model.Movie) []*dto.MovieDto {
	var movieDtos []*dto.MovieDto
	for _, m := range movieSlice {
		movieDtos = append(movieDtos, MapToMovieDto(m, true, false))
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

func MapToEpisodeDto(episode *model.Episode, isRestrictResource, isPrivilege bool) *dto.EpisodeDto {
	// Filter video path for release date
	if isRestrictResource ||
		(!isPrivilege && time.Now().Before(episode.AirDate)) {
		episode.VideoPath = sql.NullString{}
	}

	var season *dto.SeasonDto
	if episode.Season != nil {
		season = MapToSeasonDto(episode.Season)
	}

	return &dto.EpisodeDto{
		ID:        episode.ID,
		Name:      episode.Name,
		AirDate:   episode.AirDate,
		Runtime:   episode.Runtime,
		VideoPath: episode.VideoPath.String,
		SeasonID:  episode.SeasonID,
		Price:     episode.Price.Float64,
		Season:    season,
	}
}

func MapToEpisodeDtoSlice(seasonSlice []*model.Episode, isRestrictResource, isPrivilege bool) []*dto.EpisodeDto {
	var episodeDtos []*dto.EpisodeDto
	for _, e := range seasonSlice {
		episodeDtos = append(episodeDtos, MapToEpisodeDto(e, isRestrictResource, isPrivilege))
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

func MapToCollectionDto(collection *model.Collection) *dto.CollectionDto {
	return &dto.CollectionDto{
		Username:  collection.Username,
		MovieID:   uint(collection.MovieID.Int64),
		EpisodeID: uint(collection.EpisodeID.Int64),
	}
}

func MapToCollectionDetailDto(collection *model.CollectionDetail) *dto.CollectionDetailDto {
	return &dto.CollectionDetailDto{
		Username:    collection.Username,
		MovieID:     collection.MovieID,
		EpisodeID:   collection.EpisodeID,
		TypeCode:    collection.TypeCode,
		Title:       collection.Title,
		SeasonName:  collection.SeasonName,
		EpisodeName: collection.EpisodeName,
		ReleaseDate: collection.ReleaseDate,
		ImageUrl:    collection.ImageUrl,
		Description: collection.Description,
		Price:       collection.Amount,
		CreatedAt:   collection.CreatedAt,
	}
}

func MapToCollectionDetailDtoSlice(collections []*model.CollectionDetail) []*dto.CollectionDetailDto {
	var collectionDtos []*dto.CollectionDetailDto
	for _, c := range collections {
		collectionDtos = append(collectionDtos, MapToCollectionDetailDto(c))
	}
	return collectionDtos
}
