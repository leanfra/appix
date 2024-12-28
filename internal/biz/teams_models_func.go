package biz

import (
	"appix/internal/data/repo"
	"fmt"
)

func (f *Team) Validate(isNew bool) error {
	if len(f.Name) == 0 || len(f.Code) == 0 || len(f.Leader) == 0 {
		return fmt.Errorf("InvalidNameCodeLeader")
	}
	if !isNew {
		if f.Id <= 0 {
			return fmt.Errorf("InvalidId")
		}
	}
	return nil
}

func (lf *ListTeamsFilter) Validate() error {
	if lf == nil {
		return nil
	}
	if len(lf.Codes) > MaxFilterValues ||
		len(lf.Ids) > MaxFilterValues ||
		len(lf.Leaders) > MaxFilterValues ||
		len(lf.Names) > MaxFilterValues {
		return ErrFilterValuesExceedMax
	}
	if lf.PageSize == 0 || lf.PageSize > MaxPageSize {
		return ErrFilterInvalidPagesize
	}
	return nil
}

func DefaultTeamsFilter() *ListTeamsFilter {
	return &ListTeamsFilter{
		Page:     1,
		PageSize: DefaultPageSize,
	}
}

func ToDBTeam(t *Team) (*repo.Team, error) {
	return &repo.Team{
		ID:          t.Id,
		Name:        t.Name,
		Code:        t.Code,
		Leader:      t.Leader,
		Description: t.Description,
	}, nil
}

func ToDBTeams(ts []*Team) ([]*repo.Team, error) {
	var teams = make([]*repo.Team, len(ts))
	for i, t := range ts {
		nt, err := ToDBTeam(t)
		if err != nil {
			return nil, err
		}
		teams[i] = nt
	}
	return teams, nil
}

func ToBizTeam(t *repo.Team) (*Team, error) {
	return &Team{
		Id:          t.ID,
		Code:        t.Code,
		Description: t.Description,
		Leader:      t.Leader,
		Name:        t.Name,
	}, nil
}

func ToBizTeams(teams []*repo.Team) ([]*Team, error) {
	var biz_teams = make([]*Team, len(teams))
	for i, t := range teams {
		biz_teams[i] = &Team{
			Id:          t.ID,
			Code:        t.Code,
			Description: t.Description,
			Leader:      t.Leader,
			Name:        t.Name,
		}
	}
	return biz_teams, nil
}

func ToDBTeamsFilter(filter *ListTeamsFilter) *repo.TeamsFilter {
	return &repo.TeamsFilter{
		Codes:    filter.Codes,
		Ids:      filter.Ids,
		Leaders:  filter.Leaders,
		Names:    filter.Names,
		Page:     filter.Page,
		PageSize: filter.PageSize,
	}
}
