package data

import (
	"appix/internal/biz"
)

type Env struct {
	ID          uint32 `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"type:varchar(255);index:idx_name,unique"`
	Description string `gorm:"type:varchar(255);"`
}

func NewEnv(t *biz.Env) (*Env, error) {
	if t == nil {
		return nil, nil
	}
	return &Env{
		ID:          t.Id,
		Name:        t.Name,
		Description: t.Description,
	}, nil
}

func NewEnvs(es []*biz.Env) ([]*Env, error) {
	var envs = make([]*Env, len(es))
	for i, f := range es {
		nf, err := NewEnv(f)
		if err != nil {
			return nil, err
		}
		envs[i] = nf
	}
	return envs, nil
}

func NewBizEnv(t *Env) (*biz.Env, error) {
	return &biz.Env{
		Id:          t.ID,
		Name:        t.Name,
		Description: t.Description,
	}, nil
}

func NewBizEnvs(es []*Env) ([]*biz.Env, error) {
	var biz_envs = make([]*biz.Env, len(es))
	for i, f := range es {
		biz_envs[i] = &biz.Env{
			Id:          f.ID,
			Name:        f.Name,
			Description: f.Description,
		}
	}
	return biz_envs, nil
}
