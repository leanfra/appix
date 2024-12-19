package biz

import "fmt"

func (f *Env) Validate(isNew bool) error {
	if len(f.Name) == 0 {
		return fmt.Errorf("InvalidNameValue")
	}
	if !isNew {
		if f.Id <= 0 {
			return fmt.Errorf("InvalidId")
		}
	}
	return nil
}

func (lf *ListEnvsFilter) Validate() error {
	if lf == nil {
		return nil
	}
	if len(lf.Names) > MaxFilterValues || len(lf.Ids) > MaxFilterValues {
		return ErrFilterValuesExceedMax
	}
	return nil
}
