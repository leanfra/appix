package biz

import (
	"appix/internal/data/repo"
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
)

type FeaturesUsecase struct {
	repo repo.FeaturesRepo
	log  *log.Helper
	txm  repo.TxManager
}

func NewFeaturesUsecase(repo repo.FeaturesRepo, logger log.Logger, txm repo.TxManager) *FeaturesUsecase {
	return &FeaturesUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
		txm:  txm,
	}
}

func (s *FeaturesUsecase) validate(isNew bool, features []*Feature) error {
	for _, f := range features {
		if err := f.Validate(isNew); err != nil {
			return err
		}
	}
	return nil
}

// CreateFeatures is
func (s *FeaturesUsecase) CreateFeatures(ctx context.Context, features []*Feature) error {
	if err := s.validate(true, features); err != nil {
		return err
	}
	_f, e := ToDBFeatures(features)
	if e != nil {
		return e
	}
	return s.repo.CreateFeatures(ctx, _f)
}

// UpdateFeatures is
func (s *FeaturesUsecase) UpdateFeatures(ctx context.Context, features []*Feature) error {
	if err := s.validate(false, features); err != nil {
		return err
	}
	_f, e := ToDBFeatures(features)
	if e != nil {
		return e
	}
	return s.repo.UpdateFeatures(ctx, _f)
}

// DeleteFeatures is
func (s *FeaturesUsecase) DeleteFeatures(ctx context.Context, ids []uint32) error {
	if len(ids) == 0 {
		return fmt.Errorf("EmptyIds")
	}
	return s.repo.DeleteFeatures(ctx, ids)
}

// GetFeatures is
func (s *FeaturesUsecase) GetFeatures(ctx context.Context, id uint32) (*Feature, error) {
	if id <= 0 {
		return nil, fmt.Errorf("EmptyId")
	}
	_f, e := s.repo.GetFeatures(ctx, id)
	if e != nil {
		return nil, e
	}
	return ToBizFeature(_f)
}

// ListFeatures is
func (s *FeaturesUsecase) ListFeatures(ctx context.Context,
	filter *ListFeaturesFilter) ([]*Feature, error) {

	if filter != nil {
		if err := filter.Validate(); err != nil {
			return nil, err
		}
	}
	_f, e := s.repo.ListFeatures(ctx, nil, ToDBFeaturesFilter(filter))
	if e != nil {
		return nil, e
	}
	return ToBizFeatures(_f)
}
