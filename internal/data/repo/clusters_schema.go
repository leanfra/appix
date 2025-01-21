package repo

import "context"

const ClusterTable = "clusters"

type Cluster struct {
	ID          uint32 `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"type:varchar(255);index:idx_cluster_name,unique"`
	Description string `gorm:"type:varchar(255);"`
}

type ClustersFilter struct {
	Page     uint32
	PageSize uint32
	Names    []string
	Ids      []uint32
}

func (f *ClustersFilter) GetIds() []uint32 {
	return f.Ids
}

type ClustersRepo interface {
	RequireCounter
	CreateClusters(ctx context.Context, tx TX, cs []*Cluster) error
	UpdateClusters(ctx context.Context, tx TX, cs []*Cluster) error
	DeleteClusters(ctx context.Context, tx TX, ids []uint32) error
	GetClusters(ctx context.Context, id uint32) (*Cluster, error)
	ListClusters(ctx context.Context, tx TX, filter *ClustersFilter) ([]*Cluster, error)
	CountClusters(ctx context.Context, tx TX, filter CountFilter) (int64, error)
}
