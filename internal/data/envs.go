package data

import (
	"appix/internal/biz"
	"context"

	"github.com/go-kratos/kratos/v2/log"
	//  TODO: modify project name
	// biz "appix/internal/biz"
)

type EnvsRepoImpl struct {
	data *Data
	log  *log.Helper
}

func NewEnvsRepoImpl(data *Data, logger log.Logger) (biz.EnvsRepo, error) {

	if err := validateData(data); err != nil {
		return nil, err
	}
	if err := initTable(data.db, &Env{}, "env"); err != nil {
		return nil, err
	}
	return &EnvsRepoImpl{
		data: data,
		log:  log.NewHelper(logger),
	}, nil
}

// CreateEnvs is
func (d *EnvsRepoImpl) CreateEnvs(ctx context.Context, envs []biz.Env) error {

	db_env, err := NewEnvs(envs)
	if err != nil {
		return err
	}
	r := d.data.db.WithContext(ctx).Create(db_env)
	if r.Error != nil {
		return r.Error
	}
	return nil
}

// UpdateEnvs is
func (d *EnvsRepoImpl) UpdateEnvs(ctx context.Context, envs []biz.Env) error {

	db_envs, err := NewEnvs(envs)
	if err != nil {
		return err
	}
	r := d.data.db.WithContext(ctx).Save(db_envs)
	if r.Error != nil {
		return r.Error
	}
	return nil
}

// DeleteEnvs is
func (d *EnvsRepoImpl) DeleteEnvs(ctx context.Context, ids []int64) error {

	r := d.data.db.WithContext(ctx).Where("id in (?)", ids).Delete(&Env{})
	if r.Error != nil {
		return r.Error
	}
	return nil
}

// GetEnvs is
func (d *EnvsRepoImpl) GetEnvs(ctx context.Context, id int64) (*biz.Env, error) {

	env := &Env{}
	r := d.data.db.WithContext(ctx).First(env, id)
	if r.Error != nil {
		return nil, r.Error
	}

	return NewBizEnv(env)
}

// ListEnvs is
func (d *EnvsRepoImpl) ListEnvs(ctx context.Context, filter *biz.ListEnvsFilter) ([]biz.Env, error) {

	envs := []Env{}
	query := d.data.db.WithContext(ctx)
	if filter != nil {
		var offset int
		if filter.Page > 0 && filter.PageSize > 0 {
			offset = int((filter.Page - 1) * filter.PageSize)
			query = query.Offset(offset).Limit(int(filter.PageSize))
		}
		for _, pair := range filter.Filters {
			query = query.Where("name =?", pair.Name)
		}
	}
	r := query.Find(&envs)
	if r.Error != nil {
		return nil, r.Error
	}

	return NewBizEnvs(envs)
}
