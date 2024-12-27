package sqldb

import (
	"appix/internal/data/repo"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type ProductsRepoGorm struct {
	data *DataGorm
	log  *log.Helper
}

func NewProductsRepoGorm(data *DataGorm, logger log.Logger) (repo.ProductsRepo, error) {

	if err := validateData(data); err != nil {
		return nil, err
	}

	if err := initTable(data.DB, &repo.Product{}, repo.ProductTable); err != nil {
		return nil, err
	}

	return &ProductsRepoGorm{
		data: data,
		log:  log.NewHelper(logger),
	}, nil
}

// CreateProducts is
func (d *ProductsRepoGorm) CreateProducts(ctx context.Context, ps []*repo.Product) error {
	r := d.data.DB.WithContext(ctx).Create(ps)
	if r.Error != nil {
		return r.Error
	}

	return nil
}

// UpdateProducts is
func (d *ProductsRepoGorm) UpdateProducts(ctx context.Context, ps []*repo.Product) error {

	r := d.data.DB.WithContext(ctx).Save(ps)
	if r.Error != nil {
		return r.Error
	}

	return nil
}

// DeleteProducts is
func (d *ProductsRepoGorm) DeleteProducts(ctx context.Context, ids []uint32) error {

	r := d.data.DB.WithContext(ctx).Where("id in (?)", ids).Delete(&repo.Product{})
	if r.Error != nil {
		return r.Error
	}
	return nil
}

// GetProducts is
func (d *ProductsRepoGorm) GetProducts(ctx context.Context, id uint32) (*repo.Product, error) {

	product := &repo.Product{}
	r := d.data.DB.WithContext(ctx).Where("id = ?", id).First(product)
	if r.Error != nil {
		return nil, r.Error
	}
	return product, nil
}

// ListProducts is
func (d *ProductsRepoGorm) ListProducts(ctx context.Context,
	tx repo.TX,
	filter *repo.ProductsFilter) ([]*repo.Product, error) {

	db_ps := []*repo.Product{}
	query := d.data.WithTX(tx).WithContext(ctx)
	if filter != nil {
		var offset int
		if filter.Page > 0 && filter.PageSize > 0 {
			offset = int((filter.Page - 1) * filter.PageSize)
			query = query.Offset(offset).Limit(int(filter.PageSize))
		}
		if len(filter.Names) > 0 {
			nameConditions := buildOrLike("name", len(filter.Names))
			query = query.Where(nameConditions, filter.Names)
		}
		if len(filter.Codes) > 0 {
			codeConditions := buildOrLike("code", len(filter.Codes))
			query = query.Where(codeConditions, filter.Codes)
		}
		if len(filter.Ids) > 0 {
			query = query.Where("id in (?)", filter.Ids)
		}
	}
	r := query.Find(&db_ps)
	if r.Error != nil {
		return nil, r.Error
	}
	return db_ps, nil
}

// CountProducts is count by id
func (d *ProductsRepoGorm) CountProducts(ctx context.Context,
	tx repo.TX,
	filter repo.CountFilter) (int64, error) {

	var db_ps int64
	query := d.data.WithTX(tx).WithContext(ctx).Model(&repo.Product{})
	if filter != nil {
		if len(filter.GetIds()) > 0 {
			query = query.Where("id in (?)", filter.GetIds())
		}
	}
	r := query.Count(&db_ps)
	if r.Error != nil {
		return 0, r.Error
	}

	return db_ps, nil
}
