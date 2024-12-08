package data

import (
	"appix/internal/biz"
	"context"

	"github.com/go-kratos/kratos/v2/log"
	//  TODO: modify project name
	// biz "appix/internal/biz"
)

type TeamsRepoImpl struct {
	data *Data
	log  *log.Helper
}

func NewTeamsRepoImpl(data *Data, logger log.Logger) (*TeamsRepoImpl, error) {

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
func (d *TeamsRepoImpl) CreateTeams(ctx context.Context, teams []biz.Team) error {
	// TODO database operations
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
func (d *TeamsRepoImpl) UpdateTeams(ctx context.Context, teams []biz.Team) error {
	// TODO database operations
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
func (d *TeamsRepoImpl) DeleteTeams(ctx context.Context, ids []int64) error {
	// TODO database operations
	r := d.data.db.WithContext(ctx).Where("id in (?)", ids).Delete(&Team{})
	if r.Error != nil {
		return r.Error
	}

	return nil
}

// GetTeams is
func (d *TeamsRepoImpl) GetTeams(ctx context.Context, id int64) (*biz.Team, error) {
	// TODO database operations
	team := &Team{}
	r := d.data.db.WithContext(ctx).First(team, id)
	if r.Error != nil {
		return nil, r.Error
	}

	return nil, nil
}

// ListTeams is
func (d *TeamsRepoImpl) ListTeams(ctx context.Context,
	filter *biz.ListTeamsFilter) ([]biz.Team, error) {

	// TODO database operations
	db_teams := []Team{}
	query := d.data.db.WithContext(ctx)
	if filter != nil {
		var offset int
		if filter.Page > 0 && filter.PageSize > 0 {
			offset = int((filter.Page - 1) * filter.PageSize)
			query = query.Offset(offset).Limit(int(filter.PageSize))
		}

		orConditions := make([]interface{}, len(filter.Filters))
		for i, pair := range filter.Filters {

			andConditions := map[string]interface{}{}
			if pair.Code != "" {
				andConditions["code"] = pair.Code
			}
			if pair.Leader != "" {
				andConditions["leader"] = pair.Leader
			}
			if pair.Name != "" {
				andConditions["name"] = pair.Name
			}
			orConditions[i] = andConditions
		}
		query = query.Where(orConditions)
	}
	r := query.Find(&db_teams)
	if r.Error != nil {
		return nil, r.Error
	}

	return nil, nil
}
