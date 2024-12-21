package biz

import "fmt"

func (f *Team) Validate(isNew bool) error {
	if len(f.Name) == 0 || len(f.Code) == 0 || len(f.Leader) == 0 {
		return fmt.Errorf("InvalidNameCodeLeader")
	}
	if !isNew {
		if f.Id <= 0 {
			return fmt.Errorf("InvalidId")
		}
	}
	return nil
}

func (lf *ListTeamsFilter) Validate() error {
	if lf == nil {
		return nil
	}
	if len(lf.Codes) > MaxFilterValues ||
		len(lf.Ids) > MaxFilterValues ||
		len(lf.Leaders) > MaxFilterValues ||
		len(lf.Names) > MaxFilterValues {
		return ErrFilterValuesExceedMax
	}
	if lf.PageSize == 0 || lf.PageSize > MaxPageSize {
		return ErrFilterInvalidPagesize
	}
	return nil
}

func DefaultTeamsFilter() *ListTeamsFilter {
	return &ListTeamsFilter{
		Page:     1,
		PageSize: DefaultPageSize,
	}
}
