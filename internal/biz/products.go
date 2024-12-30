package biz

import (
	"appix/internal/data/repo"
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
)

type ProductsUsecase struct {
	txm      repo.TxManager
	repo     repo.ProductsRepo
	log      *log.Helper
	required []requiredBy
}

func NewProductsUsecase(
	repo repo.ProductsRepo,
	hostgrouprepo repo.HostgroupsRepo,
	apprepo repo.ApplicationsRepo,
	hprepo repo.HostgroupProductsRepo,
	logger log.Logger,
	txm repo.TxManager) *ProductsUsecase {
	return &ProductsUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
		txm:  txm,
		required: []requiredBy{
			{inst: hostgrouprepo, name: "hostgroup"},
			{inst: apprepo, name: "app"},
			{inst: hprepo, name: "hostgroup_product"},
		},
	}
}

func (s *ProductsUsecase) validate(isNew bool, ps []*Product) error {
	for _, p := range ps {
		if err := p.Validate(isNew); err != nil {
			return err
		}
	}
	return nil
}

// CreateProducts is
func (s *ProductsUsecase) CreateProducts(ctx context.Context, ps []*Product) error {
	if err := s.validate(true, ps); err != nil {
		return err
	}
	_ps, e := ToDBProducts(ps)
	if e != nil {
		return e
	}
	return s.repo.CreateProducts(ctx, _ps)
}

// UpdateProducts is
func (s *ProductsUsecase) UpdateProducts(ctx context.Context, ps []*Product) error {
	if err := s.validate(false, ps); err != nil {
		return err
	}
	dps, e := ToDBProducts(ps)
	if e != nil {
		return e
	}
	return s.repo.UpdateProducts(ctx, dps)
}

// DeleteProducts is
func (s *ProductsUsecase) DeleteProducts(ctx context.Context, ids []uint32) error {
	if len(ids) == 0 {
		return fmt.Errorf("EmptyIds")
	}
	return s.txm.RunInTX(func(tx repo.TX) error {
		for _, r := range s.required {
			c, err := r.inst.CountRequire(ctx, tx, repo.RequireProduct, ids)
			if err != nil {
				return err
			}
			if c > 0 {
				return fmt.Errorf("some %s requires", r.name)
			}
		}
		if e := s.repo.DeleteProducts(ctx, tx, ids); e != nil {
			return e
		}
		return nil
	})
}

// GetProducts is
func (s *ProductsUsecase) GetProducts(ctx context.Context, id uint32) (*Product, error) {
	if id <= 0 {
		return nil, fmt.Errorf("EmptyId")
	}
	ps, e := s.repo.GetProducts(ctx, id)
	if e != nil {
		return nil, e
	}
	return ToBizProduct(ps)
}

// ListProducts is
func (s *ProductsUsecase) ListProducts(ctx context.Context, filter *ListProductsFilter) ([]*Product, error) {
	if filter != nil {
		if err := filter.Validate(); err != nil {
			return nil, err
		}
	}
	dbps, e := s.repo.ListProducts(ctx, nil, ToDBProductFilter(filter))
	if e != nil {
		return nil, e
	}
	return ToBizProducts(dbps)
}
