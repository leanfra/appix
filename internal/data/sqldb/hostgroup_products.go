package sqldb

import (
	"appix/internal/data/repo"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type HostgroupProductsRepoGorm struct {
	data *DataGorm
	log  *log.Helper
}

func NewHostgroupProductsRepoGorm(data *DataGorm, logger log.Logger) (repo.HostgroupProductsRepo, error) {
	if err := validateData(data); err != nil {
		return nil, err
	}
	if err := initTable(data.db, &repo.HostgroupProduct{}, repo.HostgroupProductTable); err != nil {
		return nil, err
	}
	return &HostgroupProductsRepoGorm{
		data: data,
		log:  log.NewHelper(logger),
	}, nil
}

func (d *HostgroupProductsRepoGorm) CreateHostgroupProducts(ctx context.Context,
	tx repo.TX,
	hostgroups []*repo.HostgroupProduct) error {
	if len(hostgroups) == 0 {
		return nil
	}
	return d.data.db.WithContext(ctx).Create(hostgroups).Error
}

func (d *HostgroupProductsRepoGorm) UpdateHostgroupProducts(ctx context.Context,
	tx repo.TX,
	hostgroups []*repo.HostgroupProduct) error {
	if len(hostgroups) == 0 {
		return nil
	}
	return d.data.db.WithContext(ctx).Updates(hostgroups).Error
}

func (d *HostgroupProductsRepoGorm) DeleteHostgroupProducts(ctx context.Context,
	tx repo.TX,
	ids []uint32) error {
	if len(ids) == 0 {
		return nil
	}
	return d.data.db.WithContext(ctx).Delete(&repo.HostgroupProduct{}, ids).Error
}
func (d *HostgroupProductsRepoGorm) ListHostgroupProducts(ctx context.Context,
	tx repo.TX,
	filter *repo.HostgroupProductsFilter) ([]*repo.HostgroupProduct, error) {

	query := d.data.WithTX(tx).WithContext(ctx).Model(&repo.HostgroupProduct{})
	if len(filter.Ids) > 0 {
		query = query.Where("id in (?)", filter.Ids)
	}
	if len(filter.HostgroupIds) > 0 {
		query = query.Where("hostgroup_id in (?)", filter.HostgroupIds)
	}
	if len(filter.ProductIds) > 0 {
		query = query.Where("feature_id in (?)", filter.ProductIds)
	}
	if len(filter.KVs) > 0 {
		s_q, kvs := buildOrKV("hostgroup_id", "feature_id", filter.KVs)
		query = query.Where(s_q, kvs)
	}
	if filter.Page > 0 && filter.PageSize > 0 {
		offset := int(filter.PageSize * (filter.Page - 1))
		query = query.Offset(offset).Limit(int(filter.PageSize))
	}

	var hostgroups []*repo.HostgroupProduct
	if err := query.Find(&hostgroups).Error; err != nil {
		return nil, err
	}

	return hostgroups, nil
}
