package biz



type Env struct {
	Id int64
	Name string
	Description string 
}

type EnvFilter struct {
	Name string 
}

type ListEnvsFilter struct {
	Page int64
	PageSize int64
	Filters  []EnvFilter 
}

