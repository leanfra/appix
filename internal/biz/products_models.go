package biz

type Product struct {
	Id          uint32
	Name        string
	Code        string
	Description string
}

type ListProductsFilter struct {
	Page     uint32
	PageSize uint32
	Names    []string
	Codes    []string
	Ids      []uint32
}
