package repo

import "context"

const ProductTable = "products"

type Product struct {
	ID          uint32 `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"type:varchar(255);index:idx_product_name,unique"`
	Code        string `gorm:"type:varchar(255);index:idx_product_code,unique"`
	Description string `gorm:"type:varchar(255);"`
}

type ProductsFilter struct {
	Page     uint32
	PageSize uint32
	Names    []string
	Codes    []string
	Ids      []uint32
}

func (f *ProductsFilter) GetIds() []uint32 {
	return f.Ids
}

type ProductsRepo interface {
	RequireCounter
	CreateProducts(ctx context.Context, tx TX, ps []*Product) error
	UpdateProducts(ctx context.Context, tx TX, ps []*Product) error
	DeleteProducts(ctx context.Context, tx TX, ids []uint32) error
	GetProducts(ctx context.Context, id uint32) (*Product, error)
	ListProducts(ctx context.Context, tx TX, filter *ProductsFilter) ([]*Product, error)
	CountProducts(ctx context.Context, tx TX, filter CountFilter) (int64, error)
}
