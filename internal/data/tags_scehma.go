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
	var tags = make([]*Tag, len(ts))
	for i, t := range ts {
		nt, err := NewTag(t)
		if err != nil {
			return nil, err
		}
		tags[i] = nt
	}
	return tags, nil
}

func NewBizTag(t *Tag) (*biz.Tag, error) {
	return &biz.Tag{
		Id:    int64(t.ID),
		Key:   t.Key,
		Value: t.Value,
	}, nil
}

func NewBizTags(tags []Tag) ([]biz.Tag, error) {
	var biz_tags = make([]biz.Tag, len(tags))
	for i, t := range tags {
		biz_tags[i] = biz.Tag{
			Id:    int64(t.ID),
			Key:   t.Key,
			Value: t.Value,
		}
	}
	return biz_tags, nil
}
