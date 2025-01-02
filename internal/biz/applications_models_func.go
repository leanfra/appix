package biz

import (
	"appix/internal/data/repo"
	"fmt"
)

func (m *Application) Validate(isNew bool) error {
	if len(m.Name) == 0 {
		return fmt.Errorf("InvalidNameValue")
	}
	if e := ValidateName(m.Name); e != nil {
		return e
	}
	if !isNew {
		if m.Id == 0 {
			return fmt.Errorf("InvalidId")
		}
	}
	if len(m.Name) == 0 {
		return fmt.Errorf("InvalidNameValue")
	}
	if len(m.Owner) == 0 {
		return fmt.Errorf("InvalidOwnerValue")
	}
	if m.ProductId <= 0 {
		return fmt.Errorf("InvalidProductId")
	}
	if m.TeamId <= 0 {
		return fmt.Errorf("InvalidTeamId")
	}
	return nil
}

func (m *ListApplicationsFilter) Validate() error {
	if m == nil {
		return nil
	}
	if len(m.Ids) > MaxFilterValues ||
		len(m.Names) > MaxFilterValues ||
		len(m.ProductsId) > MaxFilterValues ||
		len(m.TeamsId) > MaxFilterValues ||
		len(m.FeaturesId) > MaxFilterValues ||
		len(m.HostgroupsId) > MaxFilterValues ||
		len(m.TagsId) > MaxFilterValues {

		return ErrFilterValuesExceedMax
	}

	if m.PageSize == 0 || m.PageSize > MaxPageSize {
		return ErrFilterInvalidPagesize
	}
	if m.Page == 0 {
		return ErrFilterInvalidPage
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

func ToDBApplication(app *Application) (*repo.Application, error) {
	if app == nil {
		return nil, nil
	}
	return &repo.Application{
		Id:          app.Id,
		Name:        app.Name,
		Description: app.Description,
		Owner:       app.Owner,
		IsStateful:  app.IsStateful,
		ProductId:   app.ProductId,
		TeamId:      app.TeamId,
	}, nil
}

func ToDBApplications(apps []*Application) ([]*repo.Application, error) {
	db_apps := make([]*repo.Application, len(apps))
	for i, a := range apps {
		db_app, e := ToDBApplication(a)
		if e != nil {
			return nil, e
		}
		db_apps[i] = db_app
	}
	return db_apps, nil
}

func ToBizApplication(t *repo.Application) (*Application, error) {
	if t == nil {
		return nil, nil
	}
	return &Application{
		Id:          t.Id,
		Name:        t.Name,
		Description: t.Description,
		Owner:       t.Owner,
		IsStateful:  t.IsStateful,
		ProductId:   t.ProductId,
		TeamId:      t.TeamId,
	}, nil
}

func ToBizApplications(es []*repo.Application) ([]*Application, error) {
	apps := make([]*Application, len(es))
	for i, e := range es {
		app, err := ToBizApplication(e)
		if err != nil {
			return nil, err
		}
		apps[i] = app
	}
	return apps, nil
}

func ToDBApplicationsFilter(filter *ListApplicationsFilter) *repo.ApplicationsFilter {
	if filter == nil {
		return nil
	}
	return &repo.ApplicationsFilter{
		Ids:        filter.Ids,
		Names:      filter.Names,
		ProductsId: filter.ProductsId,
		TeamsId:    filter.TeamsId,
		IsStateful: filter.IsStateful,
		Page:       filter.Page,
		PageSize:   filter.PageSize,
	}
}
