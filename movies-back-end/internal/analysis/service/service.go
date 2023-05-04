package service

import (
	"context"
	"movies-service/internal/analysis"
	"movies-service/internal/control"
	"movies-service/internal/dto"
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

func (as *analysisService) GetNumberOfMoviesByGenre(ctx context.Context) (*dto.ResultDto, error) {
	result, err := as.analysisRepository.CountMoviesByGenre(ctx)
	if err != nil {
		return nil, err
	}

	var dataSlice []*dto.DataDto
	for _, r := range result {
		dataSlice = append(dataSlice, &dto.DataDto{
			Genre: r.Genre,
			Count: r.NumMovies,
		})
	}

	return &dto.ResultDto{Data: dataSlice}, nil
}

func (as *analysisService) GetNumberOfMoviesByReleaseDate(ctx context.Context, year string, months []string) (*dto.ResultDto, error) {
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

func (as *analysisService) GetNumberOfViewsByGenreAndViewedDate(ctx context.Context, genre, year string, months []string) (*dto.ResultDto, error) {
	result, err := as.analysisRepository.CountViewsByGenreAndViewedDate(ctx, genre, year, months)
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
