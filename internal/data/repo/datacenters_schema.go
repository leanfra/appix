package repo

import "context"

const DatacenterTable = "datacenters"

type Datacenter struct {
	ID          uint32 `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"type:varchar(255);index:idx_dc_name,unique"`
	Description string `gorm:"type:varchar(255);"`
}

type DatacentersFilter struct {
	Page     uint32
	PageSize uint32
	Ids      []uint32
	Names    []string
}

func (f *DatacentersFilter) GetIds() []uint32 {
	return f.Ids
}

type DatacentersRepo interface {
	RequireCounter
	CreateDatacenters(ctx context.Context, tx TX, dcs []*Datacenter) error
	UpdateDatacenters(ctx context.Context, tx TX, dcs []*Datacenter) error
	DeleteDatacenters(ctx context.Context, tx TX, ids []uint32) error
	GetDatacenters(ctx context.Context, id uint32) (*Datacenter, error)
	ListDatacenters(ctx context.Context, tx TX, filter *DatacentersFilter) ([]*Datacenter, error)
	CountDatacenters(ctx context.Context, tx TX, filter CountFilter) (int64, error)
}
