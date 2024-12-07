package biz

type Feature struct {
	Id    int64
	Name  string
	Value string
}

type FeatureFilter struct {
	Name  string
	Value string
}

type ListFeaturesFilter struct {
	Page     int64
	PageSize int64
	Filters  []FeatureFilter
}
