package biz

import "fmt"

func (f *Product) Validate(isNew bool) error {
	if len(f.Name) == 0 || len(f.Code) == 0 {
		return fmt.Errorf("InvalidNameCode")
	}
	if !isNew {
		if f.Id <= 0 {
			return fmt.Errorf("InvalidId")
		}
	}
	return nil
}

func (lf *ListProductsFilter) Validate() error {
	if lf == nil {
		return nil
	}
	if len(lf.Codes) > MaxFilterValues ||
		len(lf.Ids) > MaxFilterValues ||
		len(lf.Names) > MaxFilterValues {

		return ErrFilterValuesExceedMax
	}

	return nil
}
