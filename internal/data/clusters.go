package data

import (
	"appix/internal/biz"
	"context"

	"github.com/go-kratos/kratos/v2/log"
	//  TODO: modify project name
	// biz "appix/internal/biz"
)

type ClustersRepoImpl struct {
	data *Data
	log  *log.Helper
}

func NewClustersRepoImpl(data *Data, logger log.Logger) (biz.ClustersRepo, error) {

	if err := validateData(data); err != nil {
		return nil, err
	}

	if err := initTable(data.db, &Cluster{}, "cluster"); err != nil {
		return nil, err
	}

	return &ClustersRepoImpl{
		data: data,
		log:  log.NewHelper(logger),
	}, nil
}

// CreateClusters is
func (d *ClustersRepoImpl) CreateClusters(ctx context.Context, cs []biz.Cluster) error {

	db_cs, err := NewClusters(cs)
	if err != nil {
		return err
	}
	r := d.data.db.WithContext(ctx).Create(db_cs)
	if r.Error != nil {
		return r.Error
	}
	return nil
}

// UpdateClusters is
func (d *ClustersRepoImpl) UpdateClusters(ctx context.Context, cs []biz.Cluster) error {

	db_cs, err := NewClusters(cs)
	if err != nil {
		return err
	}
	r := d.data.db.WithContext(ctx).Save(db_cs)
	if r.Error != nil {
		return r.Error
	}
	return nil
}

// DeleteClusters is
func (d *ClustersRepoImpl) DeleteClusters(ctx context.Context, ids []int64) error {

	r := d.data.db.WithContext(ctx).Where("id in (?)", ids).Delete(&Cluster{})
	if r.Error != nil {
		return r.Error
	}
	return nil
}

// GetClusters is
func (d *ClustersRepoImpl) GetClusters(ctx context.Context, id int64) (*biz.Cluster, error) {

	cs := &Cluster{}
	r := d.data.db.WithContext(ctx).Where("id = ?", id).First(cs)
	if r.Error != nil {
		return nil, r.Error
	}
	return NewBizCluster(cs)
}

// ListClusters is
func (d *ClustersRepoImpl) ListClusters(ctx context.Context, filter *biz.ListClustersFilter) ([]biz.Cluster, error) {
	cs := []Cluster{}
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
	r := query.Find(&cs)
	if r.Error != nil {
		return nil, r.Error
	}
	return NewBizClusters(cs)
}
