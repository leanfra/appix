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
	if e := ValidateName(f.Name); e != nil {
		return e
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

func ToDBCluster(t *Cluster) (*repo.Cluster, error) {
	if t == nil {
		return nil, nil
	}
	return &repo.Cluster{
		ID:          t.Id,
		Name:        t.Name,
		Description: t.Description,
	}, nil
}

func ToDBClusters(es []*Cluster) ([]*repo.Cluster, error) {
	var clusters = make([]*repo.Cluster, len(es))
	for i, f := range es {
		nf, err := ToDBCluster(f)
		if err != nil {
			return nil, err
		}
		clusters[i] = nf
	}
	return clusters, nil
}

func ToBizCluster(t *repo.Cluster) (*Cluster, error) {
	return &Cluster{
		Id:          t.ID,
		Name:        t.Name,
		Description: t.Description,
	}, nil
}

func ToBizClusters(es []*repo.Cluster) ([]*Cluster, error) {
	var biz_clusters = make([]*Cluster, len(es))
	for i, f := range es {
		biz_cluster, err := ToBizCluster(f)
		if err != nil {
			return nil, err
		}

		biz_clusters[i] = biz_cluster
	}
	return biz_clusters, nil
}

func ToDBClustersFilter(filter *ListClustersFilter) *repo.ClustersFilter {
	return &repo.ClustersFilter{
		Ids:      filter.Ids,
		Names:    filter.Names,
		Page:     filter.Page,
		PageSize: filter.PageSize,
	}
}
