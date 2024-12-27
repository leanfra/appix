package biz

type Hostgroup struct {
	Id              uint32
	Name            string
	Description     string
	ClusterId       uint32
	DatacenterId    uint32
	EnvId           uint32
	ProductId       uint32
	TeamId          uint32
	FeaturesId      []uint32
	TagsId          []uint32
	ShareProductsId []uint32
	ShareTeamsId    []uint32
}

type ListHostgroupsFilter struct {
	Page            uint32
	PageSize        uint32
	Ids             []uint32
	Names           []string
	ClustersId      []uint32
	DatacentersId   []uint32
	EnvsId          []uint32
	ProductsId      []uint32
	TeamsId         []uint32
	FeaturesId      []uint32
	TagsId          []uint32
	ShareProductsId []uint32
	ShareTeamsId    []uint32
}
