package biz

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
)

type FeaturesRepo interface {
	CreateFeatures(ctx context.Context, features []Feature) error
	UpdateFeatures(ctx context.Context, features []Feature) error
	DeleteFeatures(ctx context.Context, ids []int64) error
	GetFeatures(ctx context.Context, id int64) (*Feature, error)
	ListFeatures(ctx context.Context, filter *ListFeaturesFilter) ([]Feature, error)
}

type FeaturesUsecase struct {
	repo FeaturesRepo
	log  *log.Helper
}

func NewFeaturesUsecase(repo FeaturesRepo, logger log.Logger) *FeaturesUsecase {
	return &FeaturesUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (s *FeaturesUsecase) validate(isNew bool, features []Feature) error {
	for _, f := range features {
		if err := f.Validate(isNew); err != nil {
			return err
		}
	}
	return nil
}

// CreateFeatures is
func (s *FeaturesUsecase) CreateFeatures(ctx context.Context, features []Feature) error {
	if err := s.validate(true, features); err != nil {
		return err
	}
	return s.repo.CreateFeatures(ctx, features)
}

// UpdateFeatures is
func (s *FeaturesUsecase) UpdateFeatures(ctx context.Context, features []Feature) error {
	if err := s.validate(false, features); err != nil {
		return err
	}
	return s.repo.UpdateFeatures(ctx, features)
}

// DeleteFeatures is
func (s *FeaturesUsecase) DeleteFeatures(ctx context.Context, ids []int64) error {
	if len(ids) == 0 {
		return fmt.Errorf("EmptyIds")
	}
	return s.repo.DeleteFeatures(ctx, ids)
}

// GetFeatures is
func (s *FeaturesUsecase) GetFeatures(ctx context.Context, id int64) (*Feature, error) {
	if id <= 0 {
		return nil, fmt.Errorf("EmptyId")
	}
	return s.repo.GetFeatures(ctx, id)
}

// ListFeatures is
func (s *FeaturesUsecase) ListFeatures(ctx context.Context,
	filter *ListFeaturesFilter) ([]Feature, error) {

	if filter != nil {
		if err := filter.Validate(); err != nil {
			return nil, err
		}
	}
	return s.repo.ListFeatures(ctx, filter)
}
