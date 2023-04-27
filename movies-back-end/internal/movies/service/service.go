package service

import (
	"context"
	"errors"
	"log"
	"movies-service/internal/dto"
	"movies-service/internal/models"
	"movies-service/internal/movies"
	"movies-service/pkg/pagination"
	"time"
)

type movieService struct {
	movieRepository movies.MovieRepository
}

func NewMovieService(movieRepository movies.MovieRepository) movies.Service {
	return &movieService{
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
			Image:       m.Image,
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
		Image:       movie.Image,
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
			Image:       m.Image,
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
				ID:    genre.ID,
				Genre: genre.Genre,
			})
		}
	}

	movieObject := &models.Movie{
		Title:       movie.Title,
		ReleaseDate: movie.ReleaseDate,
		Runtime:     movie.Runtime,
		MpaaRating:  movie.MpaaRating,
		Description: movie.Description,
		Image:       movie.Image,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Genres:      genreObjects,
	}

	err := ms.movieRepository.InsertMovie(ctx, movieObject)
	if err != nil {
		return err
	}
	return nil
}

func (ms *movieService) UpdateMovie(ctx context.Context, movie *dto.MovieDto) error {
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
		Image:       movie.Image,
		UpdatedAt:   time.Now(),
	}

	err = ms.movieRepository.UpdateMovie(ctx, movieObj)
	if err != nil {
		return err
	}

	var genreObjects []*models.Genre
	for _, genre := range movie.Genres {
		if genre.Checked {
			genreObjects = append(genreObjects, &models.Genre{
				ID: genre.ID,
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
	err := ms.movieRepository.DeleteMovieById(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
