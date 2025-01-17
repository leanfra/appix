package biz

type Team struct {
	Id          uint32
	Name        string
	Code        string
	Description string
	LeaderId    uint32
}

type ListTeamsFilter struct {
	Page      uint32
	PageSize  uint32
	Names     []string
	Codes     []string
	LeadersId []uint32
	Ids       []uint32
}
