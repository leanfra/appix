package sqldb

import (
	"appix/internal/data/repo"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type HostgroupTagsRepoGorm struct {
	data *DataGorm
	log  *log.Helper
}

func NewHostgroupTagsRepoGorm(data *DataGorm, logger log.Logger) (repo.HostgroupTagsRepo, error) {
	if err := validateData(data); err != nil {
		return nil, err
	}
	if err := initTable(data.DB, &repo.HostgroupTag{}, repo.HostgroupTagTable); err != nil {
		return nil, err
	}
	return &HostgroupTagsRepoGorm{
		data: data,
		log:  log.NewHelper(logger),
	}, nil
}

func (d *HostgroupTagsRepoGorm) CreateHostgroupTags(ctx context.Context,
	tx repo.TX,
	hostgroups []*repo.HostgroupTag) error {
	if len(hostgroups) == 0 {
		return nil
	}
	return d.data.DB.WithContext(ctx).Create(hostgroups).Error
}

func (d *HostgroupTagsRepoGorm) UpdateHostgroupTags(ctx context.Context,
	tx repo.TX,
	hostgroups []*repo.HostgroupTag) error {
	if len(hostgroups) == 0 {
		return nil
	}
	return d.data.DB.WithContext(ctx).Updates(hostgroups).Error
}

func (d *HostgroupTagsRepoGorm) DeleteHostgroupTags(ctx context.Context,
	tx repo.TX,
	ids []uint32) error {
	if len(ids) == 0 {
		return nil
	}
	return d.data.DB.WithContext(ctx).Delete(&repo.HostgroupTag{}, ids).Error
}
func (d *HostgroupTagsRepoGorm) ListHostgroupTags(ctx context.Context,
	tx repo.TX,
	filter *repo.HostgroupTagsFilter) ([]*repo.HostgroupTag, error) {

	query := d.data.WithTX(tx).WithContext(ctx).Model(&repo.HostgroupTag{})
	if len(filter.Ids) > 0 {
		query = query.Where("id in (?)", filter.Ids)
	}
	if len(filter.HostgroupIds) > 0 {
		query = query.Where("hostgroup_id in (?)", filter.HostgroupIds)
	}
	if len(filter.TagIds) > 0 {
		query = query.Where("feature_id in (?)", filter.TagIds)
	}
	if len(filter.KVs) > 0 {
		s_q, kvs := buildOrKV("hostgroup_id", "feature_id", filter.KVs)
		query = query.Where(s_q, kvs)
	}
	if filter.Page > 0 && filter.PageSize > 0 {
		offset := int(filter.PageSize * (filter.Page - 1))
		query = query.Offset(offset).Limit(int(filter.PageSize))
	}

	var hostgroups []*repo.HostgroupTag
	if err := query.Find(&hostgroups).Error; err != nil {
		return nil, err
	}

	return hostgroups, nil
}
