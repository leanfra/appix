package biz

type Tag struct {
	Id          uint32
	Key         string
	Value       string
	Description string
}

type ListTagsFilter struct {
	Page     uint32
	PageSize uint32
	Keys     []string
	Kvs      []string
	Ids      []uint32
}
