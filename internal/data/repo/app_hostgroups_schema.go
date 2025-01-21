package repo

const AppHostgroupTable = "app_hostgroups"

type AppHostgroup struct {
	Id          uint32 `gorm:"primaryKey;autoIncrement"`
	AppID       uint32 `gorm:"index:idx_app_id_hostgroup_id,unique"`
	HostgroupID uint32 `gorm:"index:idx_app_id_hostgroup_id,unique"`
}

type AppHostgroupsFilter struct {
	Ids          []uint32
	AppIds       []uint32
	HostgroupIds []uint32
	KVs          []string // k:v format
	Page         uint32
	PageSize     uint32
}
