package data

import (
	"appix/internal/biz"
)

type Env struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"type:varchar(255);index:idx_name,unique"`
	Description string `gorm:"type:varchar(255);"`
}

func NewEnv(t biz.Env) (*Env, error) {
	return &Env{
		ID:          uint(t.Id),
		Name:        t.Name,
		Description: t.Description,
	}, nil
}

func NewEnvs(es []biz.Env) ([]*Env, error) {
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
		Id:          int64(t.ID),
		Name:        t.Name,
		Description: t.Description,
	}, nil
}

func NewBizEnvs(es []Env) ([]biz.Env, error) {
	var biz_envs = make([]biz.Env, len(es))
	for i, f := range es {
		biz_envs[i] = biz.Env{
			Id:          int64(f.ID),
			Name:        f.Name,
			Description: f.Description,
		}
	}
	return biz_envs, nil
}
