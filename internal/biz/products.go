package biz

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
)

type ProductsRepo interface {
	CreateProducts(ctx context.Context, ps []Product) error
	UpdateProducts(ctx context.Context, ps []Product) error
	DeleteProducts(ctx context.Context, ids []int64) error
	GetProducts(ctx context.Context, id int64) (*Product, error)
	ListProducts(ctx context.Context, filter *ListProductsFilter) ([]Product, error)
}

type ProductsUsecase struct {
	repo ProductsRepo
	log  *log.Helper
}

func NewProductsUsecase(repo ProductsRepo, logger log.Logger) *ProductsUsecase {
	return &ProductsUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (s *ProductsUsecase) validate(isNew bool, ps []Product) error {
	for _, p := range ps {
		if err := p.Validate(isNew); err != nil {
			return err
		}
	}
	return nil
}

// CreateProducts is
func (s *ProductsUsecase) CreateProducts(ctx context.Context, ps []Product) error {
	if err := s.validate(true, ps); err != nil {
		return err
	}
	return s.repo.CreateProducts(ctx, ps)
}

// UpdateProducts is
func (s *ProductsUsecase) UpdateProducts(ctx context.Context, ps []Product) error {
	if err := s.validate(false, ps); err != nil {
		return err
	}
	return s.repo.UpdateProducts(ctx, ps)
}

// DeleteProducts is
func (s *ProductsUsecase) DeleteProducts(ctx context.Context, ids []int64) error {
	if len(ids) == 0 {
		return fmt.Errorf("EmptyIds")
	}
	return s.repo.DeleteProducts(ctx, ids)
}

// GetProducts is
func (s *ProductsUsecase) GetProducts(ctx context.Context, id int64) (*Product, error) {
	if id <= 0 {
		return nil, fmt.Errorf("EmptyId")
	}
	return s.repo.GetProducts(ctx, id)
}

// ListProducts is
func (s *ProductsUsecase) ListProducts(ctx context.Context, filter *ListProductsFilter) ([]Product, error) {
	if filter != nil {
		if err := filter.Validate(); err != nil {
			return nil, err
		}
	}
	return s.repo.ListProducts(ctx, filter)
}
