package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"movies-service/internal/control"
	"movies-service/internal/dto"
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
			return nil, errors.New("not found")
		}
	} else {
		allGenres, err = gs.genreRepository.FindAllGenres(ctx)
		if err != nil {
			log.Println(err)
			return nil, errors.New("not found")
		}
	}

	genreDtos := mapper.MaptoGenreDtoSlice(allGenres)
	return genreDtos, nil
}

func (gs *genreService) AddGenres(ctx *gin.Context, genreDtos []dto.Genres) error {
	// Get author
	author := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	if !gs.mgmtCtrl.CheckPrivilege(author) {
		return errors.New("unauthorized")
	}

	var newGenres []*models.Genre
	for _, g := range genreDtos {
		foundedGenre, err := gs.genreRepository.FindGenreByNameAndTypeCode(ctx, &models.Genre{
			Name: fmt.Sprintf("%s", g),
		})
		if err != nil {
			return err
		}

		if !checkGenreType(g.TypeCode) {
			return errors.New("invalid type code")
		}

		if foundedGenre.Name == g.Name {
			return errors.New(fmt.Sprintf("cannot add %s", g))
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
