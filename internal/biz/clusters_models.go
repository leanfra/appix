package biz



type Cluster struct {
	Id int64
	Name string
	Description string 
}

type ClusterFilter struct {
	Name string 
}

type ListClustersFilter struct {
	Page int64
	PageSize int64
	Filters  []ClusterFilter 
}

