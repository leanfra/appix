package biz

import (
	"fmt"
)

func (f *Feature) Validate(isNew bool) error {
	if len(f.Name) == 0 || len(f.Value) == 0 {
		return fmt.Errorf("InvalidNameValue")
	}
	if !isNew {
		if f.Id <= 0 {
			return fmt.Errorf("InvalidId")
		}
	}
	return nil
}

func (lf *ListFeaturesFilter) Validate() error {
	if lf == nil {
		return nil
	}
	if len(lf.Ids) > MaxFilterValues ||
		len(lf.Names) > MaxFilterValues ||
		len(lf.Kvs) > MaxFilterValues {

		return ErrFilterValuesExceedMax
	}

	if len(lf.Kvs) > 0 {
		for _, kv := range lf.Kvs {
			if e := filterKvValidate(kv); e != nil {
				return e
			}
		}
	}
	return nil
}
