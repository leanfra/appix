package biz



type Datacenter struct {
	Id int64
	Name string
	Description string 
}

type DatacenterFilter struct {
	Name string 
}

type ListDatacentersFilter struct {
	Page int64
	PageSize int64
	Filters  []DatacenterFilter 
}

