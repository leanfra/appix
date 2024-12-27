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
	if err := initTable(data.db, &repo.AppHostgroup{}, repo.AppHostgroupTable); err != nil {
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
	return d.data.db.WithContext(ctx).Create(apps).Error
}

func (d *AppHostgroupsRepoGorm) UpdateAppHostgroups(ctx context.Context,
	tx repo.TX,
	apps []*repo.AppHostgroup) error {
	if len(apps) == 0 {
		return nil
	}
	return d.data.db.WithContext(ctx).Updates(apps).Error
}

func (d *AppHostgroupsRepoGorm) DeleteAppHostgroups(ctx context.Context,
	tx repo.TX,
	ids []uint32) error {
	if len(ids) == 0 {
		return nil
	}
	return d.data.db.WithContext(ctx).Delete(&repo.AppHostgroup{}, ids).Error
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
		query = query.Where("feature_id in (?)", filter.HostgroupIds)
	}
	if len(filter.KVs) > 0 {
		s_q, kvs := buildOrKV("app_id", "feature_id", filter.KVs)
		query = query.Where(s_q, kvs)
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
