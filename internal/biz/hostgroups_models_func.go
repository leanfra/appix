package biz

import "fmt"

func (f *Hostgroup) Validate(isNew bool) error {
	if len(f.Name) == 0 {
		return fmt.Errorf("InvalidNameValue")
	}
	if !isNew {
		if f.Id <= 0 {
			return fmt.Errorf("InvalidId")
		}
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
		len(lf.Clusters) > MaxFilterValues ||
		len(lf.Datacenters) > MaxFilterValues ||
		len(lf.Envs) > MaxFilterValues ||
		len(lf.Products) > MaxFilterValues ||
		len(lf.Teams) > MaxFilterValues ||
		len(lf.Features) > MaxFilterValues ||
		len(lf.Tags) > MaxFilterValues ||
		len(lf.ShareProducts) > MaxFilterValues ||
		len(lf.ShareTeams) > MaxFilterValues {

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
