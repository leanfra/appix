package data

import (
	"appix/internal/biz"
)

type Tag struct {
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Key   string `gorm:"type:varchar(255);index:idx_key_value,unique"`
	Value string `gorm:"type:varchar(255);index:idx_key_value"`
}

func NewTag(t biz.Tag) (*Tag, error) {
	return &Tag{
		ID:    uint(t.Id),
		Key:   t.Key,
		Value: t.Value,
	}, nil
}

func NewTags(ts []biz.Tag) ([]*Tag, error) {
	var tags []*Tag
	for _, t := range ts {
		nt, err := NewTag(t)
		if err != nil {
			return nil, err
		}
		tags = append(tags, nt)
	}
	return tags, nil
}
