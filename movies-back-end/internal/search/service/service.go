package service

import (
	"context"
	"movies-service/internal/common/dto"
	"movies-service/internal/common/mapper"
	"movies-service/internal/common/model"
	"movies-service/internal/search"
	"movies-service/pkg/pagination"
)

type searchService struct {
	searchRepository search.Repository
}

func NewSearchService(searchRepository search.Repository) search.Service {
	return &searchService{
		searchRepository,
	}
}

func (gs *searchService) SearchMovie(ctx context.Context, searchParams *model.SearchParams) (*pagination.Page[*dto.MovieDto], error) {
	page, err := gs.searchRepository.SearchMovie(ctx, searchParams)
	if err != nil {
		return nil, err
	}

	movieDtos := mapper.MapToMovieDtoSlice(page.Content)

	return &pagination.Page[*dto.MovieDto]{
		PageSize:      searchParams.Page.PageSize,
		PageNumber:    searchParams.Page.PageNumber,
		Sort:          searchParams.Page.Sort,
		TotalElements: page.TotalElements,
		TotalPages:    page.TotalPages,
		Content:       movieDtos,
	}, nil
}
