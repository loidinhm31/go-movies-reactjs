package service

import (
	"context"
	"fmt"
	"movies-service/internal/control"
	"movies-service/internal/dto"
	"movies-service/internal/episode"
	"movies-service/internal/errors"
	"movies-service/internal/mapper"
	"movies-service/internal/middlewares"
	"movies-service/internal/model"
	"movies-service/internal/season"
	"movies-service/pkg/util"
	"time"
)

type episodeService struct {
	mgmtCtrl          control.Service
	seasonRepository  season.Repository
	episodeRepository episode.Repository
}

func NewEpisodeService(mgmtCtrl control.Service, seasonRepository season.Repository, episodeRepository episode.Repository) episode.Service {
	return &episodeService{
		mgmtCtrl:          mgmtCtrl,
		seasonRepository:  seasonRepository,
		episodeRepository: episodeRepository,
	}
}

func (es episodeService) GetEpisodesByID(ctx context.Context, id uint) (*dto.EpisodeDto, error) {
	author := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	isValidUser, isPrivilege := es.mgmtCtrl.CheckUser(author)

	result, err := es.episodeRepository.FindEpisodeByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return mapper.MapToEpisodeDto(result, !isValidUser, isPrivilege), nil
}

func (es episodeService) GetEpisodesBySeasonID(ctx context.Context, seasonID uint) ([]*dto.EpisodeDto, error) {
	author := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	isValidUser, isPrivilege := es.mgmtCtrl.CheckUser(author)

	result, err := es.episodeRepository.FindEpisodesBySeasonID(ctx, seasonID)
	if err != nil {
		return nil, err
	}
	return mapper.MapToEpisodeDtoSlice(result, !isValidUser, isPrivilege), nil
}

func (es episodeService) AddEpisode(ctx context.Context, episode *dto.EpisodeDto) error {
	if episode.ID > 0 ||
		episode.Name == "" ||
		episode.AirDate.IsZero() ||
		episode.Runtime == 0 ||
		episode.SeasonID == 0 {
		return errors.ErrInvalidInput
	}

	// Get author
	author := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	if !es.mgmtCtrl.CheckPrivilege(author) {
		return errors.ErrUnAuthorized
	}

	seasonObj, err := es.seasonRepository.FindSeasonByID(ctx, episode.SeasonID)
	if err != nil {
		return err
	}

	episodeObject := &model.Episode{
		Name:      episode.Name,
		AirDate:   episode.AirDate,
		Runtime:   episode.Runtime,
		VideoPath: util.StringToSQLNullString(episode.VideoPath),
		Season:    seasonObj,
		CreatedAt: time.Now(),
		CreatedBy: author,
		UpdatedAt: time.Now(),
		UpdatedBy: author,
	}
	err = es.episodeRepository.InsertEpisode(ctx, episodeObject)
	if err != nil {
		return err
	}
	return nil
}

func (es episodeService) UpdateEpisode(ctx context.Context, episode *dto.EpisodeDto) error {
	if episode.ID == 0 ||
		episode.Name == "" ||
		episode.AirDate.IsZero() ||
		episode.Runtime == 0 ||
		episode.SeasonID == 0 {
		return errors.ErrInvalidInput
	}

	// Check episode exists
	episodeObj, err := es.episodeRepository.FindEpisodeByID(ctx, episode.ID)
	if err != nil {
		return errors.ErrResourceNotFound
	}

	// Get author
	author := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	if !es.mgmtCtrl.CheckPrivilege(author) {
		return errors.ErrUnAuthorized
	}

	seasonObj, err := es.seasonRepository.FindSeasonByID(ctx, episode.SeasonID)
	if err != nil {
		return err
	}

	// After check object exists, write updating value
	episodeObj = &model.Episode{
		ID:        episode.ID,
		Name:      episode.Name,
		AirDate:   episode.AirDate,
		Runtime:   episode.Runtime,
		VideoPath: util.StringToSQLNullString(episode.VideoPath),
		Season:    seasonObj,
		CreatedAt: time.Now(),
		CreatedBy: author,
		UpdatedAt: time.Now(),
		UpdatedBy: author,
	}

	err = es.episodeRepository.UpdateEpisode(ctx, episodeObj)
	if err != nil {
		return err
	}
	return nil
}

func (es episodeService) RemoveEpisodeByID(ctx context.Context, id uint) error {
	if id == 0 {
		return errors.ErrInvalidInput
	}

	// Get author
	author := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	if !es.mgmtCtrl.CheckPrivilege(author) {
		return errors.ErrUnAuthorized
	}

	err := es.episodeRepository.DeleteEpisodeByID(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
