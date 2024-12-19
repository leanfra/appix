package biz

type Datacenter struct {
	Id          uint32
	Name        string
	Description string
}

type ListDatacentersFilter struct {
	Page     uint32
	PageSize uint32
	Ids      []uint32
	Names    []string
}
