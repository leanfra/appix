package biz

type Application struct {
	Id           uint32
	Name         string
	Description  string
	Owner        string
	IsStateful   bool
	ClusterId    uint32
	DatacenterId uint32
	EnvId        uint32
	ProductId    uint32
	TeamId       uint32
	FeaturesId   []uint32
	TagsId       []uint32
}

type ListApplicationsFilter struct {
	Page        uint32
	PageSize    uint32
	Ids         []uint32
	Names       []string
	Clusters    []string
	Datacenters []string
	Envs        []string
	Products    []string
	Teams       []string
	Features    []string
	Tags        []string
	IsStateful  string
}

const IsStatefulTrue = "true"
const IsStatefulFalse = "false"
const IsStatefulNone = ""
