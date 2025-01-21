package repo

const HostgroupProductTable = "hostgroup_products"

type HostgroupProduct struct {
	Id          uint32 `gorm:"primaryKey;autoIncrement"`
	HostgroupID uint32 `gorm:"index:idx_hostgroup_id_product_id,unique"`
	ProductID   uint32 `gorm:"index:idx_hostgroup_id_product_id,unique"`
}

type HostgroupProductsFilter struct {
	Ids          []uint32
	HostgroupIds []uint32
	ProductIds   []uint32
	KVs          []string // k:v format
	Page         uint32
	PageSize     uint32
}
