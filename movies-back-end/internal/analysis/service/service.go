package service

import (
	"context"
	"fmt"
	"movies-service/internal/analysis"
	"movies-service/internal/common/dto"
	"movies-service/internal/control"
	"movies-service/internal/errors"
	"movies-service/internal/middlewares"
)

type analysisService struct {
	mgmtCtrl           control.Service
	analysisRepository analysis.Repository
}

func NewAnalysisService(mgmtCtrl control.Service, analysisRepository analysis.Repository) analysis.Service {
	return &analysisService{
		mgmtCtrl:           mgmtCtrl,
		analysisRepository: analysisRepository,
	}
}

func (as *analysisService) GetNumberOfMoviesByGenre(ctx context.Context, movieType string) (*dto.ResultDto, error) {
	// Get author
	author := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	if !as.mgmtCtrl.CheckPrivilege(author) {
		return nil, errors.ErrUnAuthorized
	}

	result, err := as.analysisRepository.CountMoviesByGenre(ctx, movieType)
	if err != nil {
		return nil, err
	}

	var dataSlice []*dto.DataDto
	for _, r := range result {
		dataSlice = append(dataSlice, &dto.DataDto{
			Name:     r.Name,
			TypeCode: r.TypeCode,
			Count:    r.NumMovies,
		})
	}

	return &dto.ResultDto{Data: dataSlice}, nil
}

func (as *analysisService) GetNumberOfMoviesByReleaseDate(ctx context.Context, year string, months []string) (*dto.ResultDto, error) {
	// Get author
	author := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	if !as.mgmtCtrl.CheckPrivilege(author) {
		return nil, errors.ErrUnAuthorized
	}

	result, err := as.analysisRepository.CountMoviesByReleaseDate(ctx, year, months)
	if err != nil {
		return nil, err
	}

	var dataSlice []*dto.DataDto
	for _, r := range result {
		dataSlice = append(dataSlice, &dto.DataDto{
			Year:  r.Year,
			Month: r.Month,
			Count: r.NumMovies,
		})
	}
	return &dto.ResultDto{Data: dataSlice}, nil
}

func (as *analysisService) GetNumberOfMoviesByCreatedDate(ctx context.Context, year string, months []string) (*dto.ResultDto, error) {
	// Get author
	author := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	if !as.mgmtCtrl.CheckPrivilege(author) {
		return nil, errors.ErrUnAuthorized
	}

	result, err := as.analysisRepository.CountMoviesByCreatedDate(ctx, year, months)
	if err != nil {
		return nil, err
	}

	var dataSlice []*dto.DataDto
	for _, r := range result {
		dataSlice = append(dataSlice, &dto.DataDto{
			Year:  r.Year,
			Month: r.Month,
			Count: r.NumMovies,
		})
	}
	return &dto.ResultDto{Data: dataSlice}, nil
}

func (as *analysisService) GetNumberOfViewsByGenreAndViewedDate(ctx context.Context, request *dto.RequestData) (*dto.ResultDto, error) {
	// Get author
	author := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	if !as.mgmtCtrl.CheckPrivilege(author) {
		return nil, errors.ErrUnAuthorized
	}

	if request.TypeCode == "" {
		return &dto.ResultDto{Data: nil}, nil
	}

	result, err := as.analysisRepository.CountViewsByGenreAndViewedDate(ctx, request)
	if err != nil {
		return nil, err
	}

	var dataSlice []*dto.DataDto
	for _, r := range result {
		dataSlice = append(dataSlice, &dto.DataDto{
			Name:  request.Name,
			Year:  r.Year,
			Month: r.Month,
			Count: r.NumViewers,
		})
	}
	return &dto.ResultDto{Data: dataSlice}, nil
}

func (as *analysisService) GetCumulativeViewsByGenreAndViewedDate(ctx context.Context, request *dto.RequestData) (*dto.ResultDto, error) {
	// Get author
	author := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	if !as.mgmtCtrl.CheckPrivilege(author) {
		return nil, errors.ErrUnAuthorized
	}

	if request.TypeCode == "" {
		return &dto.ResultDto{Data: nil}, nil
	}

	result, err := as.analysisRepository.CountCumulativeViewsByGenreAndViewedDate(ctx, request)
	if err != nil {
		return nil, err
	}

	var dataSlice []*dto.DataDto
	for _, r := range result {
		dataSlice = append(dataSlice, &dto.DataDto{
			Name:       request.Name,
			Year:       r.Year,
			Month:      r.Month,
			Count:      r.NumViewers,
			Cumulative: r.Cumulative,
		})
	}
	return &dto.ResultDto{Data: dataSlice}, nil
}

func (as *analysisService) GetNumberOfViewsByViewedDate(ctx context.Context, request *dto.RequestData) (*dto.ResultDto, error) {
	// Get author
	author := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	if !as.mgmtCtrl.CheckPrivilege(author) {
		return nil, errors.ErrUnAuthorized
	}

	result, err := as.analysisRepository.CountViewsByViewedDate(ctx, request)
	if err != nil {
		return nil, err
	}

	var dataSlice []*dto.DataDto
	for _, r := range result {
		dataSlice = append(dataSlice, &dto.DataDto{
			Year:  r.Year,
			Month: r.Month,
			Count: r.NumViewers,
		})
	}
	return &dto.ResultDto{Data: dataSlice}, nil
}

func (as *analysisService) GetNumberOfMoviesByGenreAndReleasedDate(ctx context.Context, request *dto.RequestData) (*dto.ResultDto, error) {
	// Get author
	author := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	if !as.mgmtCtrl.CheckPrivilege(author) {
		return nil, errors.ErrUnAuthorized
	}

	if request.TypeCode == "" {
		return &dto.ResultDto{Data: nil}, nil
	}

	result, err := as.analysisRepository.CountMoviesByGenreAndReleasedDate(ctx, request)
	if err != nil {
		return nil, err
	}

	var dataSlice []*dto.DataDto
	for _, r := range result {
		dataSlice = append(dataSlice, &dto.DataDto{
			Name:       request.Name,
			Year:       r.Year,
			Month:      r.Month,
			Count:      r.NumMovies,
			Cumulative: r.Cumulative,
		})
	}
	return &dto.ResultDto{Data: dataSlice}, nil
}
