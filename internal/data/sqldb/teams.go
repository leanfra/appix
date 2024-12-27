package sqldb

import (
	"appix/internal/data/repo"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type TeamsRepoGorm struct {
	data *DataGorm
	log  *log.Helper
}

func NewTeamsRepoGorm(data *DataGorm, logger log.Logger) (repo.TeamsRepo, error) {

	if err := validateData(data); err != nil {
		return nil, err
	}

	if err := initTable(data.db, &repo.Team{}, repo.TeamTable); err != nil {
		return nil, err
	}

	return &TeamsRepoGorm{
		data: data,
		log:  log.NewHelper(logger),
	}, nil
}

// CreateTeams is
func (d *TeamsRepoGorm) CreateTeams(ctx context.Context, teams []*repo.Team) error {
	r := d.data.db.WithContext(ctx).Create(teams)
	if r.Error != nil {
		return r.Error
	}

	return nil
}

// UpdateTeams is
func (d *TeamsRepoGorm) UpdateTeams(ctx context.Context, teams []*repo.Team) error {
	r := d.data.db.WithContext(ctx).Save(teams)
	if r.Error != nil {
		return r.Error
	}

	return nil
}

// DeleteTeams is
func (d *TeamsRepoGorm) DeleteTeams(ctx context.Context, ids []uint32) error {
	r := d.data.db.WithContext(ctx).Where("id in (?)", ids).Delete(&repo.Team{})
	if r.Error != nil {
		return r.Error
	}

	return nil
}

// GetTeams is
func (d *TeamsRepoGorm) GetTeams(ctx context.Context, id uint32) (*repo.Team, error) {
	team := &repo.Team{}
	r := d.data.db.WithContext(ctx).First(team, id)
	if r.Error != nil {
		return nil, r.Error
	}

	return team, nil
}

// ListTeams is
func (d *TeamsRepoGorm) ListTeams(ctx context.Context,
	tx repo.TX,
	filter *repo.TeamsFilter) ([]*repo.Team, error) {

	db_teams := []*repo.Team{}
	query := d.data.WithTX(tx).WithContext(ctx)
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

	return db_teams, nil
}

func (d *TeamsRepoGorm) CountTeams(ctx context.Context,
	tx repo.TX,
	filter repo.CountFilter) (int64, error) {

	var count int64
	query := d.data.WithTX(tx).WithContext(ctx)
	if filter != nil {
		if len(filter.GetIds()) > 0 {
			query = query.Where("id in (?)", filter.GetIds())
		}
	}
	r := query.Model(&repo.Team{}).Count(&count)
	if r.Error != nil {
		return 0, r.Error
	}
	return count, nil
}
