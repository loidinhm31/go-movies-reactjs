package service

import (
	"context"
	"movies-service/internal/dto"
	"movies-service/internal/models"
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

func (gs *searchService) Search(ctx context.Context, searchParams *models.SearchParams) (*pagination.Page[*dto.MovieDto], error) {
	page, err := gs.searchRepository.Search(ctx, searchParams)
	if err != nil {
		return nil, err
	}

	var movieDtos []*dto.MovieDto
	for _, m := range page.Content {
		var genreDtos []*dto.GenreDto
		if m.Genres != nil {
			for _, g := range m.Genres {
				genreDtos = append(genreDtos, &dto.GenreDto{
					ID:   g.ID,
					Name: g.Name,
				})
			}
		}

		movieDtos = append(movieDtos, &dto.MovieDto{
			ID:          m.ID,
			Title:       m.Title,
			ReleaseDate: m.ReleaseDate,
			Runtime:     m.Runtime,
			MpaaRating:  m.MpaaRating,
			Description: m.Description,
			ImagePath:   m.ImagePath.String,
			CreatedAt:   m.CreatedAt,
			UpdatedAt:   m.UpdatedAt,
			Genres:      genreDtos,
		})
	}
	return &pagination.Page[*dto.MovieDto]{
		PageSize:      searchParams.Page.PageSize,
		PageNumber:    searchParams.Page.PageNumber,
		Sort:          searchParams.Page.Sort,
		TotalElements: page.TotalElements,
		TotalPages:    page.TotalPages,
		Content:       movieDtos,
	}, nil
}
