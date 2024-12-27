package repo

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
