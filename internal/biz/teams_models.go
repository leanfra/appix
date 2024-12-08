package biz


type Team struct {
	Id int64
	Name string
	Code string
	Description string
	Leader string 
}
type TeamFilter struct {
	Name string
	Code string
	Leader string 
}
type ListTeamsFilter struct {
	Page int64
	PageSize int64
	Filters  []TeamFilter 
}

