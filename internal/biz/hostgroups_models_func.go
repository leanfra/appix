package biz

import (
	"appix/internal/data/repo"
	"fmt"
)

func (f *Hostgroup) Validate(isNew bool) error {
	if len(f.Name) == 0 {
		return fmt.Errorf("InvalidNameValue")
	}
	if !isNew {
		if f.Id <= 0 {
			return fmt.Errorf("InvalidId")
		}
	}
	if e := ValidateName(f.Name); e != nil {
		return e
	}
	if f.ClusterId <= 0 {
		return fmt.Errorf("InvalidClusterId")
	}
	if f.DatacenterId <= 0 {
		return fmt.Errorf("InvalidDatacenterId")
	}
	if f.EnvId <= 0 {
		return fmt.Errorf("InvalidEnvId")
	}
	if f.ProductId <= 0 {
		return fmt.Errorf("InvalidProductId")
	}
	if f.TeamId <= 0 {
		return fmt.Errorf("InvalidTeamId")
	}

	return nil
}

func (lf *ListHostgroupsFilter) Validate() error {
	if lf == nil {
		return nil
	}

	if len(lf.Names) > MaxFilterValues ||
		len(lf.ClustersId) > MaxFilterValues ||
		len(lf.DatacentersId) > MaxFilterValues ||
		len(lf.EnvsId) > MaxFilterValues ||
		len(lf.ProductsId) > MaxFilterValues ||
		len(lf.TeamsId) > MaxFilterValues ||
		len(lf.FeaturesId) > MaxFilterValues ||
		len(lf.TagsId) > MaxFilterValues ||
		len(lf.ShareProductsId) > MaxFilterValues ||
		len(lf.ShareTeamsId) > MaxFilterValues {

		return ErrFilterValuesExceedMax
	}

	if lf.PageSize == 0 || lf.PageSize > MaxPageSize {
		return ErrFilterInvalidPagesize
	}
	return nil
}

func DefaultHostgroupFilter() *ListHostgroupsFilter {
	return &ListHostgroupsFilter{
		Page:     1,
		PageSize: DefaultPageSize,
	}
}

func ToDBHostgroup(t *Hostgroup) (*repo.Hostgroup, error) {
	return &repo.Hostgroup{
		Id:           t.Id,
		Name:         t.Name,
		Description:  t.Description,
		ClusterId:    t.ClusterId,
		DatacenterId: t.DatacenterId,
		EnvId:        t.EnvId,
		ProductId:    t.ProductId,
		TeamId:       t.TeamId,
	}, nil
}

func ToDBHostgroups(ts []*Hostgroup) ([]*repo.Hostgroup, error) {
	var products = make([]*repo.Hostgroup, len(ts))
	for i, t := range ts {
		nt, err := ToDBHostgroup(t)
		if err != nil {
			return nil, err
		}
		products[i] = nt
	}
	return products, nil
}

func ToBizHostgroup(t *repo.Hostgroup) (*Hostgroup, error) {
	return &Hostgroup{
		Id:           t.Id,
		Description:  t.Description,
		Name:         t.Name,
		ClusterId:    t.ClusterId,
		DatacenterId: t.DatacenterId,
		EnvId:        t.EnvId,
		ProductId:    t.ProductId,
		TeamId:       t.TeamId,
	}, nil
}

func ToBizHostgroups(ps []*repo.Hostgroup) ([]*Hostgroup, error) {
	var biz_ps []*Hostgroup
	for _, t := range ps {
		if t != nil {
			bhg, err := ToBizHostgroup(t)
			if err != nil {
				return nil, err
			}
			biz_ps = append(biz_ps, bhg)
		}
	}
	return biz_ps, nil
}

func ToDBHostgroupsFilter(filter *ListHostgroupsFilter) *repo.HostgroupsFilter {
	return &repo.HostgroupsFilter{
		Ids:           filter.Ids,
		Names:         filter.Names,
		ClustersId:    filter.ClustersId,
		DatacentersId: filter.DatacentersId,
		EnvsId:        filter.EnvsId,
		ProductsId:    filter.ProductsId,
		TeamsId:       filter.TeamsId,
	}
}
