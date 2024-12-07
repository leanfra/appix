package data

import (
	"appix/internal/biz"
	"context"
	//  TODO: modify project name
	// biz "appix/internal/biz"
)

type FeaturesRepoImpl struct {
	data *Data
}

func NewFeaturesRepoImpl(data *Data) (biz.FeaturesRepo, error) {

	if data == nil {
		return nil, ErrEmptyDatabase
	}

	return &FeaturesRepoImpl{
		data: data,
	}, nil
}

// CreateFeatures is
func (d *FeaturesRepoImpl) CreateFeatures(ctx context.Context, features []biz.Feature) error {
	// TODO database operations

	return nil
}

// UpdateFeatures is
func (d *FeaturesRepoImpl) UpdateFeatures(ctx context.Context, features []biz.Feature) error {
	// TODO database operations

	return nil
}

// DeleteFeatures is
func (d *FeaturesRepoImpl) DeleteFeatures(ctx context.Context, ids []string) error {
	// TODO database operations

	return nil
}

// GetFeatures is
func (d *FeaturesRepoImpl) GetFeatures(ctx context.Context, id string) (*biz.Feature, error) {
	// TODO database operations

	return nil, nil
}

// ListFeatures is
func (d *FeaturesRepoImpl) ListFeatures(ctx context.Context, filter *biz.ListFeaturesFilter) ([]biz.Feature, error) {
	// TODO database operations

	return []biz.Feature{}, nil
}
