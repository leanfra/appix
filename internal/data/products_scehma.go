package data

import (
	"appix/internal/biz"
)

type Product struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"type:varchar(255);index:idx_name,unique"`
	Code        string `gorm:"type:varchar(255);index:idx_code,unique"`
	Description string `gorm:"type:varchar(255);"`
}

func NewProduct(t biz.Product) (*Product, error) {
	return &Product{
		ID:          uint(t.Id),
		Name:        t.Name,
		Code:        t.Code,
		Description: t.Description,
	}, nil
}

func NewProducts(ts []biz.Product) ([]*Product, error) {
	var products = make([]*Product, len(ts))
	for i, t := range ts {
		nt, err := NewProduct(t)
		if err != nil {
			return nil, err
		}
		products[i] = nt
	}
	return products, nil
}

func NewBizProduct(t *Product) (*biz.Product, error) {
	return &biz.Product{
		Id:          int64(t.ID),
		Code:        t.Code,
		Description: t.Description,
		Name:        t.Name,
	}, nil
}

func NewBizProducts(ps []Product) ([]biz.Product, error) {
	var biz_ps = make([]biz.Product, len(ps))
	for i, t := range ps {
		biz_ps[i] = biz.Product{
			Id:          int64(t.ID),
			Code:        t.Code,
			Description: t.Description,
			Name:        t.Name,
		}
	}
	return biz_ps, nil
}
