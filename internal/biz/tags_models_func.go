package biz

import (
	"appix/internal/data/repo"
	"fmt"
)

func (t *Tag) Validate(isNew bool) error {
	if len(t.Key) == 0 || len(t.Value) == 0 {
		return fmt.Errorf("InvalidKeyValue")
	}
	if !isNew {
		if t.Id <= 0 {
			return fmt.Errorf("InvalidId")
		}
	}
	return nil
}

func (lf *ListTagsFilter) Validate() error {
	if lf == nil {
		return nil
	}
	if len(lf.Ids) > MaxFilterValues || len(lf.Keys) > MaxFilterValues || len(lf.Kvs) > MaxFilterValues {
		return ErrFilterValuesExceedMax
	}
	for _, kv := range lf.Kvs {
		if e := filterKvValidate(kv); e != nil {
			return e
		}
	}
	if lf.PageSize == 0 || lf.PageSize > MaxPageSize {
		return ErrFilterInvalidPagesize
	}
	return nil
}

func DefaultTagsFilter() *ListTagsFilter {
	return &ListTagsFilter{
		Page:     1,
		PageSize: DefaultPageSize,
	}
}

func ToTagDB(t *Tag) (*repo.Tag, error) {
	if t == nil {
		return nil, nil
	}
	return &repo.Tag{
		ID:    t.Id,
		Key:   t.Key,
		Value: t.Value,
	}, nil
}

func ToTagsDB(ts []*Tag) ([]*repo.Tag, error) {
	var tags = make([]*repo.Tag, len(ts))
	for i, t := range ts {
		nt, err := ToTagDB(t)
		if err != nil {
			return nil, err
		}
		tags[i] = nt
	}
	return tags, nil
}

func ToTagBiz(t *repo.Tag) (*Tag, error) {
	return &Tag{
		Id:    t.ID,
		Key:   t.Key,
		Value: t.Value,
	}, nil
}

func ToTagsBiz(tags []*repo.Tag) ([]*Tag, error) {
	var biz_tags = make([]*Tag, len(tags))
	for i, t := range tags {
		biz_tags[i] = &Tag{
			Id:    t.ID,
			Key:   t.Key,
			Value: t.Value,
		}
	}
	return biz_tags, nil
}

func ToTagsFilterDB(filter *ListTagsFilter) *repo.TagsFilter {
	if filter == nil {
		return nil
	}
	return &repo.TagsFilter{
		Ids:      filter.Ids,
		Keys:     filter.Keys,
		Kvs:      filter.Kvs,
		Page:     filter.Page,
		PageSize: filter.PageSize,
	}
}
