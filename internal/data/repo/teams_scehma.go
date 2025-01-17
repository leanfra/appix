package repo

import "context"

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

type TeamsRepo interface {
	RequireCounter
	CreateTeams(ctx context.Context, tx TX, teams []*Team) error
	UpdateTeams(ctx context.Context, tx TX, teams []*Team) error
	DeleteTeams(ctx context.Context, tx TX, ids []uint32) error
	GetTeams(ctx context.Context, id uint32) (*Team, error)
	ListTeams(ctx context.Context, tx TX, filter *TeamsFilter) ([]*Team, error)
	CountTeams(ctx context.Context, tx TX, filter CountFilter) (int64, error)
}
