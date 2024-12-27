package repo

const FeatureTable = "features"

type Feature struct {
	Id    uint32 `gorm:"primaryKey;autoIncrement"`
	Name  string `gorm:"type:varchar(255);index:idx_feature_name_value,unique"`
	Value string `gorm:"type:varchar(255);index:idx_feature_name_value"`
}

type FeaturesFilter struct {
	Page     uint32
	PageSize uint32
	Ids      []uint32
	Names    []string
	Kvs      []string
}

func (f *FeaturesFilter) GetIds() []uint32 {
	return f.Ids
}
