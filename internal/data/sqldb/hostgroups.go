package sqldb

import (
	"opspillar/internal/data/repo"
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
)

type HostgroupsRepoGorm struct {
	data *DataGorm
	log  *log.Helper
}

func NewHostgroupsRepoGorm(data *DataGorm, logger log.Logger) (repo.HostgroupsRepo, error) {

	if err := validateData(data); err != nil {
		return nil, err
	}
	if err := initTable(data.DB, &repo.Hostgroup{}, repo.HostgroupTable); err != nil {
		return nil, err
	}
	return &HostgroupsRepoGorm{
		data: data,
		log:  log.NewHelper(logger),
	}, nil
}

// CreateHostgroups is
func (d *HostgroupsRepoGorm) CreateHostgroups(
	ctx context.Context,
	tx repo.TX,
	hgs []*repo.Hostgroup) error {

	r := d.data.WithTX(tx).WithContext(ctx).Create(hgs)
	if r.Error != nil {
		return r.Error
	}
	return nil
}

// UpdateHostgroups is
func (d *HostgroupsRepoGorm) UpdateHostgroups(
	ctx context.Context,
	tx repo.TX,
	hgs []*repo.Hostgroup) error {

	r := d.data.WithTX(tx).WithContext(ctx).Model(&repo.Hostgroup{}).Save(hgs)
	if r.Error != nil {
		return r.Error
	}

	return nil
}

// DeleteHostgroups is
func (d *HostgroupsRepoGorm) DeleteHostgroups(ctx context.Context, tx repo.TX, ids []uint32) error {

	r := d.data.WithTX(tx).WithContext(ctx).Where("id in (?)", ids).Delete(&repo.Hostgroup{})
	if r.Error != nil {
		return r.Error
	}
	if r.RowsAffected != int64(len(ids)) {
		return fmt.Errorf("delete failed. rows affected not equal wanted. affected %d. want %d", r.RowsAffected, len(ids))
	}
	return nil
}

// GetHostgroups is
func (d *HostgroupsRepoGorm) GetHostgroups(
	ctx context.Context, id uint32) (*repo.Hostgroup, error) {

	hg := &repo.Hostgroup{}
	r := d.data.DB.WithContext(ctx).Where("id = ?", id).First(hg)
	if r.Error != nil {
		return nil, r.Error
	}
	return hg, nil
}

// ListHostgroups is
func (d *HostgroupsRepoGorm) ListHostgroups(ctx context.Context,
	tx repo.TX,
	filter *repo.HostgroupsFilter) ([]*repo.Hostgroup, error) {

	query := d.data.WithTX(tx).WithContext(ctx).Model(&repo.Hostgroup{})
	if filter != nil {
		if len(filter.Ids) > 0 {
			query = query.Where("id in (?)", filter.Ids)
		}
		if len(filter.Names) > 0 {
			s_q := buildOrLike("name", len(filter.Names))
			params := make([]interface{}, len(filter.Names))
			for i, v := range filter.Names {
				params[i] = "%" + v + "%"
			}
			query = query.Where(s_q, params...)
		}
		if len(filter.ProductsId) > 0 {
			query = query.Where("product_id in (?)", filter.ProductsId)
		}
		if len(filter.DatacentersId) > 0 {
			query = query.Where("datacenter_id in (?)", filter.DatacentersId)
		}
		if len(filter.EnvsId) > 0 {
			query = query.Where("env_id in (?)", filter.EnvsId)
		}
		if len(filter.ClustersId) > 0 {
			query = query.Where("cluster_id in (?)", filter.ClustersId)
		}
		if len(filter.TeamsId) > 0 {
			query = query.Where("team_id in (?)", filter.TeamsId)
		}
		if filter.Page > 0 && filter.PageSize > 0 {
			offset := int((filter.Page - 1) * filter.PageSize)
			query = query.Offset(offset).Limit(int(filter.PageSize))
		}
	}
	var hgs []*repo.Hostgroup
	r := query.Find(&hgs)
	if r.Error != nil {
		return nil, r.Error
	}
	return hgs, nil
}

func (d *HostgroupsRepoGorm) CountHostgroups(ctx context.Context,
	tx repo.TX,
	filter repo.CountFilter) (int64, error) {

	var count int64
	query := d.data.WithTX(tx).WithContext(ctx).Model(&repo.Hostgroup{})
	if filter != nil {
		if len(filter.GetIds()) > 0 {
			query = query.Where("id in (?)", filter.GetIds())
		}
	}
	r := query.Count(&count)
	if r.Error != nil {
		return 0, r.Error
	}

	return count, nil
}

func (d *HostgroupsRepoGorm) CountRequire(ctx context.Context,
	tx repo.TX,
	need repo.RequireType,
	ids []uint32) (int64, error) {

	if len(ids) == 0 {
		return 0, repo.ErrorRequireIds
	}

	var condition string
	switch need {
	case repo.RequireCluster:
		condition = "cluster_id in (?)"
	case repo.RequireDatacenter:
		condition = "datacenter_id in (?)"
	case repo.RequireEnv:
		condition = "env_id in (?)"
	case repo.RequireProduct:
		condition = "product_id in (?)"
	case repo.RequireTeam:
		condition = "team_id in (?)"
	default:
		return 0, repo.ErrorRequireIds
	}

	var count int64
	r := d.data.WithTX(tx).WithContext(ctx).Model(&repo.Hostgroup{}).
		Where(condition, ids).Count(&count)
	if r.Error != nil {
		return 0, r.Error
	}
	// require nothing
	return count, nil
}
