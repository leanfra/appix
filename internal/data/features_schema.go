package data

import (
	"appix/internal/biz"
)

type Feature struct {
	ID    uint32 `gorm:"primaryKey;autoIncrement"`
	Name  string `gorm:"type:varchar(255);index:idx_feature_name_value,unique"`
	Value string `gorm:"type:varchar(255);index:idx_feature_name_value"`
}

func NewFeature(t *biz.Feature) (*Feature, error) {
	return &Feature{
		ID:    t.Id,
		Name:  t.Name,
		Value: t.Value,
	}, nil
}

func NewFeatures(fs []*biz.Feature) ([]*Feature, error) {
	var features = make([]*Feature, len(fs))
	for i, f := range fs {
		nf, err := NewFeature(f)
		if err != nil {
			return nil, err
		}
		features[i] = nf
	}
	return features, nil
}

func NewBizFeature(t *Feature) (*biz.Feature, error) {
	return &biz.Feature{
		Id:    t.ID,
		Name:  t.Name,
		Value: t.Value,
	}, nil
}

func NewBizFeatures(fs []Feature) ([]biz.Feature, error) {
	var biz_fts = make([]biz.Feature, len(fs))
	for i, f := range fs {
		biz_fts[i] = biz.Feature{
			Id:    f.ID,
			Name:  f.Name,
			Value: f.Value,
		}
	}
	return biz_fts, nil
}
