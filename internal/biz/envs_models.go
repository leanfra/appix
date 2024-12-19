package biz

type Env struct {
	Id          uint32
	Name        string
	Description string
}

type ListEnvsFilter struct {
	Page     uint32
	PageSize uint32
	Names    []string
	Ids      []uint32
}
