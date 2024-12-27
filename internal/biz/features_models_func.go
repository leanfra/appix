package biz

import (
	"appix/internal/data/repo"
	"fmt"
)

func (f *Feature) Validate(isNew bool) error {
	if len(f.Name) == 0 || len(f.Value) == 0 {
		return fmt.Errorf("InvalidNameValue")
	}
	if !isNew {
		if f.Id <= 0 {
			return fmt.Errorf("InvalidId")
		}
	}
	return nil
}

func (lf *ListFeaturesFilter) Validate() error {
	if lf == nil {
		return nil
	}
	if len(lf.Ids) > MaxFilterValues ||
		len(lf.Names) > MaxFilterValues ||
		len(lf.Kvs) > MaxFilterValues {

		return ErrFilterValuesExceedMax
	}

	if len(lf.Kvs) > 0 {
		for _, kv := range lf.Kvs {
			if e := filterKvValidate(kv); e != nil {
				return e
			}
		}
	}
	if lf.PageSize == 0 || lf.PageSize > MaxPageSize {
		return ErrFilterInvalidPagesize
	}

	return nil
}

func DefaultFeaturesFilter() *ListFeaturesFilter {
	return &ListFeaturesFilter{
		Page:     1,
		PageSize: DefaultPageSize,
	}
}

func NewFeature(t *Feature) (*repo.Feature, error) {
	return &repo.Feature{
		Id:    t.Id,
		Name:  t.Name,
		Value: t.Value,
	}, nil
}

func NewFeatures(fs []*Feature) ([]*repo.Feature, error) {
	var features = make([]*repo.Feature, len(fs))
	for i, f := range fs {
		nf, err := NewFeature(f)
		if err != nil {
			return nil, err
		}
		features[i] = nf
	}
	return features, nil
}

func NewBizFeature(t *repo.Feature) (*Feature, error) {
	return &Feature{
		Id:    t.Id,
		Name:  t.Name,
		Value: t.Value,
	}, nil
}

func NewBizFeatures(fs []*repo.Feature) ([]*Feature, error) {
	var biz_fts = make([]*Feature, len(fs))
	for i, f := range fs {
		biz_fts[i] = &Feature{
			Id:    f.Id,
			Name:  f.Name,
			Value: f.Value,
		}
	}
	return biz_fts, nil
}

func NewFeaturesFilter(filter *ListFeaturesFilter) *repo.FeaturesFilter {
	return &repo.FeaturesFilter{
		Ids:      filter.Ids,
		Kvs:      filter.Kvs,
		Names:    filter.Names,
		Page:     filter.Page,
		PageSize: filter.PageSize,
	}
}
