package sqldb

import (
	"appix/internal/data/repo"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type HostgroupFeaturesRepoGorm struct {
	data *DataGorm
	log  *log.Helper
}

func NewHostgroupFeaturesRepoGorm(data *DataGorm, logger log.Logger) (repo.HostgroupFeaturesRepo, error) {
	if err := validateData(data); err != nil {
		return nil, err
	}
	if err := initTable(data.DB, &repo.HostgroupFeature{}, repo.HostgroupFeatureTable); err != nil {
		return nil, err
	}
	return &HostgroupFeaturesRepoGorm{
		data: data,
		log:  log.NewHelper(logger),
	}, nil
}

func (d *HostgroupFeaturesRepoGorm) CreateHostgroupFeatures(ctx context.Context,
	tx repo.TX,
	hostgroups []*repo.HostgroupFeature) error {
	if len(hostgroups) == 0 {
		return nil
	}
	return d.data.DB.WithContext(ctx).Create(hostgroups).Error
}

func (d *HostgroupFeaturesRepoGorm) UpdateHostgroupFeatures(ctx context.Context,
	tx repo.TX,
	hostgroups []*repo.HostgroupFeature) error {
	if len(hostgroups) == 0 {
		return nil
	}
	return d.data.DB.WithContext(ctx).Updates(hostgroups).Error
}

func (d *HostgroupFeaturesRepoGorm) DeleteHostgroupFeatures(ctx context.Context,
	tx repo.TX,
	ids []uint32) error {
	if len(ids) == 0 {
		return nil
	}
	return d.data.DB.WithContext(ctx).Delete(&repo.HostgroupFeature{}, ids).Error
}
func (d *HostgroupFeaturesRepoGorm) ListHostgroupFeatures(ctx context.Context,
	tx repo.TX,
	filter *repo.HostgroupFeaturesFilter) ([]*repo.HostgroupFeature, error) {

	query := d.data.WithTX(tx).WithContext(ctx).Model(&repo.HostgroupFeature{})
	if len(filter.Ids) > 0 {
		query = query.Where("id in (?)", filter.Ids)
	}
	if len(filter.HostgroupIds) > 0 {
		query = query.Where("hostgroup_id in (?)", filter.HostgroupIds)
	}
	if len(filter.FeatureIds) > 0 {
		query = query.Where("feature_id in (?)", filter.FeatureIds)
	}
	if len(filter.KVs) > 0 {
		s_q, kvs := buildOrKV("hostgroup_id", "feature_id", filter.KVs)
		query = query.Where(s_q, kvs)
	}
	if filter.Page > 0 && filter.PageSize > 0 {
		offset := int(filter.PageSize * (filter.Page - 1))
		query = query.Offset(offset).Limit(int(filter.PageSize))
	}

	var hostgroups []*repo.HostgroupFeature
	if err := query.Find(&hostgroups).Error; err != nil {
		return nil, err
	}

	return hostgroups, nil
}
