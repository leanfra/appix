package repo

const HostgroupTeamTable = "hostgroup_tags"

type HostgroupTeam struct {
	Id          uint32 `gorm:"primaryKey;autoIncrement"`
	HostgroupID uint32 `gorm:"index:idx_hostgroup_id_team_id,unique"`
	TeamID      uint32 `gorm:"index:idx_hostgroup_id_team_id,unique"`
}

type HostgroupTeamsFilter struct {
	Ids          []uint32
	HostgroupIds []uint32
	TeamIds      []uint32
	KVs          []string // k:v format
	Page         uint32
	PageSize     uint32
}
