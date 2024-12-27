package sqldb

import (
	"appix/internal/data/repo"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type EnvsRepoGorm struct {
	data *DataGorm
	log  *log.Helper
}

func NewEnvsRepoGorm(data *DataGorm, logger log.Logger) (repo.EnvsRepo, error) {

	if err := validateData(data); err != nil {
		return nil, err
	}
	if err := initTable(data.DB, &repo.Env{}, repo.EnvTable); err != nil {
		return nil, err
	}
	return &EnvsRepoGorm{
		data: data,
		log:  log.NewHelper(logger),
	}, nil
}

// CreateEnvs is
func (d *EnvsRepoGorm) CreateEnvs(ctx context.Context, envs []*repo.Env) error {

	r := d.data.DB.WithContext(ctx).Create(envs)
	if r.Error != nil {
		return r.Error
	}
	return nil
}

// UpdateEnvs is
func (d *EnvsRepoGorm) UpdateEnvs(ctx context.Context, envs []*repo.Env) error {

	r := d.data.DB.WithContext(ctx).Save(envs)
	if r.Error != nil {
		return r.Error
	}
	return nil
}

// DeleteEnvs is
func (d *EnvsRepoGorm) DeleteEnvs(ctx context.Context, ids []uint32) error {

	r := d.data.DB.WithContext(ctx).Where("id in (?)", ids).Delete(&repo.Env{})
	if r.Error != nil {
		return r.Error
	}
	return nil
}

// GetEnvs is
func (d *EnvsRepoGorm) GetEnvs(ctx context.Context, id uint32) (*repo.Env, error) {

	env := &repo.Env{}
	r := d.data.DB.WithContext(ctx).First(env, id)
	if r.Error != nil {
		return nil, r.Error
	}

	return env, nil
}

// ListEnvs is
func (d *EnvsRepoGorm) ListEnvs(ctx context.Context,
	tx repo.TX,
	filter *repo.EnvsFilter) ([]*repo.Env, error) {

	envs := []*repo.Env{}
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
		if len(filter.Names) > 0 {
			nameConditions := buildOrLike("name", len(filter.Names))
			query = query.Where(nameConditions, filter.Names)
		}
	}
	r := query.Find(&envs)
	if r.Error != nil {
		return nil, r.Error
	}

	return envs, nil
}
