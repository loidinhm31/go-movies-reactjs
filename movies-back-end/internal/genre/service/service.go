package service

import (
	"context"
	"fmt"
	"log"
	"movies-service/internal/common/dto"
	"movies-service/internal/common/entity"
	"movies-service/internal/common/mapper"
	"movies-service/internal/control"
	"movies-service/internal/errors"
	"movies-service/internal/genre"
	"movies-service/internal/middlewares"
	"time"
)

type genreService struct {
	mgmtCtrl        control.Service
	genreRepository genre.Repository
}

func NewGenreService(mgmtCtrl control.Service, genreRepository genre.Repository) genre.Service {
	return &genreService{
		mgmtCtrl:        mgmtCtrl,
		genreRepository: genreRepository,
	}
}

func (gs *genreService) GetAllGenresByTypeCode(ctx context.Context, movieType string) ([]*dto.GenreDto, error) {
	var err error
	var allGenres []*entity.Genre

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

	var newGenres []*entity.Genre
	for _, g := range genreDtos {
		foundedGenre, err := gs.genreRepository.FindGenreByNameAndTypeCode(ctx, &entity.Genre{
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

		newGenres = append(newGenres, &entity.Genre{
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
