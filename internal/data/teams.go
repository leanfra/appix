package data

import (
	"appix/internal/biz"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type TeamsRepoImpl struct {
	data *Data
	log  *log.Helper
}

func NewTeamsRepoImpl(data *Data, logger log.Logger) (biz.TeamsRepo, error) {

	if err := validateData(data); err != nil {
		return nil, err
	}

	if err := initTable(data.db, &Team{}, "team"); err != nil {
		return nil, err
	}

	return &TeamsRepoImpl{
		data: data,
		log:  log.NewHelper(logger),
	}, nil
}

// CreateTeams is
func (d *TeamsRepoImpl) CreateTeams(ctx context.Context, teams []*biz.Team) error {
	db_teams, err := NewTeams(teams)
	if err != nil {
		return err
	}
	r := d.data.db.WithContext(ctx).Create(db_teams)
	if r.Error != nil {
		return r.Error
	}

	return nil
}

// UpdateTeams is
func (d *TeamsRepoImpl) UpdateTeams(ctx context.Context, teams []*biz.Team) error {
	db_teams, err := NewTeams(teams)
	if err != nil {
		return err
	}
	r := d.data.db.WithContext(ctx).Save(db_teams)
	if r.Error != nil {
		return r.Error
	}

	return nil
}

// DeleteTeams is
func (d *TeamsRepoImpl) DeleteTeams(ctx context.Context, ids []uint32) error {
	r := d.data.db.WithContext(ctx).Where("id in (?)", ids).Delete(&Team{})
	if r.Error != nil {
		return r.Error
	}

	return nil
}

// GetTeams is
func (d *TeamsRepoImpl) GetTeams(ctx context.Context, id uint32) (*biz.Team, error) {
	team := &Team{}
	r := d.data.db.WithContext(ctx).First(team, id)
	if r.Error != nil {
		return nil, r.Error
	}

	return NewBizTeam(team)
}

// ListTeams is
func (d *TeamsRepoImpl) ListTeams(ctx context.Context,
	filter *biz.ListTeamsFilter) ([]*biz.Team, error) {

	db_teams := []*Team{}
	query := d.data.db.WithContext(ctx)
	if filter != nil {
		var offset int
		if filter.Page > 0 && filter.PageSize > 0 {
			offset = int((filter.Page - 1) * filter.PageSize)
			query = query.Offset(offset).Limit(int(filter.PageSize))
		}

		if len(filter.Ids) > 0 {
			query = query.Where("id in (?)", filter.Ids)
		}
		if len(filter.Codes) > 0 {
			codeConditions := buildOrLike("code", len(filter.Codes))
			query = query.Where(codeConditions, filter.Codes)
		}
		if len(filter.Leaders) > 0 {
			leaderConditions := buildOrLike("leader", len(filter.Leaders))
			query = query.Where(leaderConditions, filter.Leaders)
		}
		if len(filter.Names) > 0 {
			nameConditions := buildOrLike("name", len(filter.Names))
			query = query.Where(nameConditions, filter.Names)
		}
	}
	r := query.Find(&db_teams)
	if r.Error != nil {
		return nil, r.Error
	}

	return NewBizTeams(db_teams)
}
