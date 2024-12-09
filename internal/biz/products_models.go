package biz


type Product struct {
	Id int64
	Name string
	Code string
	Description string 
}
type ProductFilter struct {
	Name string
	Code string 
}
type ListProductsFilter struct {
	Page int64
	PageSize int64
	Filters  []ProductFilter 
}

