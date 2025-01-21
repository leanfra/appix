package repo

import "context"

const EnvTable = "envs"

type Env struct {
	ID          uint32 `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"type:varchar(255);index:idx_env_name,unique"`
	Description string `gorm:"type:varchar(255);"`
}

type EnvsFilter struct {
	Page     uint32
	PageSize uint32
	Names    []string
	Ids      []uint32
}

func (f *EnvsFilter) GetIds() []uint32 {
	return f.Ids
}

type EnvsRepo interface {
	RequireCounter
	CreateEnvs(ctx context.Context, tx TX, envs []*Env) error
	UpdateEnvs(ctx context.Context, tx TX, envs []*Env) error
	DeleteEnvs(ctx context.Context, tx TX, ids []uint32) error
	GetEnvs(ctx context.Context, id uint32) (*Env, error)
	ListEnvs(ctx context.Context, tx TX, filter *EnvsFilter) ([]*Env, error)
	CountEnvs(ctx context.Context, tx TX, filter CountFilter) (int64, error)
}
