package sqldb

import (
	"appix/internal/data/repo"
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
)

type DatacentersRepoGorm struct {
	data *DataGorm
	log  *log.Helper
}

func NewDatacentersRepoGorm(data *DataGorm, logger log.Logger) (repo.DatacentersRepo, error) {

	if err := validateData(data); err != nil {
		return nil, err
	}

	if err := initTable(data.DB, &repo.Datacenter{}, repo.DatacenterTable); err != nil {
		return nil, err
	}

	return &DatacentersRepoGorm{
		data: data,
		log:  log.NewHelper(logger),
	}, nil
}

// CreateDatacenters is
func (d *DatacentersRepoGorm) CreateDatacenters(ctx context.Context,
	tx repo.TX,
	dcs []*repo.Datacenter) error {

	r := d.data.WithTX(tx).WithContext(ctx).Create(dcs)
	if r.Error != nil {
		return r.Error
	}

	return nil
}

// UpdateDatacenters is
func (d *DatacentersRepoGorm) UpdateDatacenters(ctx context.Context,
	tx repo.TX,
	dcs []*repo.Datacenter) error {

	r := d.data.WithTX(tx).WithContext(ctx).Save(dcs)
	if r.Error != nil {
		return r.Error
	}

	return nil
}

// DeleteDatacenters is
func (d *DatacentersRepoGorm) DeleteDatacenters(
	ctx context.Context,
	tx repo.TX,
	ids []uint32) error {

	r := d.data.WithTX(tx).WithContext(ctx).Where("id in (?)", ids).Delete(&repo.Datacenter{})
	if r.Error != nil {
		return r.Error
	}
	if r.RowsAffected != int64(len(ids)) {
		return fmt.Errorf("delete failed. rows affected not equal wanted. affected %d. want %d",
			r.RowsAffected, len(ids))
	}

	return nil
}

// GetDatacenters is
func (d *DatacentersRepoGorm) GetDatacenters(ctx context.Context, id uint32) (*repo.Datacenter, error) {
	cs := &repo.Datacenter{}
	r := d.data.DB.WithContext(ctx).Where("id = ?", id).First(cs)
	if r.Error != nil {
		return nil, r.Error
	}

	return cs, nil
}

// ListDatacenters is
func (d *DatacentersRepoGorm) ListDatacenters(ctx context.Context,
	tx repo.TX,
	filter *repo.DatacentersFilter) ([]*repo.Datacenter, error) {

	dcs := []*repo.Datacenter{}
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
			for i, name := range filter.Names {
				params[i] = "%" + name + "%"
			}
			query = query.Where(nameConditions, params)
		}
	}
	r := query.Find(&dcs)
	if r.Error != nil {
		return nil, r.Error
	}

	return dcs, nil
}

// CountDatacenters is count by ids
func (d *DatacentersRepoGorm) CountDatacenters(ctx context.Context,
	tx repo.TX,
	filter repo.CountFilter) (int64, error) {

	var count int64
	query := d.data.WithTX(tx).WithContext(ctx)
	if filter != nil {
		if len(filter.GetIds()) > 0 {
			query = query.Where("id in (?)", filter.GetIds())
		}
	}
	r := query.Model(&repo.Datacenter{}).Count(&count)
	if r.Error != nil {
		return 0, r.Error
	}
	return count, nil
}

func (d *DatacentersRepoGorm) CountRequire(ctx context.Context,
	tx repo.TX,
	need repo.RequireType,
	ids []uint32) (int64, error) {

	if len(ids) == 0 {
		return 0, repo.ErrorRequireIds
	}

	// require nothing
	return 0, nil

}
