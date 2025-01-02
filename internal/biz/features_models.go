package biz

type Feature struct {
	Id          uint32
	Name        string
	Value       string
	Description string
}

type ListFeaturesFilter struct {
	Page     uint32
	PageSize uint32
	Ids      []uint32
	Names    []string
	Kvs      []string
}
