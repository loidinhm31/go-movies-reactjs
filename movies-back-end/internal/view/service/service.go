package service

import (
	"context"
	"movies-service/internal/control"
	"movies-service/internal/dto"
	"movies-service/internal/errors"
	"movies-service/internal/view"
)

type viewService struct {
	mgmtCtrl       control.Service
	viewRepository view.Repository
}

func NewViewService(
	mgmtCtrl control.Service,
	viewRepository view.Repository,
) view.Service {
	return &viewService{
		mgmtCtrl:       mgmtCtrl,
		viewRepository: viewRepository,
	}
}

func (vs viewService) RecognizeViewForMovie(ctx context.Context, viewer *dto.Viewer) error {
	if viewer.Viewer == "" {
		return errors.ErrInvalidClient
	}

	if viewer.Viewer != "anonymous" {
		if !vs.mgmtCtrl.CheckUser(viewer.Viewer) {
			return errors.ErrInvalidClient
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
