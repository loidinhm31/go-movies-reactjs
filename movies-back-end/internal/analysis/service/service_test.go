package service

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"movies-service/internal/analysis"
	"movies-service/internal/dto"
	"movies-service/internal/errors"
	"movies-service/internal/model"
	"movies-service/internal/test/helper"
	"testing"
)

func initMock() (*helper.MockManagementCtrl, *helper.MockAnalysisRepository, analysis.Service) {
	mockCtrl := new(helper.MockManagementCtrl)
	mockRepo := new(helper.MockAnalysisRepository)

	analysisService := NewAnalysisService(mockCtrl, mockRepo)

	return mockCtrl, mockRepo, analysisService
}

func TestAnalysisService_GetNumberOfMoviesByGenre(t *testing.T) {
	t.Run("Unauthorized", func(t *testing.T) {
		mockCtrl, _, analysisService := initMock()

		mockCtrl.On("CheckPrivilege", mock.Anything).Return(false)

		// Call the method being tested
		result, err := analysisService.GetNumberOfMoviesByGenre(context.Background(), "TV")

		// Assert the expected error
		assert.Error(t, err)
		assert.Equal(t, errors.ErrUnAuthorized.Error(), err.Error())
		assert.Nil(t, result)
	})

	t.Run("Valid MovieType", func(t *testing.T) {
		mockCtrl, mockRepo, analysisService := initMock()

		mockCtrl.On("CheckPrivilege", mock.Anything).Return(true)

		mockRepo.On("CountMoviesByGenre", mock.Anything, "MOVIE").
			Return([]*model.GenreCount{
				{TypeCode: "MOVIE", Name: "Action", NumMovies: 10},
				{TypeCode: "MOVIE", Name: "Adventure", NumMovies: 5},
			}, nil)

		// Call the method being tested
		result, err := analysisService.GetNumberOfMoviesByGenre(context.Background(), "MOVIE")

		// Assert the expected result and error
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result.Data, 2)
		assert.Equal(t, "Action", result.Data[0].Name)
		assert.Equal(t, 10, result.Data[0].Count)
		assert.Equal(t, "Adventure", result.Data[1].Name)
		assert.Equal(t, 5, result.Data[1].Count)
	})

	t.Run("Error Counting Movie", func(t *testing.T) {
		mockCtrl, mockRepo, analysisService := initMock()

		mockCtrl.On("CheckPrivilege", mock.Anything).Return(true)

		mockRepo.On("CountMoviesByGenre", context.Background(), "TV").
			Return([]*model.GenreCount{}, fmt.Errorf("unexpected error occurred"))

		// Call the method being tested
		result, err := analysisService.GetNumberOfMoviesByGenre(context.Background(), "TV")

		// Assert the expected error
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

func TestAnalysisService_GetNumberOfMoviesByReleaseDate(t *testing.T) {
	t.Run("Unauthorized", func(t *testing.T) {
		mockCtrl, _, analysisService := initMock()

		mockCtrl.On("CheckPrivilege", mock.Anything).Return(false)

		// Call the method being tested
		result, err := analysisService.GetNumberOfMoviesByReleaseDate(context.Background(), "2022", []string{"6", "7", "8"})

		// Assert the expected error
		assert.Error(t, err)
		assert.Equal(t, errors.ErrUnAuthorized.Error(), err.Error())
		assert.Nil(t, result)
	})

	t.Run("Valid Year, Month", func(t *testing.T) {
		mockCtrl, mockRepo, analysisService := initMock()

		mockCtrl.On("CheckPrivilege", mock.Anything).Return(true)

		mockRepo.On("CountMoviesByReleaseDate", context.Background(), "2022", []string{"6", "7", "8"}).
			Return([]*model.MovieCount{
				{Year: "2022", Month: "6", NumMovies: 10},
				{Year: "2022", Month: "8", NumMovies: 5},
			}, nil)

		// Call the method being tested
		result, err := analysisService.GetNumberOfMoviesByReleaseDate(context.Background(), "2022", []string{"6", "7", "8"})

		// Assert the expected result and error
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result.Data, 2)
		assert.Equal(t, "2022", result.Data[0].Year)
		assert.Equal(t, "6", result.Data[0].Month)
		assert.Equal(t, 10, result.Data[0].Count)
		assert.Equal(t, "2022", result.Data[1].Year)
		assert.Equal(t, "8", result.Data[1].Month)
		assert.Equal(t, 5, result.Data[1].Count)
	})

	t.Run("Error Counting Movie", func(t *testing.T) {
		mockCtrl, mockRepo, analysisService := initMock()

		mockCtrl.On("CheckPrivilege", mock.Anything).Return(true)

		mockRepo.On("CountMoviesByReleaseDate", context.Background(), "2022", []string{"6", "7", "8"}).
			Return([]*model.MovieCount{}, fmt.Errorf("unexpected error occurred"))

		// Call the method being tested
		result, err := analysisService.GetNumberOfMoviesByReleaseDate(context.Background(), "2022", []string{"6", "7", "8"})

		// Assert the expected error
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

func TestAnalysisService_GetNumberOfMoviesByCreatedDate(t *testing.T) {
	t.Run("Unauthorized", func(t *testing.T) {
		mockCtrl, _, analysisService := initMock()

		// Set up the mock repository's behavior
		mockCtrl.On("CheckPrivilege", mock.Anything).Return(false)

		// Call the method being tested
		result, err := analysisService.GetNumberOfMoviesByCreatedDate(context.Background(), "2022", []string{"6", "7", "8"})

		// Assert the expected error
		assert.Error(t, err)
		assert.Equal(t, errors.ErrUnAuthorized.Error(), err.Error())
		assert.Nil(t, result)
	})

	t.Run("Valid Year, Month", func(t *testing.T) {
		mockCtrl, mockRepo, analysisService := initMock()

		mockCtrl.On("CheckPrivilege", mock.Anything).Return(true)

		mockRepo.On("CountMoviesByCreatedDate", context.Background(), "2022", []string{"6", "7", "8"}).
			Return([]*model.MovieCount{
				{Year: "2022", Month: "6", NumMovies: 10},
				{Year: "2022", Month: "8", NumMovies: 5},
			}, nil)

		// Call the method being tested
		result, err := analysisService.GetNumberOfMoviesByCreatedDate(context.Background(), "2022", []string{"6", "7", "8"})

		// Assert the expected result and error
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result.Data, 2)
		assert.Equal(t, "2022", result.Data[0].Year)
		assert.Equal(t, "6", result.Data[0].Month)
		assert.Equal(t, 10, result.Data[0].Count)
		assert.Equal(t, "2022", result.Data[1].Year)
		assert.Equal(t, "8", result.Data[1].Month)
		assert.Equal(t, 5, result.Data[1].Count)
	})

	t.Run("Error Counting Movie", func(t *testing.T) {
		mockCtrl, mockRepo, analysisService := initMock()

		mockCtrl.On("CheckPrivilege", mock.Anything).Return(true)

		mockRepo.On("CountMoviesByCreatedDate", context.Background(), "2022", []string{"6", "7", "8"}).
			Return([]*model.MovieCount{}, fmt.Errorf("unexpected error occurred"))

		// Call the method being tested
		result, err := analysisService.GetNumberOfMoviesByCreatedDate(context.Background(), "2022", []string{"6", "7", "8"})

		// Assert the expected error
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

func TestAnalysisService_GetNumberOfViewsByGenreAndViewedDate(t *testing.T) {
	t.Run("Unauthorized", func(t *testing.T) {
		mockCtrl, _, analysisService := initMock()

		mockCtrl.On("CheckPrivilege", mock.Anything).Return(false)

		// Call the method being tested
		result, err := analysisService.GetNumberOfViewsByGenreAndViewedDate(context.Background(), &dto.RequestData{})

		// Assert the expected error
		assert.Error(t, err)
		assert.Equal(t, errors.ErrUnAuthorized.Error(), err.Error())
		assert.Nil(t, result)
	})

	t.Run("Invalid Type Code", func(t *testing.T) {
		mockCtrl, _, analysisService := initMock()

		mockCtrl.On("CheckPrivilege", mock.Anything).Return(true)

		// Call the method being tested
		request := &dto.RequestData{
			TypeCode: "",
			Name:     "Action",
		}
		result, err := analysisService.GetNumberOfViewsByGenreAndViewedDate(context.Background(), request)

		// Assert the expected error
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Nil(t, result.Data)
	})

	t.Run("Valid Request", func(t *testing.T) {
		mockCtrl, mockRepo, analysisService := initMock()

		// Set up the mock repository's behavior
		expectedResult := []*model.ViewCount{
			{Year: "2023", Month: "1", NumViewers: 100},
			{Year: "2023", Month: "2", NumViewers: 150},
		}

		mockCtrl.On("CheckPrivilege", mock.Anything).Return(true)

		mockRepo.On("CountViewsByGenreAndViewedDate", mock.Anything, mock.Anything).
			Return(expectedResult, nil)

		// Call the method being tested
		request := &dto.RequestData{
			TypeCode: "MOVIE",
			Name:     "Action",
		}
		result, err := analysisService.GetNumberOfViewsByGenreAndViewedDate(context.Background(), request)

		// Assert the expected result and error
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result.Data, 2)
		assert.Equal(t, "2023", result.Data[0].Year)
		assert.Equal(t, "1", result.Data[0].Month)
		assert.Equal(t, 100, result.Data[0].Count)
		assert.Equal(t, "2023", result.Data[1].Year)
		assert.Equal(t, "2", result.Data[1].Month)
		assert.Equal(t, 150, result.Data[1].Count)
	})
}

func TestAnalysisService_GetCumulativeViewsByGenreAndViewedDate(t *testing.T) {
	t.Run("Unauthorized", func(t *testing.T) {
		mockCtrl, _, analysisService := initMock()

		mockCtrl.On("CheckPrivilege", mock.Anything).Return(false)

		// Call the method being tested
		result, err := analysisService.GetCumulativeViewsByGenreAndViewedDate(context.Background(), &dto.RequestData{})

		// Assert the expected error
		assert.Error(t, err)
		assert.Equal(t, errors.ErrUnAuthorized.Error(), err.Error())
		assert.Nil(t, result)
	})

	t.Run("Invalid Type Code", func(t *testing.T) {
		mockCtrl, _, analysisService := initMock()

		mockCtrl.On("CheckPrivilege", mock.Anything).Return(true)

		// Call the method being tested
		request := &dto.RequestData{
			TypeCode: "",
			Name:     "Action",
		}
		result, err := analysisService.GetCumulativeViewsByGenreAndViewedDate(context.Background(), request)

		// Assert the expected error
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Nil(t, result.Data)
	})

	t.Run("Valid Request", func(t *testing.T) {
		mockCtrl, mockRepo, analysisService := initMock()

		// Set up the mock repository's behavior
		expectedResult := []*model.ViewCount{
			{Year: "2023", Month: "1", NumViewers: 100, Cumulative: 100},
			{Year: "2023", Month: "2", NumViewers: 150, Cumulative: 250},
		}

		mockCtrl.On("CheckPrivilege", mock.Anything).Return(true)

		mockRepo.On("CountCumulativeViewsByGenreAndViewedDate", mock.Anything, mock.Anything).
			Return(expectedResult, nil)

		// Call the method being tested
		request := &dto.RequestData{
			TypeCode: "MOVIE",
			Name:     "Action",
		}
		result, err := analysisService.GetCumulativeViewsByGenreAndViewedDate(context.Background(), request)

		// Assert the expected result and error
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result.Data, 2)
		assert.Equal(t, "2023", result.Data[0].Year)
		assert.Equal(t, "1", result.Data[0].Month)
		assert.Equal(t, 100, result.Data[0].Count)
		assert.Equal(t, 100, result.Data[0].Cumulative)
		assert.Equal(t, "2023", result.Data[1].Year)
		assert.Equal(t, "2", result.Data[1].Month)
		assert.Equal(t, 150, result.Data[1].Count)
		assert.Equal(t, result.Data[0].Count+result.Data[1].Count, result.Data[1].Cumulative)
	})
}

func TestAnalysisService_GetNumberOfViewsByViewedDate(t *testing.T) {
	t.Run("Unauthorized", func(t *testing.T) {
		mockCtrl, _, analysisService := initMock()

		mockCtrl.On("CheckPrivilege", mock.Anything).Return(false)

		// Call the method being tested
		result, err := analysisService.GetNumberOfViewsByViewedDate(context.Background(), &dto.RequestData{})

		// Assert the expected error
		assert.Error(t, err)
		assert.Equal(t, errors.ErrUnAuthorized.Error(), err.Error())
		assert.Nil(t, result)
	})

	t.Run("Valid Request", func(t *testing.T) {
		mockCtrl, mockRepo, analysisService := initMock()

		// Set up the mock repository's behavior
		expectedResult := []*model.ViewCount{
			{Year: "2023", Month: "1", NumViewers: 100},
			{Year: "2023", Month: "2", NumViewers: 150},
		}

		mockCtrl.On("CheckPrivilege", mock.Anything).Return(true)

		mockRepo.On("CountViewsByViewedDate", mock.Anything, mock.Anything).
			Return(expectedResult, nil)

		// Call the method being tested
		request := &dto.RequestData{
			TypeCode: "MOVIE",
			Name:     "Action",
		}
		result, err := analysisService.GetNumberOfViewsByViewedDate(context.Background(), request)

		// Assert the expected result and error
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result.Data, 2)
		assert.Equal(t, "2023", result.Data[0].Year)
		assert.Equal(t, "1", result.Data[0].Month)
		assert.Equal(t, 100, result.Data[0].Count)
		assert.Equal(t, "2023", result.Data[1].Year)
		assert.Equal(t, "2", result.Data[1].Month)
		assert.Equal(t, 150, result.Data[1].Count)
	})
}

func TestAnalysisService_GetNumberOfMoviesByGenreAndReleasedDate(t *testing.T) {
	t.Run("Unauthorized", func(t *testing.T) {
		mockCtrl, _, analysisService := initMock()

		mockCtrl.On("CheckPrivilege", mock.Anything).Return(false)

		// Call the method being tested
		result, err := analysisService.GetNumberOfMoviesByGenreAndReleasedDate(context.Background(), &dto.RequestData{})

		// Assert the expected error
		assert.Error(t, err)
		assert.Equal(t, errors.ErrUnAuthorized.Error(), err.Error())
		assert.Nil(t, result)
	})

	t.Run("Invalid Type Code", func(t *testing.T) {
		mockCtrl, _, analysisService := initMock()

		mockCtrl.On("CheckPrivilege", mock.Anything).Return(true)

		// Call the method being tested
		request := &dto.RequestData{
			TypeCode: "",
			Name:     "Action",
		}
		result, err := analysisService.GetNumberOfMoviesByGenreAndReleasedDate(context.Background(), request)

		// Assert the expected error
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Nil(t, result.Data)
	})

	t.Run("Valid Request", func(t *testing.T) {
		mockCtrl, mockRepo, analysisService := initMock()

		// Set up the mock repository's behavior
		data := []*model.MovieCount{
			{Year: "2023", Month: "1", NumMovies: 100, Cumulative: 100},
			{Year: "2023", Month: "2", NumMovies: 150, Cumulative: 250},
		}

		mockCtrl.On("CheckPrivilege", mock.Anything).Return(true)

		mockRepo.On("CountMoviesByGenreAndReleasedDate", mock.Anything, mock.Anything).
			Return(data, nil)

		// Call the method being tested
		request := &dto.RequestData{
			TypeCode: "MOVIE",
			Name:     "Action",
		}
		result, err := analysisService.GetNumberOfMoviesByGenreAndReleasedDate(context.Background(), request)

		// Assert the expected result and error
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result.Data, 2)
		assert.Equal(t, "2023", result.Data[0].Year)
		assert.Equal(t, "1", result.Data[0].Month)
		assert.Equal(t, 100, result.Data[0].Count)
		assert.Equal(t, 100, result.Data[0].Cumulative)
		assert.Equal(t, "2023", result.Data[1].Year)
		assert.Equal(t, "2", result.Data[1].Month)
		assert.Equal(t, 150, result.Data[1].Count)
		assert.Equal(t, result.Data[0].Count+result.Data[1].Count, result.Data[1].Cumulative)
	})
}
