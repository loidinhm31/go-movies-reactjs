package service

import (
	"context"
	"errors"
	"log"
	"movies-service/internal/dto"
	"movies-service/internal/movies"
)

type movieService struct {
	movieRepository movies.MovieRepository
}

func NewMovieService(movieRepository movies.MovieRepository) movies.Service {
	return &movieService{
		movieRepository: movieRepository,
	}
}

func (ms *movieService) GetAllMovies(ctx context.Context) ([]*dto.MovieDto, error) {
	movies, err := ms.movieRepository.FindAllMovies(ctx)
	if err != nil {
		log.Println(err)
		return nil, errors.New("not found")
	}

	var movieDtos []*dto.MovieDto
	for _, m := range movies {
		var genreDtos []*dto.GenreDto
		if m.Genres != nil {
			for _, g := range m.Genres {
				genreDtos = append(genreDtos, &dto.GenreDto{
					ID:      g.ID,
					Genre:   g.Genre,
					Checked: g.Checked,
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
	return movieDtos, nil
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
				ID:      g.ID,
				Genre:   g.Genre,
				Checked: g.Checked,
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

func (ms *movieService) GetMoviesByGenre(ctx context.Context, genreId int) ([]*dto.MovieDto, error) {
	movies, err := ms.movieRepository.FindMoviesByGenre(ctx, genreId)
	if err != nil {
		log.Println(err)
		return nil, errors.New("not found")
	}

	var movieDtos []*dto.MovieDto
	for _, m := range movies {
		var genreDtos []*dto.GenreDto
		if m.Genres != nil {
			for _, g := range m.Genres {
				genreDtos = append(genreDtos, &dto.GenreDto{
					ID:      g.ID,
					Genre:   g.Genre,
					Checked: g.Checked,
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
	return movieDtos, nil
}

func (ms *movieService) UpdateMovie(ctx context.Context, movie *dto.MovieDto) error {
	//TODO implement me
	panic("implement me")
}

func (ms *movieService) DeleteMovieById(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}
