package biz

type Tag struct {
	Id    int64
	Key   string
	Value string
}

type TagFilter struct {
	Key   string
	Value string
}

type ListTagsFilter struct {
	Page     int64
	PageSize int64
	Filters  []TagFilter
}
