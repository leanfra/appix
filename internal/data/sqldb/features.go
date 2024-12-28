package sqldb

import (
	"appix/internal/data/repo"
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type FeaturesRepoGorm struct {
	data *DataGorm
	log  *log.Helper
}

func NewFeaturesRepoGorm(data *DataGorm, logger log.Logger) (repo.FeaturesRepo, error) {

	if err := validateData(data); err != nil {
		return nil, err
	}

	if err := initTable(data.DB, &repo.Feature{}, repo.FeatureTable); err != nil {
		return nil, err
	}

	return &FeaturesRepoGorm{
		data: data,
		log:  log.NewHelper(logger),
	}, nil
}

// XXX all data passed in should be validated.

// CreateFeatures is
func (d *FeaturesRepoGorm) CreateFeatures(ctx context.Context, features []*repo.Feature) error {

	r := d.data.DB.WithContext(ctx).Create(features)
	if r.Error != nil {
		return r.Error
	}

	return nil
}

// UpdateFeatures is
func (d *FeaturesRepoGorm) UpdateFeatures(ctx context.Context, features []*repo.Feature) error {

	r := d.data.DB.WithContext(ctx).Save(features)
	if r.Error != nil {
		return r.Error
	}

	return nil
}

// DeleteFeatures is
func (d *FeaturesRepoGorm) DeleteFeatures(ctx context.Context, ids []uint32) error {

	r := d.data.DB.WithContext(ctx).Where("id in (?)", ids).Delete(&repo.Feature{})
	if r.Error != nil {
		return r.Error
	}
	if r.RowsAffected != int64(len(ids)) {
		return fmt.Errorf("delete failed. rows affected not equal wanted. affected %d. want %d",
			r.RowsAffected, len(ids))
	}
	return nil
}

// GetFeatures is
func (d *FeaturesRepoGorm) GetFeatures(ctx context.Context, id uint32) (*repo.Feature, error) {

	feature := &repo.Feature{}
	r := d.data.DB.WithContext(ctx).First(feature, id)
	if r.Error != nil {
		return nil, r.Error
	}
	return feature, nil

}

// ListFeatures is
func (d *FeaturesRepoGorm) ListFeatures(ctx context.Context,
	tx repo.TX,
	filter *repo.FeaturesFilter) ([]*repo.Feature, error) {

	features := []*repo.Feature{}

	var r *gorm.DB
	query := d.data.WithTX(tx).WithContext(ctx)
	if filter != nil {
		var offset int
		if filter.Page > 0 && filter.PageSize > 0 {
			offset = int((filter.Page - 1) * filter.PageSize)
			query = query.Offset(offset).Limit(int(filter.PageSize))
		}
		if len(filter.Ids) > 0 {
			query = query.Where("id in (?)", filter.Ids)
		}
		if len(filter.Names) > 0 {
			nameConditions := buildOrLike("name", len(filter.Names))
			params := make([]interface{}, len(filter.Names))
			for i, v := range filter.Names {
				params[i] = "%" + v + "%"
			}
			query = query.Where(nameConditions, params...)
		}
		if len(filter.Kvs) > 0 {
			kvConditions, kvs := buildOrKV("name", "value", filter.Kvs)
			query = query.Where(kvConditions, kvs...)
		}
	}
	r = query.Find(&features)

	if r.Error != nil {
		return nil, r.Error
	}
	return features, nil
}

func (d *FeaturesRepoGorm) CountFeatures(ctx context.Context,
	tx repo.TX,
	filter repo.CountFilter) (int64, error) {

	var count int64
	query := d.data.WithTX(tx).WithContext(ctx)
	if filter != nil {
		if len(filter.GetIds()) > 0 {
			query = query.Where("id in (?)", filter.GetIds())
		}
	}
	r := query.Model(&repo.Feature{}).Count(&count)
	if r.Error != nil {
		return 0, r.Error
	}
	return count, nil
}

func (d *FeaturesRepoGorm) CountRequire(ctx context.Context,
	tx repo.TX,
	need repo.RequireType,
	ids []uint32) (int64, error) {

	if len(ids) == 0 {
		return 0, repo.ErrorRequireIds
	}

	// require nothing
	return 0, nil

}
