package repo

import (
	"context"
)

const HostgroupTable = "hostgroups"

type Hostgroup struct {
	ChangeInfo
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
	Page          uint32
	PageSize      uint32
	Ids           []uint32
	Names         []string
	ClustersId    []uint32
	DatacentersId []uint32
	EnvsId        []uint32
	ProductsId    []uint32
	TeamsId       []uint32
}

func (f *HostgroupsFilter) GetIds() []uint32 {
	return f.Ids
}

type HostgroupsRepo interface {
	RequireCounter
	CreateHostgroups(ctx context.Context, tx TX, hgs []*Hostgroup) error
	UpdateHostgroups(ctx context.Context, tx TX, hgs []*Hostgroup) error
	DeleteHostgroups(ctx context.Context, tx TX, ids []uint32) error
	GetHostgroups(ctx context.Context, id uint32) (*Hostgroup, error)
	ListHostgroups(ctx context.Context, tx TX, filter *HostgroupsFilter) ([]*Hostgroup, error)
	CountHostgroups(ctx context.Context, tx TX, filter CountFilter) (int64, error)
}
