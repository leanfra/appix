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

func (ff *ProductFilter) Validate() error {
	if len(ff.Name) == 0 || len(ff.Code) == 0 {
		return fmt.Errorf("InvalidTeamFilter")
	}
	return nil
}

func (lf *ListProductsFilter) Validate() error {
	if lf.Page < 0 || lf.PageSize < 0 {
		return fmt.Errorf("InvalidPageSize")
	}
	for _, f := range lf.Filters {
		if err := f.Validate(); err != nil {
			return err
		}
	}
	return nil
}
