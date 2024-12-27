package repo

const HostgroupTable = "hostgroups"

type Hostgroup struct {
	Id           uint32 `gorm:"primaryKey;autoIncrement"`
	Name         string `gorm:"type:varchar(255);index:idx_hg_name,unique"`
	Description  string `gorm:"type:varchar(255);"`
	ClusterId    uint32
	DatacenterId uint32
	EnvId        uint32
	ProductId    uint32
	TeamId       uint32
}

type HostgroupsFilter struct {
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

func (f *HostgroupsFilter) GetIds() []uint32 {
	return f.Ids
}
