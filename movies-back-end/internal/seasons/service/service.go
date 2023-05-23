package service

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"movies-service/internal/control"
	"movies-service/internal/dto"
	"movies-service/internal/episodes"
	"movies-service/internal/errors"
	"movies-service/internal/mapper"
	"movies-service/internal/middlewares"
	"movies-service/internal/models"
	"movies-service/internal/movies"
	"movies-service/internal/seasons"
	"time"
)

type seasonService struct {
	mgmtCtrl          control.Service
	movieRepository   movies.MovieRepository
	seasonRepository  seasons.Repository
	episodeRepository episodes.Repository
}

func NewSeasonService(mgmtCtrl control.Service, movieRepository movies.MovieRepository, seasonRepository seasons.Repository, episodeRepository episodes.Repository) seasons.Service {
	return &seasonService{
		mgmtCtrl:          mgmtCtrl,
		movieRepository:   movieRepository,
		seasonRepository:  seasonRepository,
		episodeRepository: episodeRepository,
	}
}

func (s seasonService) GetSeasonsByID(ctx *gin.Context, id int) (*dto.SeasonDto, error) {
	result, err := s.seasonRepository.FindSeasonByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return mapper.MapToSeasonDto(result), nil
}

func (s seasonService) GetSeasonsByMovieID(ctx context.Context, movieID int) ([]*dto.SeasonDto, error) {
	result, err := s.seasonRepository.FindSeasonsByMovieID(ctx, movieID)
	if err != nil {
		return nil, err
	}
	return mapper.MapToSeasonDtoSlice(result), nil
}

func (s seasonService) AddSeason(ctx context.Context, season *dto.SeasonDto) error {
	if season.ID > 0 ||
		season.Name == "" ||
		season.Description == "" ||
		season.AirDate.IsZero() ||
		season.MovieID == 0 {
		return errors.ErrInvalidInput
	}

	// Get author
	author := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	if !s.mgmtCtrl.CheckPrivilege(author) {
		return errors.ErrUnAuthorized
	}

	movieObj, err := s.movieRepository.FindMovieById(ctx, season.MovieID)
	if err != nil {
		return err
	}

	seasonObject := &models.Season{
		Name:        season.Name,
		AirDate:     season.AirDate,
		Description: season.Description,
		Movie:       movieObj,
		CreatedAt:   time.Now(),
		CreatedBy:   author,
		UpdatedAt:   time.Now(),
		UpdatedBy:   author,
	}
	err = s.seasonRepository.InsertSeason(ctx, seasonObject)
	if err != nil {
		return err
	}
	return nil
}

func (s seasonService) UpdateSeason(ctx context.Context, season *dto.SeasonDto) error {
	if season.ID == 0 ||
		season.Name == "" ||
		season.Description == "" ||
		season.AirDate.IsZero() ||
		season.MovieID == 0 {
		return errors.ErrInvalidInput
	}

	// Check season exists
	seasonObj, err := s.seasonRepository.FindSeasonByID(ctx, season.ID)
	if err != nil {
		return errors.ErrResourceNotFound
	}

	// Get author
	author := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	if !s.mgmtCtrl.CheckPrivilege(author) {
		return errors.ErrUnAuthorized
	}

	movieObj, err := s.movieRepository.FindMovieById(ctx, season.MovieID)
	if err != nil {
		return err
	}

	// After check object exists, write updating value
	seasonObj = &models.Season{
		ID:          season.ID,
		Name:        season.Name,
		AirDate:     season.AirDate,
		Description: season.Description,
		Movie:       movieObj,
		CreatedAt:   time.Now(),
		CreatedBy:   author,
		UpdatedAt:   time.Now(),
		UpdatedBy:   author,
	}

	err = s.seasonRepository.UpdateSeason(ctx, seasonObj)
	if err != nil {
		return err
	}
	return nil
}

func (s seasonService) DeleteSeasonById(ctx context.Context, id int) error {
	if id == 0 {
		return errors.ErrInvalidInput
	}

	// Get author
	author := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	if !s.mgmtCtrl.CheckPrivilege(author) {
		return errors.ErrUnAuthorized
	}

	err := s.episodeRepository.DeleteEpisodeBySeasonID(ctx, id)
	if err != nil {
		return err
	}

	err = s.seasonRepository.DeleteSeasonByID(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
