package service

import (
	"context"
	"fmt"
	"movies-service/internal/collection"
	"movies-service/internal/control"
	"movies-service/internal/dto"
	"movies-service/internal/errors"
	"movies-service/internal/mapper"
	"movies-service/internal/middlewares"
	"movies-service/internal/model"
	"movies-service/pkg/pagination"
	"time"
)

type collectionService struct {
	mgmtCtrl             control.Service
	collectionRepository collection.Repository
}

func NewCollectionService(mgmtCtrl control.Service, collectionRepository collection.Repository) collection.Service {
	return &collectionService{
		mgmtCtrl:             mgmtCtrl,
		collectionRepository: collectionRepository,
	}
}

func (fs collectionService) AddCollection(ctx context.Context, movieID uint) error {
	if movieID == 0 {
		return errors.ErrInvalidInput
	}

	// Get author
	author := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	isValidUser, _ := fs.mgmtCtrl.CheckUser(author)
	if !isValidUser {
		return errors.ErrUnAuthorized
	}

	err := fs.collectionRepository.InsertCollection(ctx, &model.Collection{
		Username:  author,
		MovieID:   movieID,
		CreatedAt: time.Now(),
		CreatedBy: author,
	})
	if err != nil {
		return err
	}
	return nil
}

func (fs collectionService) GetCollectionsByUsernameAndType(ctx context.Context, movieType string, keyword string, pageRequest *pagination.PageRequest) (*pagination.Page[*dto.CollectionDto], error) {
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
	return &pagination.Page[*dto.CollectionDto]{
		PageSize:      pageRequest.PageSize,
		PageNumber:    pageRequest.PageNumber,
		Sort:          pageRequest.Sort,
		TotalElements: results.TotalElements,
		TotalPages:    results.TotalPages,
		Content:       collectionDtos,
	}, nil
}

func (fs collectionService) GetCollectionByUsernameAndMovieID(ctx context.Context, username string, movieID uint) (*dto.CollectionDto, error) {
	results, err := fs.collectionRepository.FindCollectionByUsernameAndMovieID(ctx, username, movieID)
	if err != nil {
		return nil, err
	}
	return mapper.MapToCollectionDto(results), nil
}
