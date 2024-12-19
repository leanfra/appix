package biz

type Team struct {
	Id          uint32
	Name        string
	Code        string
	Description string
	Leader      string
}

type ListTeamsFilter struct {
	Page     uint32
	PageSize uint32
	Names    []string
	Codes    []string
	Leaders  []string
	Ids      []uint32
}
