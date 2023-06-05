package mapper

import (
	"database/sql"
	dto2 "movies-service/internal/common/dto"
	model2 "movies-service/internal/common/model"
	"time"
)

func MapToMovieDto(movie *model2.Movie, isRestrictResource bool, isPrivilege bool) *dto2.MovieDto {
	genreDtos := MapToGenreDtoSlice(movie.Genres)

	// Filter video path for release date
	if isRestrictResource ||
		(!isPrivilege && time.Now().Before(movie.ReleaseDate)) {
		movie.VideoPath = sql.NullString{}
	}

	return &dto2.MovieDto{
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

func MapToMovieDtoSlice(movieSlice []*model2.Movie) []*dto2.MovieDto {
	var movieDtos []*dto2.MovieDto
	for _, m := range movieSlice {
		movieDtos = append(movieDtos, MapToMovieDto(m, true, false))
	}
	return movieDtos
}

func MapToGenreDto(genre *model2.Genre) *dto2.GenreDto {
	return &dto2.GenreDto{
		ID:       genre.ID,
		Name:     genre.Name,
		TypeCode: genre.TypeCode,
	}
}

func MapToGenreDtoSlice(genreSlice []*model2.Genre) []*dto2.GenreDto {
	var genreDtos []*dto2.GenreDto
	for _, g := range genreSlice {
		genreDtos = append(genreDtos, MapToGenreDto(g))
	}
	return genreDtos
}

func MapToSeasonDto(season *model2.Season) *dto2.SeasonDto {
	return &dto2.SeasonDto{
		ID:          season.ID,
		Name:        season.Name,
		AirDate:     season.AirDate,
		Description: season.Description,
		MovieID:     season.MovieID,
		EpisodeDtos: nil,
	}
}

func MapToSeasonDtoSlice(seasonSlice []*model2.Season) []*dto2.SeasonDto {
	var seasonDtos []*dto2.SeasonDto
	for _, s := range seasonSlice {
		seasonDtos = append(seasonDtos, MapToSeasonDto(s))
	}
	return seasonDtos
}

func MapToEpisodeDto(episode *model2.Episode, isRestrictResource, isPrivilege bool) *dto2.EpisodeDto {
	// Filter video path for release date
	if isRestrictResource ||
		(!isPrivilege && time.Now().Before(episode.AirDate)) {
		episode.VideoPath = sql.NullString{}
	}

	var season *dto2.SeasonDto
	if episode.Season != nil {
		season = MapToSeasonDto(episode.Season)
	}

	return &dto2.EpisodeDto{
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

func MapToEpisodeDtoSlice(seasonSlice []*model2.Episode, isRestrictResource, isPrivilege bool) []*dto2.EpisodeDto {
	var episodeDtos []*dto2.EpisodeDto
	for _, e := range seasonSlice {
		episodeDtos = append(episodeDtos, MapToEpisodeDto(e, isRestrictResource, isPrivilege))
	}
	return episodeDtos
}

func MapToRatingDto(rating *model2.Rating) *dto2.RatingDto {
	return &dto2.RatingDto{
		ID:   rating.ID,
		Code: rating.Code,
		Name: rating.Name,
	}
}

func MapToRatingDtoSlice(ratingSlice []*model2.Rating) []*dto2.RatingDto {
	var ratingDtos []*dto2.RatingDto
	for _, r := range ratingSlice {
		ratingDtos = append(ratingDtos, MapToRatingDto(r))
	}
	return ratingDtos
}

func MapToCollectionDto(collection *model2.Collection) *dto2.CollectionDto {
	return &dto2.CollectionDto{
		UserID:    collection.UserID,
		MovieID:   uint(collection.MovieID.Int64),
		EpisodeID: uint(collection.EpisodeID.Int64),
	}
}

func MapToCollectionDetailDto(collection *model2.CollectionDetail) *dto2.CollectionDetailDto {
	return &dto2.CollectionDetailDto{
		UserID:      collection.Username,
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

func MapToCollectionDetailDtoSlice(collections []*model2.CollectionDetail) []*dto2.CollectionDetailDto {
	var collectionDtos []*dto2.CollectionDetailDto
	for _, c := range collections {
		collectionDtos = append(collectionDtos, MapToCollectionDetailDto(c))
	}
	return collectionDtos
}

func MapToPaymentDto(payment *model2.Payment) *dto2.PaymentDto {
	return &dto2.PaymentDto{
		ID:            payment.ID,
		UserID:        payment.UserID,
		RefID:         payment.RefID,
		TypeCode:      payment.TypeCode,
		Provider:      payment.Provider,
		Amount:        payment.Amount,
		Currency:      payment.Currency,
		PaymentMethod: payment.PaymentMethod,
		Status:        payment.Status,
		CreatedAt:     payment.CreatedAt,
	}
}

func MapToPaymentDtoSlice(payments []*model2.Payment) []*dto2.PaymentDto {
	var paymentDtos []*dto2.PaymentDto
	for _, p := range payments {
		paymentDtos = append(paymentDtos, MapToPaymentDto(p))
	}
	return paymentDtos
}
