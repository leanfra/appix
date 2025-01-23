package biz

import (
	"opspillar/internal/data/repo"
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
	if e := ValidateName(f.Name); e != nil {
		return e
	}
	if e := ValidateCode(f.Code); e != nil {
		return e
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
	if lf.Page == 0 {
		return ErrFilterInvalidPage
	}

	return nil
}

func DefaultProductsFilter() *ListProductsFilter {
	return &ListProductsFilter{
		Page:     1,
		PageSize: DefaultPageSize,
	}
}

func ToDBProduct(t *Product) (*repo.Product, error) {
	return &repo.Product{
		ID:          t.Id,
		Name:        t.Name,
		Code:        t.Code,
		Description: t.Description,
	}, nil
}

func ToDBProducts(ts []*Product) ([]*repo.Product, error) {
	var products = make([]*repo.Product, len(ts))
	for i, t := range ts {
		nt, err := ToDBProduct(t)
		if err != nil {
			return nil, err
		}
		products[i] = nt
	}
	return products, nil
}

func ToBizProduct(t *repo.Product) (*Product, error) {
	return &Product{
		Id:          t.ID,
		Code:        t.Code,
		Description: t.Description,
		Name:        t.Name,
	}, nil
}

func ToBizProducts(ps []*repo.Product) ([]*Product, error) {
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

func ToDBProductFilter(filter *ListProductsFilter) *repo.ProductsFilter {
	return &repo.ProductsFilter{
		Codes:    filter.Codes,
		Ids:      filter.Ids,
		Names:    filter.Names,
		Page:     filter.Page,
		PageSize: filter.PageSize,
	}
}
