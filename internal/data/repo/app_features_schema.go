package repo

const AppFeatureTable = "app_features"

type AppFeature struct {
	Id        uint32 `gorm:"primaryKey;autoIncrement"`
	AppID     uint32 `gorm:"index:idx_app_id_feature_id,unique"`
	FeatureID uint32 `gorm:"index:idx_app_id_feature_id,unique"`
}

type AppFeaturesFilter struct {
	Ids        []uint32
	AppIds     []uint32
	FeatureIds []uint32
	KVs        []string
	Page       uint32
	PageSize   uint32
}
