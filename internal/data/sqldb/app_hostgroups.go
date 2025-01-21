package sqldb

import (
	"appix/internal/data/repo"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type AppHostgroupsRepoGorm struct {
	data *DataGorm
	log  *log.Helper
}

func NewAppHostgroupsRepoGorm(data *DataGorm, logger log.Logger) (repo.AppHostgroupsRepo, error) {
	if err := validateData(data); err != nil {
		return nil, err
	}
	if err := initTable(data.DB, &repo.AppHostgroup{}, repo.AppHostgroupTable); err != nil {
		return nil, err
	}
	return &AppHostgroupsRepoGorm{
		data: data,
		log:  log.NewHelper(logger),
	}, nil
}

func (d *AppHostgroupsRepoGorm) CreateAppHostgroups(ctx context.Context,
	tx repo.TX,
	apps []*repo.AppHostgroup) error {
	if len(apps) == 0 {
		return nil
	}
	return d.data.WithTX(tx).WithContext(ctx).Create(apps).Error
}

func (d *AppHostgroupsRepoGorm) UpdateAppHostgroups(ctx context.Context,
	tx repo.TX,
	apps []*repo.AppHostgroup) error {
	if len(apps) == 0 {
		return nil
	}
	return d.data.WithTX(tx).WithContext(ctx).Updates(apps).Error
}

func (d *AppHostgroupsRepoGorm) DeleteAppHostgroups(ctx context.Context,
	tx repo.TX,
	ids []uint32) error {
	if len(ids) == 0 {
		return nil
	}
	return d.data.WithTX(tx).WithContext(ctx).Delete(&repo.AppHostgroup{}, ids).Error
}
func (d *AppHostgroupsRepoGorm) ListAppHostgroups(ctx context.Context,
	tx repo.TX,
	filter *repo.AppHostgroupsFilter) ([]*repo.AppHostgroup, error) {

	query := d.data.WithTX(tx).WithContext(ctx)
	if len(filter.Ids) > 0 {
		query = query.Where("id in (?)", filter.Ids)
	}
	if len(filter.AppIds) > 0 {
		query = query.Where("app_id in (?)", filter.AppIds)
	}
	if len(filter.HostgroupIds) > 0 {
		query = query.Where("hostgroup_id in (?)", filter.HostgroupIds)
	}
	if len(filter.KVs) > 0 {
		s_q, kvs := buildOrKV("app_id", "hostgroup_id", filter.KVs)
		query = query.Where(s_q, kvs...)
	}
	if filter.Page > 0 && filter.PageSize > 0 {
		offset := int(filter.PageSize * (filter.Page - 1))
		query = query.Offset(offset).Limit(int(filter.PageSize))
	}

	var apps []*repo.AppHostgroup
	if err := query.Find(&apps).Error; err != nil {
		return nil, err
	}

	return apps, nil
}

func (d *AppHostgroupsRepoGorm) DeleteAppHostgroupsByAppId(ctx context.Context,
	tx repo.TX,
	appids []uint32) error {
	if len(appids) == 0 {
		return nil
	}
	return d.data.WithTX(tx).
		WithContext(ctx).
		Delete(&repo.AppHostgroup{}, "app_id in (?)", appids).Error
}

func (d *AppHostgroupsRepoGorm) CountRequire(ctx context.Context,
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
	case repo.RequireHostgroup:
		condition = "hostgroup_id in (?)"
	default:
		return 0, repo.ErrorRequireIds
	}

	var count int64
	r := d.data.WithTX(tx).WithContext(ctx).Model(&repo.AppHostgroup{}).
		Where(condition, ids).Count(&count)
	if r.Error != nil {
		return 0, r.Error
	}

	return count, nil
}
