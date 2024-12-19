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
	Page        uint32
	PageSize    uint32
	Names       []string
	Clusters    []string
	Datacenters []string
	Envs        []string
	Products    []string
	Teams       []string
	Features    []string
	Tags        []string
}
