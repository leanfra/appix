package biz

type Application struct {
	Id           uint32
	Name         string
	Description  string
	Owner        string
	IsStateful   bool
	ProductId    uint32
	TeamId       uint32
	FeaturesId   []uint32
	TagsId       []uint32
	HostgroupsId []uint32
}

type ListApplicationsFilter struct {
	Page         uint32
	PageSize     uint32
	Ids          []uint32
	Names        []string
	IsStateful   string
	ProductsId   []uint32
	TeamsId      []uint32
	FeaturesId   []uint32
	TagsId       []uint32
	HostgroupsId []uint32
}

type MatchAppHostgroupsFilter struct {
	FeaturesId []uint32
	ProductId  uint32
	TeamId     uint32
}
