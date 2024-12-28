package repo

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
