package biz

type Cluster struct {
	Id          uint32
	Name        string
	Description string
}

type ListClustersFilter struct {
	Page     uint32
	PageSize uint32
	Names    []string
	Ids      []uint32
}
