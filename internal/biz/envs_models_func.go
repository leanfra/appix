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

func (ff *EnvFilter) Validate() error {
	if len(ff.Name) == 0 {
		return fmt.Errorf("InvalidEnvFilterNameValue")
	}
	return nil
}

func (lf *ListEnvsFilter) Validate() error {
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
