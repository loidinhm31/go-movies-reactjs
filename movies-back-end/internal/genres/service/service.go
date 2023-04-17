package service

import (
	"context"
	"errors"
	"log"
	"movies-service/internal/dto"
	"movies-service/internal/genres"
)

type genreService struct {
	genreRepository genres.GenreRepository
}

func NewGenreService(genreRepository genres.GenreRepository) genres.Service {
	return &genreService{
		genreRepository: genreRepository,
	}
}

func (gs *genreService) GetAllGenres(ctx context.Context) ([]*dto.GenreDto, error) {
	genres, err := gs.genreRepository.FindAllGenres(ctx)
	if err != nil {
		log.Println(err)
		return nil, errors.New("not found")
	}

	var genreDtos []*dto.GenreDto
	for _, g := range genres {
		genreDtos = append(genreDtos, &dto.GenreDto{
			ID:        g.ID,
			Genre:     g.Genre,
			Checked:   g.Checked,
			CreatedAt: g.CreatedAt,
			UpdatedAt: g.UpdatedAt,
		})
	}
	return genreDtos, nil
}
