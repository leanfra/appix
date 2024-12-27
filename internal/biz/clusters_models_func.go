package biz

import (
	"appix/internal/data/repo"
	"fmt"
)

func (f *Cluster) Validate(isNew bool) error {
	if len(f.Name) == 0 {
		return fmt.Errorf("InvalidNameValue")
	}
	if !isNew {
		if f.Id <= 0 {
			return fmt.Errorf("InvalidId")
		}
	}
	return nil
}

func (lf *ListClustersFilter) Validate() error {
	if lf == nil {
		return nil
	}
	if len(lf.Ids) > MaxFilterValues || len(lf.Names) > MaxFilterValues {
		return ErrFilterValuesExceedMax
	}
	if lf.PageSize == 0 || lf.PageSize > MaxPageSize {
		return ErrFilterInvalidPagesize
	}
	return nil
}

func DefaultClusterFilter() *ListClustersFilter {
	return &ListClustersFilter{
		Page:     1,
		PageSize: DefaultPageSize,
	}
}

func NewCluster(t *Cluster) (*repo.Cluster, error) {
	if t == nil {
		return nil, nil
	}
	return &repo.Cluster{
		ID:          t.Id,
		Name:        t.Name,
		Description: t.Description,
	}, nil
}

func NewClusters(es []*Cluster) ([]*repo.Cluster, error) {
	var clusters = make([]*repo.Cluster, len(es))
	for i, f := range es {
		nf, err := NewCluster(f)
		if err != nil {
			return nil, err
		}
		clusters[i] = nf
	}
	return clusters, nil
}

func NewBizCluster(t *repo.Cluster) (*Cluster, error) {
	return &Cluster{
		Id:          t.ID,
		Name:        t.Name,
		Description: t.Description,
	}, nil
}

func NewBizClusters(es []*repo.Cluster) ([]*Cluster, error) {
	var biz_clusters = make([]*Cluster, len(es))
	for i, f := range es {
		biz_cluster, err := NewBizCluster(f)
		if err != nil {
			return nil, err
		}

		biz_clusters[i] = biz_cluster
	}
	return biz_clusters, nil
}

func NewClustersFilter(filter *ListClustersFilter) *repo.ClustersFilter {
	return &repo.ClustersFilter{
		Ids:      filter.Ids,
		Names:    filter.Names,
		Page:     filter.Page,
		PageSize: filter.PageSize,
	}
}
