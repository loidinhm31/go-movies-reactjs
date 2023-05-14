package service

import (
	"context"
	"errors"
	"movies-service/internal/control"
	"movies-service/internal/dto"
	"movies-service/internal/views"
)

type viewService struct {
	mgmtCtrl       control.Service
	viewRepository views.Repository
}

func NewViewService(
	mgmtCtrl control.Service,
	viewRepository views.Repository,
) views.Service {
	return &viewService{
		mgmtCtrl:       mgmtCtrl,
		viewRepository: viewRepository,
	}
}

func (vs viewService) RecognizeViewForMovie(ctx context.Context, viewer *dto.Viewer) error {
	if viewer.Viewer == "" {
		return errors.New("invalid")
	}

	if viewer.Viewer != "anonymous" {
		if !vs.mgmtCtrl.CheckUser(viewer.Viewer) {
			return errors.New("invalid user")
		}
	}

	err := vs.viewRepository.InsertView(ctx, viewer)
	if err != nil {
		return err
	}
	return nil
}

func (vs viewService) GetNumberOfViewsByMovieId(ctx context.Context, movieId int) (int64, error) {
	totalViews, err := vs.viewRepository.CountViewsByMovieId(ctx, movieId)
	if err != nil {
		return 0, err
	}
	return totalViews, nil
}
