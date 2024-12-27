package sqldb

import (
	"appix/internal/data/repo"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type ClustersRepoGorm struct {
	data *DataGorm
	log  *log.Helper
}

func NewClustersRepoGorm(data *DataGorm, logger log.Logger) (repo.ClustersRepo, error) {

	if err := validateData(data); err != nil {
		return nil, err
	}

	if err := initTable(data.DB, &repo.Cluster{}, repo.ClusterTable); err != nil {
		return nil, err
	}

	return &ClustersRepoGorm{
		data: data,
		log:  log.NewHelper(logger),
	}, nil
}

// CreateClusters is
func (d *ClustersRepoGorm) CreateClusters(ctx context.Context, cs []*repo.Cluster) error {

	r := d.data.DB.WithContext(ctx).Create(cs)
	if r.Error != nil {
		return r.Error
	}
	return nil
}

// UpdateClusters is
func (d *ClustersRepoGorm) UpdateClusters(ctx context.Context, cs []*repo.Cluster) error {

	r := d.data.DB.WithContext(ctx).Save(cs)
	if r.Error != nil {
		return r.Error
	}
	return nil
}

// DeleteClusters is
func (d *ClustersRepoGorm) DeleteClusters(ctx context.Context, ids []uint32) error {

	r := d.data.DB.WithContext(ctx).Where("id in (?)", ids).Delete(&repo.Cluster{})
	if r.Error != nil {
		return r.Error
	}
	return nil
}

// GetClusters is
func (d *ClustersRepoGorm) GetClusters(ctx context.Context, id uint32) (*repo.Cluster, error) {

	cs := &repo.Cluster{}
	r := d.data.DB.WithContext(ctx).Where("id = ?", id).First(cs)
	if r.Error != nil {
		return nil, r.Error
	}
	return cs, nil
}

// ListClusters is
func (d *ClustersRepoGorm) ListClusters(ctx context.Context,
	tx repo.TX,
	filter *repo.ClustersFilter) ([]*repo.Cluster, error) {
	cs := []*repo.Cluster{}
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
			query = query.Where(nameConditions, filter.Names)
		}
	}
	r := query.Find(&cs)
	if r.Error != nil {
		return nil, r.Error
	}
	return cs, nil
}

func (d *ClustersRepoGorm) CountClusters(ctx context.Context,
	tx repo.TX,
	filter repo.CountFilter) (int64, error) {

	var count int64
	query := d.data.WithTX(tx).WithContext(ctx)
	if filter != nil {
		if len(filter.GetIds()) > 0 {
			query = query.Where("id in (?)", filter.GetIds())
		}
	}
	r := query.Model(&repo.Cluster{}).Count(&count)
	if r.Error != nil {
		return 0, r.Error
	}
	return count, nil
}
