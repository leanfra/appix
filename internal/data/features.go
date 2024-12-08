package data

import (
	"appix/internal/biz"
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	//  TODO: modify project name
	// biz "appix/internal/biz"
)

type FeaturesRepoImpl struct {
	data *Data
	log  *log.Helper
}

func NewFeaturesRepoImpl(data *Data, logger log.Logger) (biz.FeaturesRepo, error) {

	if err := validateData(data); err != nil {
		return nil, err
	}

	if err := initTable(data.db, &Feature{}, "feature"); err != nil {
		return nil, err
	}

	return &FeaturesRepoImpl{
		data: data,
		log:  log.NewHelper(logger),
	}, nil
}

// XXX all data passed in should be validated.

// CreateFeatures is
func (d *FeaturesRepoImpl) CreateFeatures(ctx context.Context, features []biz.Feature) error {

	db_ft, err := NewFeatures(features)
	if err != nil {
		return err
	}
	r := d.data.db.WithContext(ctx).Create(db_ft)
	if r.Error != nil {
		return r.Error
	}

	return nil
}

// UpdateFeatures is
func (d *FeaturesRepoImpl) UpdateFeatures(ctx context.Context, features []biz.Feature) error {

	db_fts, err := NewFeatures(features)
	if err != nil {
		return err
	}
	r := d.data.db.WithContext(ctx).Save(db_fts)
	if r.Error != nil {
		return r.Error
	}

	return nil
}

// DeleteFeatures is
func (d *FeaturesRepoImpl) DeleteFeatures(ctx context.Context, ids []int64) error {

	r := d.data.db.WithContext(ctx).Where("id in (?)", ids).Delete(&Feature{})
	if r.Error != nil {
		return r.Error
	}
	return nil
}

// GetFeatures is
func (d *FeaturesRepoImpl) GetFeatures(ctx context.Context, id int64) (*biz.Feature, error) {

	feature := &Feature{}
	r := d.data.db.WithContext(ctx).First(feature, id)
	if r.Error != nil {
		return nil, r.Error
	}
	return NewBizFeature(feature)

}

// ListFeatures is
func (d *FeaturesRepoImpl) ListFeatures(ctx context.Context, filter *biz.ListFeaturesFilter) ([]biz.Feature, error) {

	features := []Feature{}

	var r *gorm.DB
	query := d.data.db.WithContext(ctx)
	if filter != nil {
		var offset int
		if filter.Page > 0 && filter.PageSize > 0 {
			offset = int((filter.Page - 1) * filter.PageSize)
			query = query.Offset(offset).Limit(int(filter.PageSize))
		}

		for _, pair := range filter.Filters {
			query = query.Where("name =? AND value =?", pair.Name, pair.Value)
		}
	}
	r = query.Find(&features)

	if r.Error != nil {
		return nil, r.Error
	}
	return NewBizFeatures(features)
}
