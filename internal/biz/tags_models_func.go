package biz

import "fmt"

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
