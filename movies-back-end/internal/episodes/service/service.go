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
	"movies-service/internal/seasons"
	"time"
)

type episodeService struct {
	mgmtCtrl          control.Service
	seasonRepository  seasons.Repository
	episodeRepository episodes.Repository
}

func NewSeasonService(mgmtCtrl control.Service, seasonRepository seasons.Repository, episodeRepository episodes.Repository) episodes.Service {
	return &episodeService{
		mgmtCtrl:          mgmtCtrl,
		seasonRepository:  seasonRepository,
		episodeRepository: episodeRepository,
	}
}

func (e episodeService) GetEpisodesByID(ctx *gin.Context, id int) (*dto.EpisodeDto, error) {
	result, err := e.episodeRepository.FindEpisodeByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return mapper.MapToEpisodeDto(result), nil
}

func (e episodeService) GetEpisodesBySeasonID(ctx context.Context, seasonID int) ([]*dto.EpisodeDto, error) {
	result, err := e.episodeRepository.FindEpisodesBySeasonID(ctx, seasonID)
	if err != nil {
		return nil, err
	}
	return mapper.MapToEpisodeDtoSlice(result), nil
}

func (e episodeService) AddEpisode(ctx context.Context, episode *dto.EpisodeDto) error {
	// Get author
	author := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	if !e.mgmtCtrl.CheckPrivilege(author) {
		return errors.ErrUnAuthorized
	}

	if episode.ID > 0 ||
		episode.Name == "" ||
		episode.AirDate.IsZero() ||
		episode.Runtime == 0 ||
		episode.SeasonID == 0 {
		return errors.ErrInvalidInput
	}

	seasonObj, err := e.seasonRepository.FindSeasonByID(ctx, episode.SeasonID)
	if err != nil {
		return err
	}

	episodeObject := &models.Episode{
		Name:      episode.Name,
		AirDate:   episode.AirDate,
		Runtime:   episode.Runtime,
		VideoPath: episode.VideoPath,
		Season:    seasonObj,
		CreatedAt: time.Now(),
		CreatedBy: author,
		UpdatedAt: time.Now(),
		UpdatedBy: author,
	}
	err = e.episodeRepository.InsertEpisode(ctx, episodeObject)
	if err != nil {
		return err
	}
	return nil
}

func (e episodeService) UpdateEpisode(ctx context.Context, episode *dto.EpisodeDto) error {
	// Get author
	author := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	if !e.mgmtCtrl.CheckPrivilege(author) {
		return errors.ErrUnAuthorized
	}

	if episode.ID == 0 ||
		episode.Name == "" ||
		episode.AirDate.IsZero() ||
		episode.Runtime == 0 ||
		episode.SeasonID == 0 {
		return errors.ErrInvalidInput
	}

	// Check episode exists
	episodeObj, err := e.episodeRepository.FindEpisodeByID(ctx, episode.ID)
	if err != nil {
		return errors.ErrResourceNotFound
	}

	seasonObj, err := e.seasonRepository.FindSeasonByID(ctx, episode.SeasonID)
	if err != nil {
		return err
	}

	// After check object exists, write updating value
	episodeObj = &models.Episode{
		ID:        episode.ID,
		Name:      episode.Name,
		AirDate:   episode.AirDate,
		Runtime:   episode.Runtime,
		VideoPath: episode.VideoPath,
		Season:    seasonObj,
		CreatedAt: time.Now(),
		CreatedBy: author,
		UpdatedAt: time.Now(),
		UpdatedBy: author,
	}

	err = e.episodeRepository.UpdateEpisode(ctx, episodeObj)
	if err != nil {
		return err
	}
	return nil
}

func (e episodeService) DeleteEpisodeById(ctx context.Context, id int) error {
	if id == 0 {
		return errors.ErrInvalidInput
	}

	author := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	if !e.mgmtCtrl.CheckPrivilege(author) {
		return errors.ErrUnAuthorized
	}

	err := e.episodeRepository.DeleteEpisodeById(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
