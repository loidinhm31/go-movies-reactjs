package service

import (
	"context"
	"fmt"
	"movies-service/internal/collection"
	"movies-service/internal/control"
	"movies-service/internal/dto"
	"movies-service/internal/episode"
	"movies-service/internal/errors"
	"movies-service/internal/mapper"
	"movies-service/internal/middlewares"
	"movies-service/internal/model"
	"movies-service/internal/movie"
	"movies-service/pkg/pagination"
	"movies-service/pkg/util"
	"time"
)

type collectionService struct {
	mgmtCtrl             control.Service
	movieRepository      movie.Repository
	episodeRepository    episode.Repository
	collectionRepository collection.Repository
}

func NewCollectionService(mgmtCtrl control.Service, movieRepository movie.Repository, episodeRepository episode.Repository, collectionRepository collection.Repository) collection.Service {
	return &collectionService{
		mgmtCtrl:             mgmtCtrl,
		movieRepository:      movieRepository,
		episodeRepository:    episodeRepository,
		collectionRepository: collectionRepository,
	}
}

func (fs collectionService) AddCollection(ctx context.Context, collection *dto.CollectionDto) error {
	if collection.TypeCode == "" {
		return errors.ErrInvalidInput
	}

	// Get author
	author := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	isValidUser, _ := fs.mgmtCtrl.CheckUser(author)
	if !isValidUser {
		return errors.ErrUnAuthorized
	}

	// Check free movie
	if collection.TypeCode == "MOVIE" {
		theMovie, err := fs.movieRepository.FindMovieByID(ctx, collection.MovieID)
		if err != nil {
			return err
		}

		if theMovie.Price.Valid {
			return errors.ErrPaymentNotFound
		}

		err = fs.collectionRepository.InsertCollection(ctx, &model.Collection{
			Username:  author,
			EpisodeID: util.IntToSQLNullInt(int64(theMovie.ID)),
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
			return errors.ErrPaymentNotFound
		}

		err = fs.collectionRepository.InsertCollection(ctx, &model.Collection{
			Username:  author,
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

func (fs collectionService) GetCollectionsByUsernameAndType(ctx context.Context, movieType string, keyword string, pageRequest *pagination.PageRequest) (*pagination.Page[*dto.CollectionDetailDto], error) {
	if movieType == "" {
		return nil, errors.ErrInvalidInput
	}

	// Get author
	author := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	isValidUser, _ := fs.mgmtCtrl.CheckUser(author)
	if !isValidUser {
		return nil, errors.ErrUnAuthorized
	}

	page := &pagination.Page[*model.CollectionDetail]{}

	results, err := fs.collectionRepository.FindCollectionsByUsernameAndType(ctx, author, movieType, keyword, pageRequest, page)
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

func (fs collectionService) GetCollectionByUsernameAndRefID(ctx context.Context, username string, typeCode string, refID uint) (*dto.CollectionDto, error) {
	if typeCode == "" {
		return nil, errors.ErrInvalidInput
	}

	var result *model.Collection
	var err error
	if typeCode == "MOVIE" {
		result, err = fs.collectionRepository.FindCollectionByUsernameAndMovieID(ctx, username, refID)
		if err != nil {
			return nil, err
		}
	} else if typeCode == "TV" {
		result, err = fs.collectionRepository.FindCollectionByUsernameAndEpisodeID(ctx, username, refID)
		if err != nil {
			return nil, err
		}
	}

	return mapper.MapToCollectionDto(result), nil
}
