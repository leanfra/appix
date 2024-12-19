package data

import (
	"appix/internal/biz"
)

type Tag struct {
	ID    uint32 `gorm:"primaryKey;autoIncrement"`
	Key   string `gorm:"type:varchar(255);index:idx_key_value,unique"`
	Value string `gorm:"type:varchar(255);index:idx_key_value"`
}

func NewTag(t *biz.Tag) (*Tag, error) {
	if t == nil {
		return nil, nil
	}
	return &Tag{
		ID:    t.Id,
		Key:   t.Key,
		Value: t.Value,
	}, nil
}

func NewTags(ts []*biz.Tag) ([]*Tag, error) {
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
		Id:    t.ID,
		Key:   t.Key,
		Value: t.Value,
	}, nil
}

func NewBizTags(tags []Tag) ([]biz.Tag, error) {
	var biz_tags = make([]biz.Tag, len(tags))
	for i, t := range tags {
		biz_tags[i] = biz.Tag{
			Id:    t.ID,
			Key:   t.Key,
			Value: t.Value,
		}
	}
	return biz_tags, nil
}
