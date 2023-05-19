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

func (gs *genreService) GetAllGenres(ctx context.Context) ([]*dto.GenreDto, error) {
	allGenres, err := gs.genreRepository.FindAllGenres(ctx)
	if err != nil {
		log.Println(err)
		return nil, errors.New("not found")
	}

	var genreDtos []*dto.GenreDto
	for _, g := range allGenres {
		genreDtos = append(genreDtos, &dto.GenreDto{
			ID:        g.ID,
			Name:      g.Name,
			CreatedAt: g.CreatedAt,
			UpdatedAt: g.UpdatedAt,
		})
	}
	return genreDtos, nil
}

func (gs *genreService) AddGenres(ctx *gin.Context, genreNames []string) error {
	// Get author
	author := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	if !gs.mgmtCtrl.CheckPrivilege(author) {
		return errors.New("unauthorized")
	}

	var newGenres []*models.Genre
	for _, n := range genreNames {
		foundedGenre, err := gs.genreRepository.FindGenreByName(ctx, &models.Genre{
			Name: fmt.Sprintf("%s", n),
		})
		if err != nil {
			return err
		}

		if foundedGenre.Name == n {
			return errors.New(fmt.Sprintf("cannot add %s", n))
		}
		newGenres = append(newGenres, &models.Genre{
			Name:      n,
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
