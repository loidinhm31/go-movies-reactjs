package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"movies-service/internal/model"
	"movies-service/internal/test/helper"
	"movies-service/pkg/pagination"
	"testing"
)

func TestSearchMovie(t *testing.T) {
	// Create a new instance of the mock search.Repository.
	mockRepo := new(helper.MockSearchRepository)

	// Create an instance of the searchService using the mock repository.
	service := NewSearchService(mockRepo)

	// Create test data for searchParams and expected results.
	searchParams := &model.SearchParams{
		Filters: []*model.FieldData{
			{
				Operator: "and",
				Field:    "title",
				TypeValue: model.TypeValue{
					Type:   "string",
					Values: []string{"example"},
				},
			},
		},
		Page: &pagination.PageRequest{
			PageSize:   10,
			PageNumber: 1,
			Sort: pagination.Sort{
				Orders: []*pagination.Order{
					{
						Property:  "title",
						Direction: pagination.DESC,
					},
				},
			},
		},
	}

	expectedPage := &pagination.Page[*model.Movie]{
		PageSize:      searchParams.Page.PageSize,
		PageNumber:    searchParams.Page.PageNumber,
		Sort:          searchParams.Page.Sort,
		TotalElements: 2,
		TotalPages:    1,
		Content: []*model.Movie{
			{Title: "M1"},
			{Title: "M2"},
		},
	}

	// Set up the expectations for the mock repository's SearchMovie method.
	mockRepo.On("SearchMovie", mock.Anything, searchParams).
		Return(expectedPage, nil)

	// Call the SearchMovie method of the searchService.
	result, err := service.SearchMovie(context.Background(), searchParams)

	// Assert that the mock repository's SearchMovie method was called with the correct parameters.
	mockRepo.AssertCalled(t, "SearchMovie", mock.Anything, searchParams)

	// Assert that the result and error match the expected values.
	assert.NoError(t, err)
	assert.Equal(t, "M1", result.Content[0].Title)
	assert.Equal(t, "M2", result.Content[1].Title)
	assert.Equal(t, 1, result.TotalPages)
	assert.Equal(t, int64(2), result.TotalElements)

}
