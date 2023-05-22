package service

import (
	"context"
	"fmt"
	"log"
	"movies-service/internal/control"
	"movies-service/internal/dto"
	"movies-service/internal/errors"
	"movies-service/internal/mapper"
	"movies-service/internal/middlewares"
	"movies-service/internal/models"
	"movies-service/internal/movies"
	"movies-service/pkg/pagination"
)

type movieService struct {
	mgmtCtrl        control.Service
	movieRepository movies.MovieRepository
}

func NewMovieService(mgmtCtrl control.Service, movieRepository movies.MovieRepository) movies.Service {
	return &movieService{
		mgmtCtrl:        mgmtCtrl,
		movieRepository: movieRepository,
	}
}

func (ms *movieService) GetAllMoviesByType(ctx context.Context, movieType string, pageRequest *pagination.PageRequest) (*pagination.Page[*dto.MovieDto], error) {
	page := &pagination.Page[*models.Movie]{}

	var err error
	var movieResults *pagination.Page[*models.Movie]
	if movieType != "" {
		movieResults, err = ms.movieRepository.FindAllMoviesByType(ctx, movieType, pageRequest, page)
		if err != nil {
			log.Println(err)
			return nil, errors.ErrResourceNotFound
		}
	} else {
		movieResults, err = ms.movieRepository.FindAllMovies(ctx, pageRequest, page)
		if err != nil {
			log.Println(err)
			return nil, errors.ErrResourceNotFound
		}
	}

	movieDtos := mapper.MapToMovieDtoSlice(movieResults.Content)

	return &pagination.Page[*dto.MovieDto]{
		PageSize:      pageRequest.PageSize,
		PageNumber:    pageRequest.PageNumber,
		Sort:          pageRequest.Sort,
		TotalElements: movieResults.TotalElements,
		TotalPages:    movieResults.TotalPages,
		Content:       movieDtos,
	}, nil
}

func (ms *movieService) GetMovieById(ctx context.Context, id int) (*dto.MovieDto, error) {
	movie, err := ms.movieRepository.FindMovieById(ctx, id)
	if err != nil {
		log.Println(err)
		return nil, errors.ErrResourceNotFound
	}

	movieDto := mapper.MapToMovieDto(movie)
	return movieDto, nil
}

func (ms *movieService) GetMoviesByGenre(ctx context.Context, pageRequest *pagination.PageRequest, genreId int) (*pagination.Page[*dto.MovieDto], error) {
	page := &pagination.Page[*models.Movie]{}

	movieResults, err := ms.movieRepository.FindMoviesByGenre(ctx, pageRequest, page, genreId)
	if err != nil {
		log.Println(err)
		return nil, errors.ErrResourceNotFound
	}

	movieDtos := mapper.MapToMovieDtoSlice(movieResults.Content)

	return &pagination.Page[*dto.MovieDto]{
		PageSize:      pageRequest.PageSize,
		PageNumber:    pageRequest.PageNumber,
		Sort:          pageRequest.Sort,
		TotalElements: movieResults.TotalElements,
		TotalPages:    movieResults.TotalPages,
		Content:       movieDtos,
	}, nil
}

func (ms *movieService) AddMovie(ctx context.Context, movie *dto.MovieDto) error {
	// Get author
	author := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	if !ms.mgmtCtrl.CheckPrivilege(author) {
		return errors.ErrUnAuthorized
	}

	var genreObjects []*models.Genre

	if movie.ID > 0 ||
		movie.Title == "" ||
		movie.TypeCode == "" ||
		movie.Runtime == 0 ||
		movie.Description == "" ||
		movie.ReleaseDate.IsZero() ||
		movie.MpaaRating == "" ||
		(movie.Genres == nil || len(movie.Genres) == 0) {
		return errors.ErrInvalidInput
	}

	for _, genre := range movie.Genres {
		if genre.Checked {
			genreObjects = append(genreObjects, mapper.MapToGenre(genre, author))
		}
	}

	movieObject := mapper.MapToMovie(movie, author)
	movieObject.Genres = genreObjects

	err := ms.movieRepository.InsertMovie(ctx, movieObject)
	if err != nil {
		return err
	}
	return nil
}

func (ms *movieService) UpdateMovie(ctx context.Context, movie *dto.MovieDto) error {
	// Get author
	author := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	if !ms.mgmtCtrl.CheckPrivilege(author) {
		return errors.ErrUnAuthorized
	}

	if movie.ID == 0 ||
		movie.Title == "" ||
		movie.Runtime == 0 ||
		movie.Description == "" ||
		movie.ReleaseDate.IsZero() ||
		(movie.Genres == nil || len(movie.Genres) == 0) {
		return errors.ErrInvalidInput
	}

	// Check movie exists
	movieObj, err := ms.movieRepository.FindMovieById(ctx, movie.ID)
	if err != nil {
		return errors.ErrResourceNotFound
	}

	// After check object exists, write updating value
	movieObj = mapper.MapToMovieUpdate(movie, author)

	err = ms.movieRepository.UpdateMovie(ctx, movieObj)
	if err != nil {
		return err
	}

	var genreObjects []*models.Genre
	for _, genre := range movie.Genres {
		if genre.Checked {
			genreObjects = append(genreObjects, mapper.MapToGenre(genre, author))
		}
	}
	err = ms.movieRepository.UpdateMovieGenres(ctx, movieObj, genreObjects)
	if err != nil {
		return err
	}
	return nil
}

func (ms *movieService) DeleteMovieById(ctx context.Context, id int) error {
	// Get author
	author := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	if !ms.mgmtCtrl.CheckPrivilege(author) {
		return errors.ErrUnAuthorized
	}

	err := ms.movieRepository.DeleteMovieById(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
