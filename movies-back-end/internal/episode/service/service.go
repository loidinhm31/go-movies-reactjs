package service

import (
	"context"
	"fmt"
	"log"
	"movies-service/internal/blob"
	"movies-service/internal/collection"
	"movies-service/internal/common/dto"
	"movies-service/internal/common/entity"
	"movies-service/internal/common/mapper"
	"movies-service/internal/control"
	"movies-service/internal/episode"
	"movies-service/internal/errors"
	"movies-service/internal/middlewares"
	"movies-service/internal/payment"
	"movies-service/internal/season"
	"movies-service/internal/user"
	"movies-service/pkg/util"
	"strings"
	"sync"
	"time"
)

type episodeService struct {
	mgmtCtrl             control.Service
	userRepository       user.UserRepository
	seasonRepository     season.Repository
	collectionRepository collection.Repository
	paymentRepository    payment.Repository
	episodeRepository    episode.Repository
	blobService          blob.Service
}

func NewEpisodeService(mgmtCtrl control.Service, userRepository user.UserRepository,
	seasonRepository season.Repository, collectionRepository collection.Repository,
	paymentRepository payment.Repository, episodeRepository episode.Repository,
	blobService blob.Service) episode.Service {
	return &episodeService{
		mgmtCtrl:             mgmtCtrl,
		userRepository:       userRepository,
		seasonRepository:     seasonRepository,
		collectionRepository: collectionRepository,
		paymentRepository:    paymentRepository,
		episodeRepository:    episodeRepository,
		blobService:          blobService,
	}
}

func (es episodeService) GetEpisodeByID(ctx context.Context, id uint) (*dto.EpisodeDto, error) {
	var author string
	if val := ctx.Value(middlewares.CtxUserKey); val != nil {
		author = fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	}

	result, err := es.episodeRepository.FindEpisodeByID(ctx, id)
	if err != nil {
		return nil, err
	}

	var episodeDto *dto.EpisodeDto
	if author != "" {
		theUser, err := es.userRepository.FindUserByUsernameAndIsNew(ctx, author, false)
		if err != nil {
			return nil, err
		}

		if theUser.Role.RoleCode == "BANNED" {
			return nil, errors.ErrInvalidClient
		}

		// Set valid video path
		isRestrict := false
		if result.Price.Valid {
			thePayment, err := es.paymentRepository.FindPaymentByUserIDAndTypeCodeAndRefID(ctx, theUser.ID, "TV", result.ID)
			if err != nil {
				return nil, err
			}

			// Not paid
			if !(thePayment.TypeCode == "TV" && thePayment.RefID == result.ID) {
				if !(theUser.Role.RoleCode == "ADMIN" || theUser.Role.RoleCode == "MOD") {
					isRestrict = true
				}
			}
		}

		episodeDto = mapper.MapToEpisodeDto(
			result,
			isRestrict,
			theUser.Role.RoleCode == "ADMIN" || theUser.Role.RoleCode == "MOD",
		)
	} else {
		episodeDto = mapper.MapToEpisodeDto(
			result,
			true,
			false,
		)
	}

	return episodeDto, nil
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

	episodeObject := &entity.Episode{
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
	episodeObj = &entity.Episode{
		ID:        episode.ID,
		Name:      episode.Name,
		AirDate:   episode.AirDate,
		Runtime:   episode.Runtime,
		VideoPath: util.StringToSQLNullString(episode.VideoPath),
		SeasonID:  seasonObj.ID,
		Price:     util.FloatToSQLNullFloat(episode.Price),
		CreatedAt: episodeObj.CreatedAt,
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

	// Check payment
	log.Println("checking payment before removing...")
	payments, err := es.paymentRepository.FindPaymentsByTypeCodeAndRefID(ctx, "TV", id)
	if err != nil {
		return err
	}

	if len(payments) > 0 {
		return errors.ErrCannotExecuteAction
	}

	// Check collection
	log.Println("checking collection before removing...")
	collections, err := es.collectionRepository.FindCollectionsByEpisodeID(ctx, id)
	if err != nil {
		return err
	}

	if len(collections) > 0 {
		return errors.ErrCannotExecuteAction
	}

	// Get author
	author := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	if !es.mgmtCtrl.CheckPrivilege(author) {
		return errors.ErrInvalidClient
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
