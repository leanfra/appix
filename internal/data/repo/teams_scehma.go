package repo

const TeamTable = "teams"

type Team struct {
	ID          uint32 `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"type:varchar(255);index:idx_team_name,unique"`
	Code        string `gorm:"type:varchar(255);index:idx_team_code,unique"`
	Leader      string `gorm:"type:varchar(255);index:idx_team_leader"`
	Description string `gorm:"type:varchar(255);"`
}

type TeamsFilter struct {
	Page     uint32
	PageSize uint32
	Ids      []uint32
	Names    []string
	Leaders  []string
	Codes    []string
}

func (f *TeamsFilter) GetIds() []uint32 {
	return f.Ids
}
