package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"movies-service/internal/control"
	"movies-service/internal/dto"
	"movies-service/internal/middlewares"
	"movies-service/internal/models"
	"movies-service/internal/movies"
	"movies-service/pkg/pagination"
	"movies-service/pkg/utils"
	"time"
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

func (ms *movieService) GetAllMovies(ctx context.Context, pageRequest *pagination.PageRequest) (*pagination.Page[*dto.MovieDto], error) {
	page := &pagination.Page[*models.Movie]{}

	movieResults, err := ms.movieRepository.FindAllMovies(ctx, pageRequest, page)
	if err != nil {
		log.Println(err)
		return nil, errors.New("not found")
	}

	var movieDtos []*dto.MovieDto
	for _, m := range movieResults.Data {
		var genreDtos []*dto.GenreDto
		if m.Genres != nil {
			for _, g := range m.Genres {
				genreDtos = append(genreDtos, &dto.GenreDto{
					ID:    g.ID,
					Genre: g.Genre,
				})
			}
		}

		movieDtos = append(movieDtos, &dto.MovieDto{
			ID:          m.ID,
			Title:       m.Title,
			ReleaseDate: m.ReleaseDate,
			Runtime:     m.Runtime,
			MpaaRating:  m.MpaaRating,
			Description: m.Description,
			ImagePath:   m.ImagePath.String,
			VideoPath:   m.VideoPath.String,
			CreatedAt:   m.CreatedAt,
			UpdatedAt:   m.UpdatedAt,
			Genres:      genreDtos,
		})
	}
	return &pagination.Page[*dto.MovieDto]{
		PageSize:      pageRequest.PageSize,
		PageNumber:    pageRequest.PageNumber,
		Sort:          pageRequest.Sort,
		TotalElements: movieResults.TotalElements,
		TotalPages:    movieResults.TotalPages,
		Data:          movieDtos,
	}, nil
}

func (ms *movieService) GetMovieById(ctx context.Context, id int) (*dto.MovieDto, error) {
	movie, err := ms.movieRepository.FindMovieById(ctx, id)
	if err != nil {
		log.Println(err)
		return nil, errors.New("not found")
	}

	var genreDtos []*dto.GenreDto
	if movie.Genres != nil {
		for _, g := range movie.Genres {
			genreDtos = append(genreDtos, &dto.GenreDto{
				ID:    g.ID,
				Genre: g.Genre,
			})
		}
	}

	movieDto := &dto.MovieDto{
		ID:          movie.ID,
		Title:       movie.Title,
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
	return movieDto, nil
}

func (ms *movieService) GetMoviesByGenre(ctx context.Context, pageRequest *pagination.PageRequest, genreId int) (*pagination.Page[*dto.MovieDto], error) {
	page := &pagination.Page[*models.Movie]{}

	movieResults, err := ms.movieRepository.FindMoviesByGenre(ctx, pageRequest, page, genreId)
	if err != nil {
		log.Println(err)
		return nil, errors.New("not found")
	}

	var movieDtos []*dto.MovieDto
	for _, m := range movieResults.Data {
		var genreDtos []*dto.GenreDto
		if m.Genres != nil {
			for _, g := range m.Genres {
				genreDtos = append(genreDtos, &dto.GenreDto{
					ID:    g.ID,
					Genre: g.Genre,
				})
			}
		}

		movieDtos = append(movieDtos, &dto.MovieDto{
			ID:          m.ID,
			Title:       m.Title,
			ReleaseDate: m.ReleaseDate,
			Runtime:     m.Runtime,
			MpaaRating:  m.MpaaRating,
			Description: m.Description,
			ImagePath:   m.ImagePath.String,
			VideoPath:   m.VideoPath.String,
			CreatedAt:   m.CreatedAt,
			UpdatedAt:   m.UpdatedAt,
			Genres:      genreDtos,
		})
	}
	return &pagination.Page[*dto.MovieDto]{
		PageSize:      pageRequest.PageSize,
		PageNumber:    pageRequest.PageNumber,
		Sort:          pageRequest.Sort,
		TotalElements: movieResults.TotalElements,
		TotalPages:    movieResults.TotalPages,
		Data:          movieDtos,
	}, nil
}

func (ms *movieService) AddMovie(ctx context.Context, movie *dto.MovieDto) error {
	// Get author
	author := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	if !ms.mgmtCtrl.CheckPrivilege(author) {
		return errors.New("unauthorized")
	}

	var genreObjects []*models.Genre

	if movie.ID > 0 ||
		movie.Title == "" ||
		movie.Runtime == 0 ||
		movie.Description == "" ||
		movie.ReleaseDate.IsZero() ||
		(movie.Genres == nil || len(movie.Genres) == 0) {
		return errors.New("invalid input")
	}

	for _, genre := range movie.Genres {
		if genre.Checked {
			genreObjects = append(genreObjects, &models.Genre{
				ID:        genre.ID,
				Genre:     genre.Genre,
				CreatedBy: author,
				UpdatedBy: author,
			})
		}
	}

	movieObject := &models.Movie{
		Title:       movie.Title,
		ReleaseDate: movie.ReleaseDate,
		Runtime:     movie.Runtime,
		MpaaRating:  movie.MpaaRating,
		Description: movie.Description,
		ImagePath:   utils.StringToSQLNullString(movie.ImagePath),
		VideoPath:   utils.StringToSQLNullString(movie.VideoPath),
		CreatedAt:   time.Now(),
		CreatedBy:   author,
		UpdatedAt:   time.Now(),
		UpdatedBy:   author,
		Genres:      genreObjects,
	}

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
		return errors.New("unauthorized")
	}

	if movie.ID == 0 ||
		movie.Title == "" ||
		movie.Runtime == 0 ||
		movie.Description == "" ||
		movie.ReleaseDate.IsZero() ||
		(movie.Genres == nil || len(movie.Genres) == 0) {
		return errors.New("invalid input")
	}

	movieObj, err := ms.movieRepository.FindMovieById(ctx, movie.ID)
	if err != nil {
		return errors.New("cannot find object")
	}

	movieObj = &models.Movie{
		ID:          movie.ID,
		Title:       movie.Title,
		ReleaseDate: movie.ReleaseDate,
		Runtime:     movie.Runtime,
		MpaaRating:  movie.MpaaRating,
		Description: movie.Description,
		ImagePath:   utils.StringToSQLNullString(movie.ImagePath),
		VideoPath:   utils.StringToSQLNullString(movie.VideoPath),
		UpdatedAt:   time.Now(),
		UpdatedBy:   author,
	}

	err = ms.movieRepository.UpdateMovie(ctx, movieObj)
	if err != nil {
		return err
	}

	var genreObjects []*models.Genre
	for _, genre := range movie.Genres {
		if genre.Checked {
			genreObjects = append(genreObjects, &models.Genre{
				ID:        genre.ID,
				CreatedBy: author,
				UpdatedBy: author,
			})
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
		return errors.New("unauthorized")
	}

	err := ms.movieRepository.DeleteMovieById(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
