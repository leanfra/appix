package data

import (
	"appix/internal/biz"
)

type Hostgroup struct {
	ID           uint32 `gorm:"primaryKey;autoIncrement"`
	Name         string `gorm:"type:varchar(255);index:idx_name,unique"`
	Description  string `gorm:"type:varchar(255);"`
	ClusterId    uint32
	DatacenterId uint32
	EnvId        uint32
	ProductId    uint32
	TeamId       uint32
}

func NewHostgroup(t *biz.Hostgroup) (*Hostgroup, error) {
	return &Hostgroup{
		ID:           t.Id,
		Name:         t.Name,
		Description:  t.Description,
		ClusterId:    t.ClusterId,
		DatacenterId: t.DatacenterId,
		EnvId:        t.EnvId,
		ProductId:    t.ProductId,
		TeamId:       t.TeamId,
	}, nil
}

func NewHostgroups(ts []*biz.Hostgroup) ([]*Hostgroup, error) {
	var products = make([]*Hostgroup, len(ts))
	for i, t := range ts {
		nt, err := NewHostgroup(t)
		if err != nil {
			return nil, err
		}
		products[i] = nt
	}
	return products, nil
}

func NewBizHostgroup(t *Hostgroup) (*biz.Hostgroup, error) {
	return &biz.Hostgroup{
		Id:           t.ID,
		Description:  t.Description,
		Name:         t.Name,
		ClusterId:    t.ClusterId,
		DatacenterId: t.DatacenterId,
		EnvId:        t.EnvId,
		ProductId:    t.ProductId,
		TeamId:       t.TeamId,
	}, nil
}

func NewBizHostgroups(ps []*Hostgroup) ([]*biz.Hostgroup, error) {
	var biz_ps []*biz.Hostgroup
	for _, t := range ps {
		if t != nil {
			bhg, err := NewBizHostgroup(t)
			if err != nil {
				return nil, err
			}
			biz_ps = append(biz_ps, bhg)
		}
	}
	return biz_ps, nil
}
