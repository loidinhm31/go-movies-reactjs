package service

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"movies-service/internal/blob"
	"movies-service/internal/collection"
	"movies-service/internal/control"
	"movies-service/internal/dto"
	"movies-service/internal/errors"
	"movies-service/internal/mapper"
	"movies-service/internal/middlewares"
	"movies-service/internal/model"
	"movies-service/internal/movie"
	"movies-service/internal/payment"
	"movies-service/pkg/pagination"
	"strings"
	"sync"
)

type movieService struct {
	mgmtCtrl             control.Service
	movieRepository      movie.Repository
	collectionRepository collection.Repository
	paymentRepository    payment.Repository
	blobService          blob.Service
}

func NewMovieService(mgmtCtrl control.Service, movieRepository movie.Repository, collectionRepository collection.Repository, paymentRepository payment.Repository, blobService blob.Service) movie.Service {
	return &movieService{
		mgmtCtrl:             mgmtCtrl,
		movieRepository:      movieRepository,
		collectionRepository: collectionRepository,
		paymentRepository:    paymentRepository,
		blobService:          blobService,
	}
}

func (ms *movieService) GetAllMoviesByType(ctx context.Context, keyword, movieType string, pageRequest *pagination.PageRequest) (*pagination.Page[*dto.MovieDto], error) {
	page := &pagination.Page[*model.Movie]{}

	var err error
	var movieResults *pagination.Page[*model.Movie]
	if movieType != "" {
		movieResults, err = ms.movieRepository.FindAllMoviesByType(ctx, keyword, movieType, pageRequest, page)
		if err != nil {
			log.Println(err)
			return nil, errors.ErrResourceNotFound
		}
	} else {
		movieResults, err = ms.movieRepository.FindAllMovies(ctx, keyword, pageRequest, page)
		if err != nil {
			log.Println(err)
			return nil, errors.ErrResourceNotFound
		}
	}

	movieDtos := mapper.MapToMovieDtoSlice(movieResults.Content)

	return &pagination.Page[*dto.MovieDto]{
		PageSize:      pageRequest.PageSize,
		PageNumber:    pageRequest.PageNumber,
		Sort:          pageRequest.Sort,
		TotalElements: movieResults.TotalElements,
		TotalPages:    movieResults.TotalPages,
		Content:       movieDtos,
	}, nil
}

func (ms *movieService) GetMovieByID(ctx context.Context, id uint) (*dto.MovieDto, error) {
	author := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	isValidUser, isPrivilege := ms.mgmtCtrl.CheckUser(author)

	result, err := ms.movieRepository.FindMovieByID(ctx, id)
	if err != nil {
		log.Println(err)
		return nil, errors.ErrResourceNotFound
	}

	// Check valid video path
	if result.Price.Valid {
		thePayment, err := ms.paymentRepository.FindByTypeCodeAndRefID(ctx, result.TypeCode, result.ID)
		if err != nil {
			return nil, err
		}

		if !(thePayment.TypeCode == result.TypeCode && thePayment.RefID == result.ID) {
			theCollection, err := ms.collectionRepository.FindCollectionByMovieID(ctx, result.ID)
			if err != nil {
				return nil, err
			}

			if !(theCollection.TypeCode == result.TypeCode && uint(theCollection.MovieID.Int64) == result.ID) {
				result.VideoPath = sql.NullString{}
			}
		}
	}

	movieDto := mapper.MapToMovieDto(result, !isValidUser, isPrivilege)
	return movieDto, nil
}

func (ms *movieService) GetMoviesByGenre(ctx context.Context, pageRequest *pagination.PageRequest, genreId uint) (*pagination.Page[*dto.MovieDto], error) {
	page := &pagination.Page[*model.Movie]{}

	movieResults, err := ms.movieRepository.FindMoviesByGenre(ctx, pageRequest, page, genreId)
	if err != nil {
		log.Println(err)
		return nil, errors.ErrResourceNotFound
	}

	movieDtos := mapper.MapToMovieDtoSlice(movieResults.Content)

	return &pagination.Page[*dto.MovieDto]{
		PageSize:      pageRequest.PageSize,
		PageNumber:    pageRequest.PageNumber,
		Sort:          pageRequest.Sort,
		TotalElements: movieResults.TotalElements,
		TotalPages:    movieResults.TotalPages,
		Content:       movieDtos,
	}, nil
}

