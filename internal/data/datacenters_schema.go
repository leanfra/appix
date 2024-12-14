package data

import (
	"appix/internal/biz"
)

type Datacenter struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"type:varchar(255);index:idx_name,unique"`
	Description string `gorm:"type:varchar(255);"`
}

func NewDatacenter(t biz.Datacenter) (*Datacenter, error) {
	return &Datacenter{
		ID:          uint(t.Id),
		Name:        t.Name,
		Description: t.Description,
	}, nil
}

func NewDatacenters(es []biz.Datacenter) ([]*Datacenter, error) {
	var clusters = make([]*Datacenter, len(es))
	for i, f := range es {
		nf, err := NewDatacenter(f)
		if err != nil {
			return nil, err
		}
		clusters[i] = nf
	}
	return clusters, nil
}

func NewBizDatacenter(t *Datacenter) (*biz.Datacenter, error) {
	return &biz.Datacenter{
		Id:          int64(t.ID),
		Name:        t.Name,
		Description: t.Description,
	}, nil
}

func NewBizDatacenters(es []Datacenter) ([]biz.Datacenter, error) {
	var biz_clusters = make([]biz.Datacenter, len(es))
	for i, f := range es {
		biz_clusters[i] = biz.Datacenter{
			Id:          int64(f.ID),
			Name:        f.Name,
			Description: f.Description,
		}
	}
	return biz_clusters, nil
}
