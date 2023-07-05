package service

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"movies-service/internal/common/dto"
	"movies-service/internal/common/entity"
	"movies-service/internal/common/mapper"
	"movies-service/internal/control"
	"movies-service/internal/episode"
	"movies-service/internal/errors"
	"movies-service/internal/middlewares"
	"movies-service/internal/movie"
	"movies-service/internal/season"
	"time"
)

type seasonService struct {
	mgmtCtrl          control.Service
	movieRepository   movie.Repository
	seasonRepository  season.Repository
	episodeRepository episode.Repository
}

func NewSeasonService(mgmtCtrl control.Service, movieRepository movie.Repository, seasonRepository season.Repository, episodeRepository episode.Repository) season.Service {
	return &seasonService{
		mgmtCtrl:          mgmtCtrl,
		movieRepository:   movieRepository,
		seasonRepository:  seasonRepository,
		episodeRepository: episodeRepository,
	}
}

func (s seasonService) GetSeasonsByID(ctx *gin.Context, id uint) (*dto.SeasonDto, error) {
	result, err := s.seasonRepository.FindSeasonByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return mapper.MapToSeasonDto(result), nil
}

func (s seasonService) GetSeasonsByMovieID(ctx context.Context, movieID uint) ([]*dto.SeasonDto, error) {
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

	movieObj, err := s.movieRepository.FindMovieByID(ctx, season.MovieID)
	if err != nil {
		return err
	}

	seasonObject := &entity.Season{
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

	movieObj, err := s.movieRepository.FindMovieByID(ctx, season.MovieID)
	if err != nil {
		return err
	}

	// After check object exists, write updating value
	seasonObj = &entity.Season{
		ID:          season.ID,
		Name:        season.Name,
		AirDate:     season.AirDate,
		Description: season.Description,
		Movie:       movieObj,
		CreatedAt:   seasonObj.CreatedAt,
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

func (s seasonService) RemoveSeasonByID(ctx context.Context, id uint) error {
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
