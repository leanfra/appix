package data

import (
	"appix/internal/biz"
	"context"

	"github.com/go-kratos/kratos/v2/log"
	//  TODO: modify project name
	// biz "appix/internal/biz"
)

type ProductsRepoImpl struct {
	data *Data
	log  *log.Helper
}

func NewProductsRepoImpl(data *Data, logger log.Logger) (biz.ProductsRepo, error) {

	if err := validateData(data); err != nil {
		return nil, err
	}

	if err := initTable(data.db, &Product{}, "product"); err != nil {
		return nil, err
	}

	return &ProductsRepoImpl{
		data: data,
		log:  log.NewHelper(logger),
	}, nil
}

// CreateProducts is
func (d *ProductsRepoImpl) CreateProducts(ctx context.Context, ps []biz.Product) error {
	db_ps, err := NewProducts(ps)
	if err != nil {
		return err
	}
	r := d.data.db.WithContext(ctx).Create(db_ps)
	if r.Error != nil {
		return r.Error
	}

	return nil
}

// UpdateProducts is
func (d *ProductsRepoImpl) UpdateProducts(ctx context.Context, ps []biz.Product) error {

	db_ps, err := NewProducts(ps)
	if err != nil {
		return err
	}
	r := d.data.db.WithContext(ctx).Save(db_ps)
	if r.Error != nil {
		return r.Error
	}

	return nil
}

// DeleteProducts is
func (d *ProductsRepoImpl) DeleteProducts(ctx context.Context, ids []int64) error {

	r := d.data.db.WithContext(ctx).Where("id in (?)", ids).Delete(&Product{})
	if r.Error != nil {
		return r.Error
	}
	return nil
}

// GetProducts is
func (d *ProductsRepoImpl) GetProducts(ctx context.Context, id int64) (*biz.Product, error) {

	product := &Product{}
	r := d.data.db.WithContext(ctx).Where("id = ?", id).First(product)
	if r.Error != nil {
		return nil, r.Error
	}
	return NewBizProduct(product)
}

// ListProducts is
func (d *ProductsRepoImpl) ListProducts(ctx context.Context,
	filter *biz.ListProductsFilter) ([]biz.Product, error) {

	db_ps := []Product{}
	query := d.data.db.WithContext(ctx)
	if filter != nil {
		var offset int
		if filter.Page > 0 && filter.PageSize > 0 {
			offset = int((filter.Page - 1) * filter.PageSize)
			query = query.Offset(offset).Limit(int(filter.PageSize))
		}

		orConditions := make([]interface{}, len(filter.Filters))

		for i, pair := range filter.Filters {
			andConditions := map[string]string{}
			if pair.Code != "" {
				andConditions["code"] = pair.Code
			}
			if pair.Name != "" {
				andConditions["name"] = pair.Name
			}
			orConditions[i] = andConditions
		}
		query = query.Where(orConditions)
	}
	r := query.Find(&db_ps)
	if r.Error != nil {
		return nil, r.Error
	}
	return NewBizProducts(db_ps)
}
