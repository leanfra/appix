package repo

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
