package data

import (
	"appix/internal/biz"
)

type Team struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"type:varchar(255);index:idx_name,unique"`
	Code        string `gorm:"type:varchar(255);index:idx_code,unique"`
	Leader      string `gorm:"type:varchar(255);index:idx_leader"`
	Description string `gorm:"type:varchar(255);"`
}

func NewTeam(t biz.Team) (*Team, error) {
	return &Team{
		ID:          uint(t.Id),
		Name:        t.Name,
		Code:        t.Code,
		Leader:      t.Leader,
		Description: t.Description,
	}, nil
}

func NewTeams(ts []biz.Team) ([]*Team, error) {
	var teams = make([]*Team, len(ts))
	for i, t := range ts {
		nt, err := NewTeam(t)
		if err != nil {
			return nil, err
		}
		teams[i] = nt
	}
	return teams, nil
}

func NewBizTeam(t *Team) (*biz.Team, error) {
	return &biz.Team{
		Id:          int64(t.ID),
		Code:        t.Code,
		Description: t.Description,
		Leader:      t.Leader,
		Name:        t.Name,
	}, nil
}

func NewBizTeams(teams []Team) ([]biz.Team, error) {
	var biz_teams = make([]biz.Team, len(teams))
	for i, t := range teams {
		biz_teams[i] = biz.Team{
			Id:          int64(t.ID),
			Code:        t.Code,
			Description: t.Description,
			Leader:      t.Leader,
			Name:        t.Name,
		}
	}
	return biz_teams, nil
}
