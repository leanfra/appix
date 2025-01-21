package repo

import "context"

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

type TagsRepo interface {
	RequireCounter
	CreateTags(ctx context.Context, tx TX, tags []*Tag) error
	UpdateTags(ctx context.Context, tx TX, tags []*Tag) error
	DeleteTags(ctx context.Context, tx TX, ids []uint32) error
	GetTags(ctx context.Context, id uint32) (*Tag, error)
	ListTags(ctx context.Context, tx TX, filter *TagsFilter) ([]*Tag, error)
	CountTags(ctx context.Context, tx TX, filter CountFilter) (int64, error)
}
