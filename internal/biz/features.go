package biz

import (
	"context"
)

type FeaturesRepo interface {
	CreateFeatures(ctx context.Context, features []Feature) error
	UpdateFeatures(ctx context.Context, features []Feature) error
	DeleteFeatures(ctx context.Context, ids []string) error
	GetFeatures(ctx context.Context, id string) (Feature, error)
	ListFeatures(ctx context.Context, filter *ListFeaturesFilter) ([]Feature, error)
}

type FeaturesUsecase struct {
	repo FeaturesRepo
}

func NewFeaturesUsecase(repo FeaturesRepo) *FeaturesUsecase {
	return &FeaturesUsecase{
		repo: repo,
	}
}

// CreateFeatures is
func (s *FeaturesUsecase) CreateFeatures(ctx context.Context, features []Feature) error {
	return s.repo.CreateFeatures(ctx, features)
}

// UpdateFeatures is
func (s *FeaturesUsecase) UpdateFeatures(ctx context.Context, features []Feature) error {
	return s.repo.UpdateFeatures(ctx, features)
}

// DeleteFeatures is
func (s *FeaturesUsecase) DeleteFeatures(ctx context.Context, ids []string) error {
	return s.repo.DeleteFeatures(ctx, ids)
}

// GetFeatures is
func (s *FeaturesUsecase) GetFeatures(ctx context.Context, id string) (Feature, error) {
	return s.repo.GetFeatures(ctx, id)
}

// ListFeatures is
func (s *FeaturesUsecase) ListFeatures(ctx context.Context,
	filter *ListFeaturesFilter) ([]Feature, error) {
	return s.repo.ListFeatures(ctx, filter)
}
