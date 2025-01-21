package repo

const HostgroupFeatureTable = "hostgroup_features"

type HostgroupFeature struct {
	Id          uint32 `gorm:"primaryKey;autoIncrement"`
	HostgroupID uint32 `gorm:"index:idx_hostgroup_id_feature_id,unique"`
	FeatureID   uint32 `gorm:"index:idx_hostgroup_id_feature_id,unique"`
}

type HostgroupFeaturesFilter struct {
	Ids          []uint32
	HostgroupIds []uint32
	FeatureIds   []uint32
	KVs          []string
	Page         uint32
	PageSize     uint32
}

type HostgroupMatchFeaturesFilter struct {
	FeatureIds []uint32
}
