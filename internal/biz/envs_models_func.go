package biz

import (
	"appix/internal/data/repo"
	"fmt"
)

func (f *Env) Validate(isNew bool) error {
	if len(f.Name) == 0 {
		return fmt.Errorf("InvalidNameValue")
	}
	if !isNew {
		if f.Id <= 0 {
			return fmt.Errorf("InvalidId")
		}
	}
	return nil
}

func (lf *ListEnvsFilter) Validate() error {
	if lf == nil {
		return nil
	}
	if len(lf.Names) > MaxFilterValues || len(lf.Ids) > MaxFilterValues {
		return ErrFilterValuesExceedMax
	}
	if lf.PageSize == 0 || lf.PageSize > MaxPageSize {
		return ErrFilterInvalidPagesize
	}
	return nil
}

func DefaultEnvFilter() *ListEnvsFilter {
	return &ListEnvsFilter{
		Page:     1,
		PageSize: DefaultPageSize,
	}
}

func ToDBEnv(t *Env) (*repo.Env, error) {
	if t == nil {
		return nil, nil
	}
	return &repo.Env{
		ID:          t.Id,
		Name:        t.Name,
		Description: t.Description,
	}, nil
}

func ToDBEnvs(es []*Env) ([]*repo.Env, error) {
	var envs = make([]*repo.Env, len(es))
	for i, f := range es {
		nf, err := ToDBEnv(f)
		if err != nil {
			return nil, err
		}
		envs[i] = nf
	}
	return envs, nil
}

func ToBizEnv(t *repo.Env) (*Env, error) {
	return &Env{
		Id:          t.ID,
		Name:        t.Name,
		Description: t.Description,
	}, nil
}

func ToBizEnvs(es []*repo.Env) ([]*Env, error) {
	var biz_envs = make([]*Env, len(es))
	for i, f := range es {
		biz_envs[i] = &Env{
			Id:          f.ID,
			Name:        f.Name,
			Description: f.Description,
		}
	}
	return biz_envs, nil
}

func ToDBEnvsFilter(filter *ListEnvsFilter) *repo.EnvsFilter {
	if filter == nil {
		return nil
	}
	return &repo.EnvsFilter{
		Ids:      filter.Ids,
		Names:    filter.Names,
		Page:     filter.Page,
		PageSize: filter.PageSize,
	}
}
