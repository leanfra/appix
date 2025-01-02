package repo

const TagTable = "tags"

type Tag struct {
	ID          uint32 `gorm:"primaryKey;autoIncrement"`
	Key         string `gorm:"type:varchar(255);index:idx_key_value,unique"`
	Value       string `gorm:"type:varchar(255);index:idx_key_value"`
	Description string `gorm:"type:text"`
}

type TagsFilter struct {
	Page     uint32
	PageSize uint32
	Keys     []string
	Kvs      []string
	Ids      []uint32
}

func (f *TagsFilter) GetIds() []uint32 {
	return f.Ids
}
