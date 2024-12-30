package sqldb

import (
	"appix/internal/data/repo"
	"context"
	"fmt"

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
func (d *EnvsRepoGorm) DeleteEnvs(
	ctx context.Context,
	tx repo.TX,
	ids []uint32) error {

	r := d.data.WithTX(tx).WithContext(ctx).Where("id in (?)", ids).Delete(&repo.Env{})
	if r.Error != nil {
		return r.Error
	}
	if r.RowsAffected != int64(len(ids)) {
		return fmt.Errorf("delete failed. rows affected not equal wanted. affected %d. want %d",
			r.RowsAffected, len(ids))
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
			params := make([]interface{}, len(filter.Names))
			for i, v := range filter.Names {
				params[i] = "%" + v + "%"
			}
			query = query.Where(nameConditions, params...)
		}
	}
	r := query.Find(&envs)
	if r.Error != nil {
		return nil, r.Error
	}

	return envs, nil
}

func (d *EnvsRepoGorm) CountEnvs(ctx context.Context,
	tx repo.TX,
	filter repo.CountFilter) (int64, error) {

	query := d.data.WithTX(tx).WithContext(ctx)
	var count int64
	if filter != nil {
		if len(filter.GetIds()) > 0 {
			query = query.Where("id in (?)", filter.GetIds())
		}
	}
	r := query.Model(&repo.Env{}).Count(&count)
	if r.Error != nil {
		return 0, r.Error
	}
	return count, nil
}

func (d *EnvsRepoGorm) CountRequire(ctx context.Context,
	tx repo.TX,
	need repo.RequireType,
	ids []uint32) (int64, error) {

	if len(ids) == 0 {
		return 0, repo.ErrorRequireIds
	}

	// require nothing
	return 0, nil

}
