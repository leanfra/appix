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
	return d.data.WithTX(tx).WithContext(ctx).Create(hostgroups).Error
}

func (d *HostgroupFeaturesRepoGorm) UpdateHostgroupFeatures(ctx context.Context,
	tx repo.TX,
	hostgroups []*repo.HostgroupFeature) error {
	if len(hostgroups) == 0 {
		return nil
	}
	return d.data.WithTX(tx).WithContext(ctx).Updates(hostgroups).Error
}

func (d *HostgroupFeaturesRepoGorm) DeleteHostgroupFeatures(ctx context.Context,
	tx repo.TX,
	ids []uint32) error {
	if len(ids) == 0 {
		return nil
	}
	return d.data.WithTX(tx).WithContext(ctx).Delete(&repo.HostgroupFeature{}, ids).Error
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
		query = query.Where(s_q, kvs...)
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

func (d *HostgroupFeaturesRepoGorm) CountRequire(ctx context.Context,
	tx repo.TX,
	need repo.RequireType,
	ids []uint32) (int64, error) {

	if len(ids) == 0 {
		return 0, repo.ErrorRequireIds
	}

	var condition string
	switch need {
	case repo.RequireHostgroup:
		condition = "hostgroup_id in (?)"
	case repo.RequireFeature:
		condition = "feature_id in (?)"
	default:
		return 0, repo.ErrorRequireIds
	}

	var count int64
	r := d.data.WithTX(tx).WithContext(ctx).Model(&repo.HostgroupFeature{}).
		Where(condition, ids).Count(&count)
	if r.Error != nil {
		return 0, r.Error
	}
	// require nothing
	return count, nil
}

func (d *HostgroupFeaturesRepoGorm) ListHostgroupMatchFeatures(ctx context.Context,
	tx repo.TX,
	filter *repo.HostgroupMatchFeaturesFilter) ([]uint32, error) {

	if len(filter.FeatureIds) == 0 {
		return nil, ErrRequireFeatureIds
	}

	var hostgroupIds []uint32
	query := d.data.WithTX(tx).WithContext(ctx).Model(&repo.HostgroupFeature{})
	// Find hostgroups that have all the specified features
	query = query.
		Where("feature_id IN (?)", filter.FeatureIds).
		Group("hostgroup_id").
		Having("COUNT(DISTINCT feature_id) >= ?", len(filter.FeatureIds))

	if err := query.Pluck("hostgroup_id", &hostgroupIds).Error; err != nil {
		return nil, err
	}

	return hostgroupIds, nil

}
