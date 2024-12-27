package biz

import (
	"appix/internal/data/repo"
	"fmt"
)

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
	if lf.PageSize == 0 || lf.PageSize > MaxPageSize {
		return ErrFilterInvalidPagesize
	}

	return nil
}

func DefaultProductsFilter() *ListProductsFilter {
	return &ListProductsFilter{
		Page:     1,
		PageSize: DefaultPageSize,
	}
}

func ToProductDB(t *Product) (*repo.Product, error) {
	return &repo.Product{
		ID:          t.Id,
		Name:        t.Name,
		Code:        t.Code,
		Description: t.Description,
	}, nil
}

func ToProductsDB(ts []*Product) ([]*repo.Product, error) {
	var products = make([]*repo.Product, len(ts))
	for i, t := range ts {
		nt, err := ToProductDB(t)
		if err != nil {
			return nil, err
		}
		products[i] = nt
	}
	return products, nil
}

func ToProductBiz(t *repo.Product) (*Product, error) {
	return &Product{
		Id:          t.ID,
		Code:        t.Code,
		Description: t.Description,
		Name:        t.Name,
	}, nil
}

func ToProductsBiz(ps []*repo.Product) ([]*Product, error) {
	var _ps = make([]*Product, len(ps))
	for i, t := range ps {
		_ps[i] = &Product{
			Id:          t.ID,
			Code:        t.Code,
			Description: t.Description,
			Name:        t.Name,
		}
	}
	return _ps, nil
}

func ToProductFilterDB(filter *ListProductsFilter) *repo.ProductsFilter {
	return &repo.ProductsFilter{
		Codes:    filter.Codes,
		Ids:      filter.Ids,
		Names:    filter.Names,
		Page:     filter.Page,
		PageSize: filter.PageSize,
	}
}
