package biz

import (
	"fmt"
)

func (m *Application) Validate(isNew bool) error {
	if len(m.Name) == 0 {
		return fmt.Errorf("InvalidNameValue")
	}
	if !isNew {
		if m.Id == 0 {
			return fmt.Errorf("InvalidId")
		}
	}
	return nil
}

func (m *ListApplicationsFilter) Validate() error {
	if m == nil {
		return nil
	}
	if len(m.Ids) > MaxFilterValues ||
		len(m.Names) > MaxFilterValues ||
		len(m.Clusters) > MaxFilterValues ||
		len(m.Datacenters) > MaxFilterValues ||
		len(m.Envs) > MaxFilterValues ||
		len(m.Products) > MaxFilterValues ||
		len(m.Teams) > MaxFilterValues ||
		len(m.Features) > MaxFilterValues ||
		len(m.Tags) > MaxFilterValues {

		return ErrFilterValuesExceedMax
	}

	if m.PageSize == 0 || m.PageSize > MaxPageSize {
		return ErrFilterInvalidPagesize
	}

	if m.IsStateful != IsStatefulFalse && m.IsStateful != IsStatefulTrue && m.IsStateful != IsStatefulNone {
		return fmt.Errorf("InvalidIsStateful")
	}

	return nil
}

func DefaultApplicationFilter() *ListApplicationsFilter {
	return &ListApplicationsFilter{
		Page:       1,
		PageSize:   DefaultPageSize,
		IsStateful: IsStatefulNone,
	}
}
