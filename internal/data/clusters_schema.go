package data

import (
	"appix/internal/biz"
)

type Cluster struct {
	ID          uint32 `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"type:varchar(255);index:idx_cluster_name,unique"`
	Description string `gorm:"type:varchar(255);"`
}

func NewCluster(t *biz.Cluster) (*Cluster, error) {
	if t == nil {
		return nil, nil
	}
	return &Cluster{
		ID:          t.Id,
		Name:        t.Name,
		Description: t.Description,
	}, nil
}

func NewClusters(es []*biz.Cluster) ([]*Cluster, error) {
	var clusters = make([]*Cluster, len(es))
	for i, f := range es {
		nf, err := NewCluster(f)
		if err != nil {
			return nil, err
		}
		clusters[i] = nf
	}
	return clusters, nil
}

func NewBizCluster(t *Cluster) (*biz.Cluster, error) {
	return &biz.Cluster{
		Id:          t.ID,
		Name:        t.Name,
		Description: t.Description,
	}, nil
}

func NewBizClusters(es []Cluster) ([]*biz.Cluster, error) {
	var biz_clusters = make([]*biz.Cluster, len(es))
	for i, f := range es {
		biz_cluster, err := NewBizCluster(&f)
		if err != nil {
			return nil, err
		}

		biz_clusters[i] = biz_cluster
	}
	return biz_clusters, nil
}
