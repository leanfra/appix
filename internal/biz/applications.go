package biz

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
)

type ApplicationsRepo interface {
	CreateApplications(ctx context.Context, apps []*Application) error
	UpdateApplications(ctx context.Context, apps []*Application) error
	DeleteApplications(ctx context.Context, ids []uint32) error
	GetApplications(ctx context.Context, id uint32) (*Application, error)
	ListApplications(ctx context.Context, filter *ListApplicationsFilter) ([]*Application, error)
}

type ApplicationsUsecase struct {
	repo ApplicationsRepo
	log  *log.Helper
}

func NewApplicationsUsecase(repo ApplicationsRepo, logger log.Logger) *ApplicationsUsecase {
	return &ApplicationsUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (s *ApplicationsUsecase) validate(isNew bool, apps []*Application) error {
	for _, a := range apps {
		if err := a.Validate(isNew); err != nil {
			return err
		}
	}
	return nil
}

// CreateApplications is
func (s *ApplicationsUsecase) CreateApplications(ctx context.Context, apps []*Application) error {
	if err := s.validate(true, nil); err != nil {
		return err
	}
	return s.repo.CreateApplications(ctx, apps)
}

// UpdateApplications is
func (s *ApplicationsUsecase) UpdateApplications(ctx context.Context, apps []*Application) error {
	if err := s.validate(false, nil); err != nil {
		return err
	}
	return s.repo.UpdateApplications(ctx, apps)
}

// DeleteApplications is
func (s *ApplicationsUsecase) DeleteApplications(ctx context.Context, ids []uint32) error {
	if len(ids) == 0 {
		return fmt.Errorf("EmptyIds")
	}
	return s.repo.DeleteApplications(ctx, ids)
}

// GetApplications is
func (s *ApplicationsUsecase) GetApplications(ctx context.Context, id uint32) (*Application, error) {
	if id <= 0 {
		return nil, fmt.Errorf("InvalidId")
	}
	return s.repo.GetApplications(ctx, id)
}

// ListApplications is
func (s *ApplicationsUsecase) ListApplications(ctx context.Context, filter *ListApplicationsFilter) ([]*Application, error) {
	if filter != nil {
		if err := filter.Validate(); err != nil {
			return nil, err
		}
	}
	return s.repo.ListApplications(ctx, filter)
}
