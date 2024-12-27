package sqldb

import (
	"appix/internal/data/repo"
	"context"

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

	return nil, nil
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
