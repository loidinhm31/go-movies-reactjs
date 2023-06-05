package service

import (
	"context"
	"fmt"
	"movies-service/internal/collection"
	"movies-service/internal/common/dto"
	"movies-service/internal/common/mapper"
	model2 "movies-service/internal/common/model"
	"movies-service/internal/episode"
	"movies-service/internal/errors"
	"movies-service/internal/middlewares"
	"movies-service/internal/movie"
	"movies-service/internal/payment"
	"movies-service/internal/user"
	"movies-service/pkg/pagination"
	"movies-service/pkg/util"
	"time"
)

type collectionService struct {
	userRepository       user.UserRepository
	movieRepository      movie.Repository
	episodeRepository    episode.Repository
	paymentRepository    payment.Repository
	collectionRepository collection.Repository
}

func NewCollectionService(userRepository user.UserRepository, movieRepository movie.Repository, episodeRepository episode.Repository, paymentRepository payment.Repository, collectionRepository collection.Repository) collection.Service {
	return &collectionService{
		userRepository:       userRepository,
		movieRepository:      movieRepository,
		episodeRepository:    episodeRepository,
		paymentRepository:    paymentRepository,
		collectionRepository: collectionRepository,
	}
}

func (fs *collectionService) AddCollection(ctx context.Context, collection *dto.CollectionDto) error {
	if collection.TypeCode == "" {
		return errors.ErrInvalidInput
	}

	// Get author
	author := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	theUser, _ := fs.userRepository.FindUserByUsername(ctx, author)
	if theUser.Role.RoleCode == "BANNED" {
		return errors.ErrInvalidClient
	}

	// Check free movie
	if collection.TypeCode == "MOVIE" {
		theMovie, err := fs.movieRepository.FindMovieByID(ctx, collection.MovieID)
		if err != nil {
			return err
		}

		if theMovie.Price.Valid {
			thePayment, err := fs.paymentRepository.FindPaymentByTypeCodeAndRefID(ctx, collection.TypeCode, theMovie.ID)
			if err != nil {
				return err
			}

			if !(thePayment.RefID == theMovie.ID && thePayment.UserID == theUser.ID) {
				return errors.ErrPaymentNotFound
			}
		}

		err = fs.collectionRepository.InsertCollection(ctx, &model2.Collection{
			UserID:    theUser.ID,
			MovieID:   util.IntToSQLNullInt(int64(theMovie.ID)),
			TypeCode:  collection.TypeCode,
			CreatedAt: time.Now(),
			CreatedBy: author,
		})
		if err != nil {
			return err
		}
	} else if collection.TypeCode == "TV" {
		theEpisode, err := fs.episodeRepository.FindEpisodeByID(ctx, collection.EpisodeID)
		if err != nil {
			return err
		}

		if theEpisode.Price.Valid {
			thePayment, err := fs.paymentRepository.FindPaymentByTypeCodeAndRefID(ctx, collection.TypeCode, theEpisode.ID)
			if err != nil {
				return err
			}

			if !(thePayment.RefID == theEpisode.ID && thePayment.UserID == theUser.ID) {
				return errors.ErrPaymentNotFound
			}
		}

		err = fs.collectionRepository.InsertCollection(ctx, &model2.Collection{
			UserID:    theUser.ID,
			EpisodeID: util.IntToSQLNullInt(int64(theEpisode.ID)),
			TypeCode:  collection.TypeCode,
			CreatedAt: time.Now(),
			CreatedBy: author,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (fs *collectionService) GetCollectionsByUserAndType(ctx context.Context, movieType string, keyword string, pageRequest *pagination.PageRequest) (*pagination.Page[*dto.CollectionDetailDto], error) {
	if movieType == "" {
		return nil, errors.ErrInvalidInput
	}

	// Get author
	author := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	theUser, err := fs.userRepository.FindUserByUsername(ctx, author)
	if err != nil {
		return nil, err
	}

	if theUser.Role.RoleCode == "BANNED" {
		return nil, errors.ErrInvalidClient
	}

	page := &pagination.Page[*model2.CollectionDetail]{}

	results, err := fs.collectionRepository.FindCollectionsByUserIDAndType(ctx, theUser.ID, movieType, keyword, pageRequest, page)
	if err != nil {
		return nil, err
	}

	collectionDtos := mapper.MapToCollectionDetailDtoSlice(results.Content)
	return &pagination.Page[*dto.CollectionDetailDto]{
		PageSize:      pageRequest.PageSize,
		PageNumber:    pageRequest.PageNumber,
		Sort:          pageRequest.Sort,
		TotalElements: results.TotalElements,
		TotalPages:    results.TotalPages,
		Content:       collectionDtos,
	}, nil
}

func (fs *collectionService) GetCollectionByUserAndRefID(ctx context.Context, typeCode string, refID uint) (*dto.CollectionDto, error) {
	if typeCode == "" {
		return nil, errors.ErrInvalidInput
	}

	username := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	theUser, err := fs.userRepository.FindUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	if theUser.Role.RoleCode == "BANNED" {
		return nil, errors.ErrInvalidClient
	}

	var result *model2.Collection
	if typeCode == "MOVIE" {
		result, err = fs.collectionRepository.FindCollectionByUserIDAndMovieID(ctx, theUser.ID, refID)
		if err != nil {
			return nil, err
		}
	} else if typeCode == "TV" {
		result, err = fs.collectionRepository.FindCollectionByUserIDAndEpisodeID(ctx, theUser.ID, refID)
		if err != nil {
			return nil, err
		}
	}

	return mapper.MapToCollectionDto(result), nil
}

func (fs *collectionService) RemoveCollectionByRefID(ctx context.Context, typeCode string, refID uint) error {
	if typeCode == "" {
		return errors.ErrInvalidInput
	}

	username := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	theUser, err := fs.userRepository.FindUserByUsername(ctx, username)
	if err != nil {
		return err
	}

	if theUser.Role.RoleCode == "BANNED" {
		return errors.ErrInvalidClient
	}

	if typeCode == "MOVIE" {
		err = fs.collectionRepository.DeleteCollectionByTypeCodeAndMovieID(ctx, typeCode, refID)
	} else if typeCode == "TV" {
		err = fs.collectionRepository.DeleteCollectionByTypeCodeAndEpisodeID(ctx, typeCode, refID)
	}
	return nil
}
