package data

import (
	"appix/internal/biz"
	"context"

	"github.com/go-kratos/kratos/v2/log"
	//  TODO: modify project name
	// biz "appix/internal/biz"
)

type DatacentersRepoImpl struct {
	data *Data
	log  *log.Helper
}

func NewDatacentersRepoImpl(data *Data, logger log.Logger) (biz.DatacentersRepo, error) {

	if err := validateData(data); err != nil {
		return nil, err
	}

	if err := initTable(data.db, &Datacenter{}, "datacenter"); err != nil {
		return nil, err
	}

	return &DatacentersRepoImpl{
		data: data,
		log:  log.NewHelper(logger),
	}, nil
}

// CreateDatacenters is
func (d *DatacentersRepoImpl) CreateDatacenters(ctx context.Context,
	dcs []biz.Datacenter) error {

	db_dcs, err := NewDatacenters(dcs)
	if err != nil {
		return err
	}
	r := d.data.db.WithContext(ctx).Create(db_dcs)
	if r.Error != nil {
		return r.Error
	}

	return nil
}

// UpdateDatacenters is
func (d *DatacentersRepoImpl) UpdateDatacenters(ctx context.Context,
	dcs []biz.Datacenter) error {

	db_cs, err := NewDatacenters(dcs)
	if err != nil {
		return err
	}
	r := d.data.db.WithContext(ctx).Save(db_cs)
	if r.Error != nil {
		return r.Error
	}

	return nil
}

// DeleteDatacenters is
func (d *DatacentersRepoImpl) DeleteDatacenters(ctx context.Context, ids []int64) error {

	r := d.data.db.WithContext(ctx).Where("id in (?)", ids).Delete(&Datacenter{})
	if r.Error != nil {
		return r.Error
	}

	return nil
}

// GetDatacenters is
func (d *DatacentersRepoImpl) GetDatacenters(ctx context.Context, id int64) (*biz.Datacenter, error) {
	cs := &Datacenter{}
	r := d.data.db.WithContext(ctx).Where("id = ?", id).First(cs)
	if r.Error != nil {
		return nil, r.Error
	}

	return NewBizDatacenter(cs)
}

// ListDatacenters is
func (d *DatacentersRepoImpl) ListDatacenters(ctx context.Context,
	filter *biz.ListDatacentersFilter) ([]biz.Datacenter, error) {

	dcs := []Datacenter{}
	query := d.data.db.WithContext(ctx)
	if filter != nil {
		var offset int
		if filter.Page > 0 && filter.PageSize > 0 {
			offset = int((filter.Page - 1) * filter.PageSize)
			query = query.Offset(offset).Limit(int(filter.PageSize))
		}
		for _, pair := range filter.Filters {
			query = query.Where("name =?", pair.Name)
		}
	}
	r := query.Find(&dcs)
	if r.Error != nil {
		return nil, r.Error
	}

	return NewBizDatacenters(dcs)
}
