package biz

import (
	"appix/internal/data/repo"
	"fmt"
)

func (f *Datacenter) Validate(isNew bool) error {
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

func (lf *ListDatacentersFilter) Validate() error {
	if lf == nil {
		return nil
	}
	if len(lf.Ids) > MaxFilterValues || len(lf.Names) > MaxFilterValues {
		return ErrFilterValuesExceedMax
	}
	if lf.PageSize == 0 || lf.PageSize > MaxPageSize {
		return ErrFilterInvalidPagesize
	}
	return nil
}

func DefaultDatacentersFilter() *ListDatacentersFilter {
	return &ListDatacentersFilter{
		Page:     1,
		PageSize: DefaultPageSize,
	}
}

func NewDatacenter(t *Datacenter) (*repo.Datacenter, error) {
	return &repo.Datacenter{
		ID:          t.Id,
		Name:        t.Name,
		Description: t.Description,
	}, nil
}

func NewDatacenters(es []*Datacenter) ([]*repo.Datacenter, error) {
	var clusters = make([]*repo.Datacenter, len(es))
	for i, f := range es {
		nf, err := NewDatacenter(f)
		if err != nil {
			return nil, err
		}
		clusters[i] = nf
	}
	return clusters, nil
}

func NewBizDatacenter(t *repo.Datacenter) (*Datacenter, error) {
	return &Datacenter{
		Id:          t.ID,
		Name:        t.Name,
		Description: t.Description,
	}, nil
}

func NewBizDatacenters(es []*repo.Datacenter) ([]*Datacenter, error) {
	var biz_clusters = make([]*Datacenter, len(es))
	for i, f := range es {
		biz_clusters[i] = &Datacenter{
			Id:          f.ID,
			Name:        f.Name,
			Description: f.Description,
		}
	}
	return biz_clusters, nil
}

func NewDatacentersFilter(filter *ListDatacentersFilter) *repo.DatacentersFilter {
	return &repo.DatacentersFilter{
		Ids:      filter.Ids,
		Names:    filter.Names,
		Page:     filter.Page,
		PageSize: filter.PageSize,
	}
}
