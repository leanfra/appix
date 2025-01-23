package sqldb

import (
	"opspillar/internal/data/repo"
	"context"
	"fmt"

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

	if err := initTable(data.DB, &repo.Team{}, repo.TeamTable); err != nil {
		return nil, err
	}

	return &TeamsRepoGorm{
		data: data,
		log:  log.NewHelper(logger),
	}, nil
}

// CreateTeams is
func (d *TeamsRepoGorm) CreateTeams(ctx context.Context, tx repo.TX, teams []*repo.Team) error {
	r := d.data.WithTX(tx).WithContext(ctx).Create(teams)
	if r.Error != nil {
		return r.Error
	}

	return nil
}

// UpdateTeams is
func (d *TeamsRepoGorm) UpdateTeams(ctx context.Context, tx repo.TX, teams []*repo.Team) error {
	r := d.data.WithTX(tx).WithContext(ctx).Save(teams)
	if r.Error != nil {
		return r.Error
	}

	return nil
}

// DeleteTeams is delete items by id
// return error if affected rows not equal to wanted
func (d *TeamsRepoGorm) DeleteTeams(ctx context.Context,
	tx repo.TX,
	ids []uint32) error {

	r := d.data.WithTX(tx).WithContext(ctx).Where("id in (?)", ids).Delete(&repo.Team{})
	if r.Error != nil {
		return r.Error
	}
	if r.RowsAffected != int64(len(ids)) {
		return fmt.Errorf("delete not equal expected. want %d. affected %d", len(ids), r.RowsAffected)
	}

	return nil
}

// GetTeams is
// notfound return error
func (d *TeamsRepoGorm) GetTeams(ctx context.Context, id uint32) (*repo.Team, error) {
	team := &repo.Team{}
	r := d.data.DB.WithContext(ctx).First(team, id)
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
			params := make([]interface{}, len(filter.Codes))
			for i, v := range filter.Codes {
				params[i] = "%" + v + "%"
			}
			query = query.Where(codeConditions, params...)
		}
		if len(filter.LeadersId) > 0 {
			query = query.Where("leader_id in (?)", filter.LeadersId)
		}
		if len(filter.Names) > 0 {
			nameConditions := buildOrLike("name", len(filter.Names))
			params := make([]interface{}, len(filter.Names))
			for i, v := range filter.Names {
				params[i] = "%" + v + "%"
			}
			query = query.Where(nameConditions, params...)
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

func (d *TeamsRepoGorm) CountRequire(ctx context.Context,
	tx repo.TX,
	need repo.RequireType,
	ids []uint32) (int64, error) {

	if len(ids) == 0 {
		return 0, repo.ErrorRequireIds
	}

	var condition string
	switch need {
	case repo.RequireUser:
		condition = "leader_id in (?)"
	default:
		return 0, repo.ErrorRequireIds
	}

	var count int64
	r := d.data.WithTX(tx).WithContext(ctx).Model(&repo.Team{}).
		Where(condition, ids).Count(&count)
	if r.Error != nil {
		return 0, r.Error
	}

	return count, nil

}
