package repo

const HostgroupTagTable = "hostgroup_tags"

type HostgroupTag struct {
	Id          uint32 `gorm:"primaryKey;autoIncrement"`
	HostgroupID uint32 `gorm:"index:idx_hostgroup_id_tag_id,unique"`
	TagID       uint32 `gorm:"index:idx_hostgroup_id_tag_id,unique"`
}

type HostgroupTagsFilter struct {
	Ids          []uint32
	HostgroupIds []uint32
	TagIds       []uint32
	KVs          []string // k:v format
	Page         uint32
	PageSize     uint32
}
