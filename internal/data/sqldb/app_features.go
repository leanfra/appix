package sqldb

import (
	"appix/internal/data/repo"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type AppFeaturesRepoGorm struct {
	data *DataGorm
	log  *log.Helper
}

func NewAppFeaturesRepoGorm(data *DataGorm, logger log.Logger) (repo.AppFeaturesRepo, error) {
	if err := validateData(data); err != nil {
		return nil, err
	}
	if err := initTable(data.DB, &repo.AppFeature{}, repo.AppFeatureTable); err != nil {
		return nil, err
	}
	return &AppFeaturesRepoGorm{
		data: data,
		log:  log.NewHelper(logger),
	}, nil
}

func (d *AppFeaturesRepoGorm) CreateAppFeatures(ctx context.Context,
	tx repo.TX,
	apps []*repo.AppFeature) error {
	if len(apps) == 0 {
		return nil
	}
	return d.data.WithTX(tx).WithContext(ctx).Create(apps).Error
}

func (d *AppFeaturesRepoGorm) UpdateAppFeatures(ctx context.Context,
	tx repo.TX,
	apps []*repo.AppFeature) error {
	if len(apps) == 0 {
		return nil
	}
	return d.data.WithTX(tx).WithContext(ctx).Updates(apps).Error
}

func (d *AppFeaturesRepoGorm) DeleteAppFeatures(ctx context.Context,
	tx repo.TX,
	ids []uint32) error {
	if len(ids) == 0 {
		return nil
	}
	return d.data.WithTX(tx).WithContext(ctx).Delete(&repo.AppFeature{}, ids).Error
}
func (d *AppFeaturesRepoGorm) ListAppFeatures(ctx context.Context,
	tx repo.TX,
	filter *repo.AppFeaturesFilter) ([]*repo.AppFeature, error) {

	query := d.data.WithTX(tx).WithContext(ctx)
	if len(filter.Ids) > 0 {
		query = query.Where("id in (?)", filter.Ids)
	}
	if len(filter.AppIds) > 0 {
		query = query.Where("app_id in (?)", filter.AppIds)
	}
	if len(filter.FeatureIds) > 0 {
		query = query.Where("feature_id in (?)", filter.FeatureIds)
	}
	if len(filter.KVs) > 0 {
		s_q, kvs := buildOrKV("app_id", "feature_id", filter.KVs)
		query = query.Where(s_q, kvs...)
	}
	if filter.Page > 0 && filter.PageSize > 0 {
		offset := int(filter.PageSize * (filter.Page - 1))
		query = query.Offset(offset).Limit(int(filter.PageSize))
	}

	var apps []*repo.AppFeature
	if err := query.Find(&apps).Error; err != nil {
		return nil, err
	}

	return apps, nil
}

func (d *AppFeaturesRepoGorm) DeleteAppFeaturesByAppId(ctx context.Context,
	tx repo.TX,
	appids []uint32) error {
	if len(appids) == 0 {
		return nil
	}
	return d.data.WithTX(tx).WithContext(ctx).Delete(&repo.AppFeature{}, "app_id in (?)", appids).Error
}

func (d *AppFeaturesRepoGorm) CountRequire(ctx context.Context,
	tx repo.TX,
	need repo.RequireType,
	ids []uint32) (int64, error) {

	if len(ids) == 0 {
		return 0, repo.ErrorRequireIds
	}

	var condition string
	switch need {
	case repo.RequireApp:
		condition = "app_id in (?)"
	case repo.RequireFeature:
		condition = "feature_id in (?)"
	default:
		return 0, nil
	}

	var count int64
	r := d.data.WithTX(tx).WithContext(ctx).Model(&repo.AppFeature{}).
		Where(condition, ids).Count(&count)
	if r.Error != nil {
		return 0, r.Error
	}

	return count, nil
}
