package service

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"movies-service/internal/blob"
	"movies-service/internal/collection"
	"movies-service/internal/common/dto"
	"movies-service/internal/common/mapper"
	"movies-service/internal/common/model"
	"movies-service/internal/control"
	"movies-service/internal/episode"
	"movies-service/internal/errors"
	"movies-service/internal/middlewares"
	"movies-service/internal/payment"
	"movies-service/internal/season"
	"movies-service/pkg/util"
	"strings"
	"sync"
	"time"
)

type episodeService struct {
	mgmtCtrl             control.Service
	seasonRepository     season.Repository
	collectionRepository collection.Repository
	paymentRepository    payment.Repository
	episodeRepository    episode.Repository
	blobService          blob.Service
}

func NewEpisodeService(mgmtCtrl control.Service, seasonRepository season.Repository,
	collectionRepository collection.Repository, paymentRepository payment.Repository,
	episodeRepository episode.Repository, blobService blob.Service) episode.Service {
	return &episodeService{
		mgmtCtrl:             mgmtCtrl,
		seasonRepository:     seasonRepository,
		collectionRepository: collectionRepository,
		paymentRepository:    paymentRepository,
		episodeRepository:    episodeRepository,
		blobService:          blobService,
	}
}

func (es episodeService) GetEpisodesByID(ctx context.Context, id uint) (*dto.EpisodeDto, error) {
	author := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	isValidUser, isPrivilege := es.mgmtCtrl.CheckUser(author)

	result, err := es.episodeRepository.FindEpisodeByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Check valid video path
	if result.Price.Valid {
		thePayment, err := es.paymentRepository.FindPaymentByTypeCodeAndRefID(ctx, "TV", result.ID)
		if err != nil {
			return nil, err
		}

		if !(thePayment.TypeCode == "TV" && thePayment.RefID == result.ID) {
			theCollection, err := es.collectionRepository.FindCollectionByEpisodeID(ctx, result.ID)
			if err != nil {
				return nil, err
			}

			if !(theCollection.TypeCode == "TV" && uint(theCollection.MovieID.Int64) == result.ID) {
				result.VideoPath = sql.NullString{}
			}
		}
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
		Price:     util.FloatToSQLNullFloat(episode.Price),
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
		Price:     util.FloatToSQLNullFloat(episode.Price),
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

	// Get current episode
	episodeObj, err := es.episodeRepository.FindEpisodeByID(ctx, id)
	if err != nil {
		return err
	}

	// Delete video from blob concurrently
	if episodeObj.VideoPath.Valid {
		var wg sync.WaitGroup
		wg.Add(1)

		go func() {
			defer wg.Done()
			videoPath := episodeObj.VideoPath.String
			videoPathSplit := strings.Split(videoPath, "/")
			videoKey := videoPathSplit[len(videoPathSplit)-1]
			res, err := es.blobService.DeleteFile(ctx, videoKey, "video")
			if err != nil {
				log.Println("cannot delete video")
			}
			log.Println(res)
		}()
		wg.Wait()
	}

	err = es.episodeRepository.DeleteEpisodeByID(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
