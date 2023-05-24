package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"movies-service/internal/dto"
	"movies-service/internal/errors"
	"movies-service/internal/models"
	"testing"
)

// Mock genreRepository for testing
type MockGenreRepository struct {
	mock.Mock
}

func (m *MockGenreRepository) FindAllGenresByTypeCode(ctx context.Context, movieType string) ([]*models.Genre, error) {
	args := m.Called(ctx, movieType)
	return args.Get(0).([]*models.Genre), args.Error(1)
}

func (m *MockGenreRepository) FindGenreByNameAndTypeCode(ctx context.Context, genre *models.Genre) (*models.Genre, error) {
	args := m.Called(ctx, genre)
	return args.Get(0).(*models.Genre), args.Error(1)
}

func (m *MockGenreRepository) FindAllGenres(ctx context.Context) ([]*models.Genre, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*models.Genre), args.Error(1)
}

func (m *MockGenreRepository) InsertGenres(ctx context.Context, genres []*models.Genre) error {
	args := m.Called(ctx, genres)
	return args.Error(0)
}

// MockManagementCtrl is a mock implementation of the Service interface for testing.
type MockManagementCtrl struct {
	mock.Mock
}

func (m *MockManagementCtrl) CheckPrivilege(username string) bool {
	args := m.Called(username)
	return args.Bool(0)
}

func (m *MockManagementCtrl) CheckAdminPrivilege(username string) bool {
	args := m.Called(username)
	return args.Bool(0)
}

func (m *MockManagementCtrl) CheckUser(username string) bool {
	args := m.Called(username)
	return args.Bool(0)
}

func setupMock() (*MockManagementCtrl, *MockGenreRepository, *genreService) {
	// Create a mock genre repository
	mockRepo := new(MockGenreRepository)

	// Create a mock management controller
	mockCtrl := new(MockManagementCtrl)

	// Create a genre service instance with the mock repository and controller
	genreService := &genreService{
		genreRepository: mockRepo,
		mgmtCtrl:        mockCtrl,
	}

	return mockCtrl, mockRepo, genreService
}

func TestGetAllGenresByTypeCode(t *testing.T) {

	t.Run("Valid movie type (action)", func(t *testing.T) {
		_, mockRepo, genreService := setupMock()

		// Set up mock expectations and return values
		mockRepo.On("FindAllGenresByTypeCode", context.Background(), "action").
			Return([]*models.Genre{
				{Name: "Action1", TypeCode: "MOVIE"},
				{Name: "Action2", TypeCode: "MOVIE"},
			}, nil)

		genres, err := genreService.GetAllGenresByTypeCode(context.Background(), "action")
		assert.NoError(t, err)
		assert.Equal(t, 2, len(genres))
		assert.Equal(t, "Action1", genres[0].Name)
		assert.Equal(t, "Action2", genres[1].Name)
	})

	t.Run("Invalid movie type", func(t *testing.T) {
		_, mockRepo, genreService := setupMock()

		// Set up mock expectations and return values
		mockRepo.On("FindAllGenres", context.Background()).
			Return([]*models.Genre{
				{Name: "Comedy1", TypeCode: "MOVIE"},
				{Name: "Comedy2", TypeCode: "TV"},
			}, nil)

		genres, err := genreService.GetAllGenresByTypeCode(context.Background(), "")
		assert.NoError(t, err)
		assert.Equal(t, 2, len(genres))
		assert.Equal(t, "Comedy1", genres[0].Name)
		assert.Equal(t, "MOVIE", genres[0].TypeCode)
		assert.Equal(t, "Comedy2", genres[1].Name)
		assert.Equal(t, "TV", genres[1].TypeCode)
	})
}

func TestAddGenres(t *testing.T) {

	t.Run("Valid genres and privileged user", func(t *testing.T) {
		mockCtrl, mockRepo, genreService := setupMock()

		// Set up mock expectations and return values
		mockCtrl.On("CheckPrivilege", mock.Anything).Return(true)

		genreDtos := []dto.GenreDto{
			{Name: "Genre3", TypeCode: "MOVIE"},
			{Name: "Genre4", TypeCode: "TV"},
		}

		mockRepo.On("FindGenreByNameAndTypeCode", context.Background(), &models.Genre{
			Name:     genreDtos[0].Name,
			TypeCode: genreDtos[0].TypeCode,
		}).
			Return(&models.Genre{}, nil)

		mockRepo.On("FindGenreByNameAndTypeCode", context.Background(), &models.Genre{
			Name:     genreDtos[1].Name,
			TypeCode: genreDtos[1].TypeCode,
		}).
			Return(&models.Genre{}, nil)

		mockRepo.On("InsertGenres", context.Background(), mock.Anything).
			Return(nil)
		err := genreService.AddGenres(context.Background(), genreDtos)
		assert.NoError(t, err)
	})

	t.Run("Genre with existing name and type code", func(t *testing.T) {
		mockCtrl, mockRepo, genreService := setupMock()

		// Set up mock expectations and return values
		mockCtrl.On("CheckPrivilege", mock.Anything).Return(true)

		genreDtos := []dto.GenreDto{
			{Name: "Action1", TypeCode: "MOVIE"},
			{Name: "Genre2", TypeCode: "TV"},
		}

		mockRepo.On("FindGenreByNameAndTypeCode", context.Background(), &models.Genre{
			Name:     genreDtos[0].Name,
			TypeCode: genreDtos[0].TypeCode,
		}).
			Return(&models.Genre{}, nil)

		mockRepo.On("FindGenreByNameAndTypeCode", context.Background(), &models.Genre{
			Name:     genreDtos[1].Name,
			TypeCode: genreDtos[1].TypeCode,
		}).
			Return(&models.Genre{Name: genreDtos[1].Name, TypeCode: genreDtos[1].TypeCode}, nil)

		err := genreService.AddGenres(context.Background(), genreDtos)
		assert.Error(t, err)
		assert.Equal(t, errors.ErrCannotExecuteAction.Error(), err.Error())
	})

	t.Run("Invalid genre type code", func(t *testing.T) {
		mockCtrl, mockRepo, genreService := setupMock()

		// Set up mock expectations and return values
		mockCtrl.On("CheckPrivilege", mock.Anything).Return(true)

		mockRepo.On("FindGenreByNameAndTypeCode", context.Background(), mock.Anything).
			Return(&models.Genre{}, nil)

		genreDtos := []dto.GenreDto{
			{Name: "Genre1", TypeCode: "INVALID"},
		}
		err := genreService.AddGenres(context.Background(), genreDtos)
		assert.Error(t, err)
		assert.Equal(t, errors.ErrInvalidInput.Error(), err.Error())
	})

	t.Run("Unauthorized user", func(t *testing.T) {
		mockCtrl, _, genreService := setupMock()

		// Set up mock expectations and return values
		mockCtrl.On("CheckPrivilege", mock.Anything).Return(false)

		genreDtos := []dto.GenreDto{
			{Name: "Genre1", TypeCode: "MOVIE"},
		}
		err := genreService.AddGenres(context.Background(), genreDtos)
		assert.Error(t, err)
		assert.Equal(t, errors.ErrUnAuthorized.Error(), err.Error())
	})
}
