package service

import (
	"context"
	"fmt"
	"log"
	"movies-service/internal/control"
	"movies-service/internal/dto"
	"movies-service/internal/errors"
	"movies-service/internal/genres"
	"movies-service/internal/mapper"
	"movies-service/internal/middlewares"
	"movies-service/internal/models"
	"time"
)

type genreService struct {
	mgmtCtrl        control.Service
	genreRepository genres.GenreRepository
}

func NewGenreService(mgmtCtrl control.Service, genreRepository genres.GenreRepository) genres.Service {
	return &genreService{
		mgmtCtrl:        mgmtCtrl,
		genreRepository: genreRepository,
	}
}

func (gs *genreService) GetAllGenresByTypeCode(ctx context.Context, movieType string) ([]*dto.GenreDto, error) {
	var err error
	var allGenres []*models.Genre

	if movieType != "" {
		allGenres, err = gs.genreRepository.FindAllGenresByTypeCode(ctx, movieType)
		if err != nil {
			log.Println(err)
			return nil, errors.ErrResourceNotFound
		}
	} else {
		allGenres, err = gs.genreRepository.FindAllGenres(ctx)
		if err != nil {
			log.Println(err)
			return nil, errors.ErrResourceNotFound
		}
	}

	genreDtos := mapper.MapToGenreDtoSlice(allGenres)
	return genreDtos, nil
}

func (gs *genreService) AddGenres(ctx context.Context, genreDtos []dto.GenreDto) error {
	// Get author
	author := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	if !gs.mgmtCtrl.CheckPrivilege(author) {
		return errors.ErrUnAuthorized
	}

	var newGenres []*models.Genre
	for _, g := range genreDtos {
		foundedGenre, err := gs.genreRepository.FindGenreByNameAndTypeCode(ctx, &models.Genre{
			Name:     g.Name,
			TypeCode: g.TypeCode,
		})
		if err != nil {
			return err
		}

		if !checkGenreType(g.TypeCode) {
			return errors.ErrInvalidInput
		}

		if foundedGenre.Name == g.Name && foundedGenre.TypeCode == foundedGenre.TypeCode {
			return errors.ErrCannotExecuteAction
		}

		newGenres = append(newGenres, &models.Genre{
			Name:      g.Name,
			TypeCode:  g.TypeCode,
			CreatedAt: time.Now(),
			CreatedBy: author,
			UpdatedAt: time.Now(),
			UpdatedBy: author,
		})
	}

	err := gs.genreRepository.InsertGenres(ctx, newGenres)
	if err != nil {
		return err
	}
	return nil
}

func checkGenreType(typeCode string) bool {
	if typeCode == "TV" || typeCode == "MOVIE" {
		return true
	}
	return false
}
