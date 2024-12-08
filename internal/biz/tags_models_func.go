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

func (f *TagFilter) Validate() error {
	if len(f.Key) == 0 || len(f.Value) == 0 {
		return fmt.Errorf("InvalidTagFilterKeyValue")
	}
	return nil
}

func (lf *ListTagsFilter) Validate() error {
	if lf.Page < 0 || lf.PageSize < 0 {
		return fmt.Errorf("ListTagFilterInvliadPagePagesize")
	}
	for _, f := range lf.Filters {
		if err := f.Validate(); err != nil {
			return err
		}
	}
	return nil
}
