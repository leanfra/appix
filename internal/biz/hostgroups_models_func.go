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
	if lf.Page < 0 || lf.PageSize < 0 {
		return fmt.Errorf("InvalidPageSize")
	}
	return nil
}