func (ms *movieService) AddMovie(ctx context.Context, movie *dto.MovieDto) error {
	if movie.ID > 0 ||
		movie.Title == "" ||
		movie.TypeCode == "" ||
		movie.Runtime == 0 ||
		movie.Description == "" ||
		movie.ReleaseDate.IsZero() ||
		movie.MpaaRating == "" ||
		(movie.Genres == nil || len(movie.Genres) == 0) {
		return errors.ErrInvalidInput
	}

	// Get author
	author := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	if !ms.mgmtCtrl.CheckPrivilege(author) {
		return errors.ErrUnAuthorized
	}

	var genreObjects []*model.Genre
	for _, genre := range movie.Genres {
		if genre.Checked {
			if genre.TypeCode != movie.TypeCode {
				log.Println("genre type must equal movie type")
				return errors.ErrInvalidInput
			}
			genreObjects = append(genreObjects, mapper.MapToGenre(genre, author))
		}
	}

	if len(genreObjects) == 0 {
		log.Println("empty genres")
		return errors.ErrInvalidInputDetail("genres cannot empty")
	}

	movieObject := mapper.MapToMovie(movie, author)
	movieObject.Genres = genreObjects

	err := ms.movieRepository.InsertMovie(ctx, movieObject)
	if err != nil {
		return err
	}
	return nil
}

func (ms *movieService) UpdateMovie(ctx context.Context, movie *dto.MovieDto) error {
	if movie.ID == 0 ||
		movie.Title == "" ||
		movie.TypeCode == "" ||
		movie.Runtime == 0 ||
		movie.Description == "" ||
		movie.ReleaseDate.IsZero() ||
		(movie.Genres == nil || len(movie.Genres) == 0) {
		return errors.ErrInvalidInput
	}

	// Check movie exists
	movieObj, err := ms.movieRepository.FindMovieByID(ctx, movie.ID)
	if err != nil {
		return errors.ErrResourceNotFound
	}

	// Get author
	author := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	if !ms.mgmtCtrl.CheckPrivilege(author) {
		return errors.ErrUnAuthorized
	}

	// After check object exists, write updating value
	movieObj = mapper.MapToMovieUpdate(movie, author)

	err = ms.movieRepository.UpdateMovie(ctx, movieObj)
	if err != nil {
		return err
	}

	var genreObjects []*model.Genre
	for _, genre := range movie.Genres {
		if genre.Checked {
			if genre.TypeCode != movie.TypeCode {
				log.Println("genre type must equal movie type")
				return errors.ErrInvalidInput
			}
			genreObjects = append(genreObjects, mapper.MapToGenre(genre, author))
		}
	}

	if len(genreObjects) == 0 {
		log.Println("empty genres")
		return errors.ErrInvalidInputDetail("genres cannot empty")
	}

	err = ms.movieRepository.UpdateMovieGenres(ctx, movieObj, genreObjects)
	if err != nil {
		return err
	}
	return nil
}

func (ms *movieService) RemoveMovieByID(ctx context.Context, id uint) error {
	if id == 0 {
		return errors.ErrInvalidInput
	}

	// Check collection
	log.Println("checking collection before deleting...")
	collections, err := ms.collectionRepository.FindCollectionsByMovieID(ctx, id)
	if err != nil {
		return err
	}

	if len(collections) > 0 {
		return errors.ErrCannotExecuteAction
	}

	// Get author
	author := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	if !ms.mgmtCtrl.CheckPrivilege(author) {
		return errors.ErrUnAuthorized
	}

	// Get current movie
	movieObj, err := ms.movieRepository.FindMovieByID(ctx, id)
	if err != nil {
		return err
	}

	// Delete video from blob concurrently
	if movieObj.VideoPath.Valid {
		var wg sync.WaitGroup
		wg.Add(1)

		go func() {
			defer wg.Done()
			videoPath := movieObj.VideoPath.String
			videoPathSplit := strings.Split(videoPath, "/")
			videoKey := videoPathSplit[len(videoPathSplit)-1]
			res, err := ms.blobService.DeleteFile(ctx, videoKey, "video")
			if err != nil {
				log.Println("cannot delete video")
			}
			log.Println(res)
		}()
		wg.Wait()
	}

	// Delete image from blob concurrently
	if movieObj.ImageUrl.Valid {
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			imageUrl := movieObj.ImageUrl.String
			imageUrlSplit := strings.Split(imageUrl, "/")
			imageFile := imageUrlSplit[len(imageUrlSplit)-1]
			imageKey := strings.Split(imageFile, ".")[0]
			res, err := ms.blobService.DeleteFile(ctx, imageKey, "image")
			if err != nil {
				log.Println("cannot delete image")
			}
			log.Println(res)

		}()
		wg.Wait()
	}

	err = ms.movieRepository.DeleteMovieByID(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (ms *movieService) GetMovieByEpisodeID(ctx context.Context, episodeID uint) (*dto.MovieDto, error) {
	author := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	isValidUser, isPrivilege := ms.mgmtCtrl.CheckUser(author)

	result, err := ms.movieRepository.FindMovieByEpisodeID(ctx, episodeID)
	if err != nil {
		log.Println(err)
		return nil, errors.ErrResourceNotFound
	}

	movieDto := mapper.MapToMovieDto(result, !isValidUser, isPrivilege)
	return movieDto, nil
}
