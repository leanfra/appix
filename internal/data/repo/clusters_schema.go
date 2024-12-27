package repo

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
