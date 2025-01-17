package biz

import (
	"appix/internal/data"
	"appix/internal/data/repo"
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
)

type ProductsUsecase struct {
	txm       repo.TxManager
	prdrepo   repo.ProductsRepo
	authzrepo repo.AuthzRepo
	log       *log.Helper
	required  []requiredBy
}

func NewProductsUsecase(
	repo repo.ProductsRepo,
	authzrepo repo.AuthzRepo,
	hostgrouprepo repo.HostgroupsRepo,
	apprepo repo.ApplicationsRepo,
	hprepo repo.HostgroupProductsRepo,
	logger log.Logger,
	txm repo.TxManager) *ProductsUsecase {
	return &ProductsUsecase{
		prdrepo:   repo,
		authzrepo: authzrepo,
		log:       log.NewHelper(logger),
		txm:       txm,
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

func (s *ProductsUsecase) enforce(ctx context.Context, tx repo.TX) error {
	curUser := ctx.Value(data.UserName).(string)
	ires := repo.NewResource4Sv1("team", "", "", "")
	can, err := s.authzrepo.Enforce(ctx, tx, &repo.AuthenRequest{
		Sub:      curUser,
		Resource: ires,
		Action:   repo.ActWrite,
	})
	if err != nil {
		return err
	}
	if !can {
		return fmt.Errorf("PermissionDenied")
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
	//return s.prdrepo.CreateProducts(ctx, _ps)
	err := s.txm.RunInTX(func(tx repo.TX) error {
		if err := s.enforce(ctx, tx); err != nil {
			return err
		}
		if e := s.prdrepo.CreateProducts(ctx, tx, _ps); e != nil {
			return e
		}
		return nil
	})
	return err
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
	//return s.prdrepo.UpdateProducts(ctx, dps)
	err := s.txm.RunInTX(func(tx repo.TX) error {
		if err := s.enforce(ctx, tx); err != nil {
			return err
		}
		if e := s.prdrepo.UpdateProducts(ctx, tx, dps); e != nil {
			return e
		}
		return nil
	})
	return err
}

// DeleteProducts is
func (s *ProductsUsecase) DeleteProducts(ctx context.Context, ids []uint32) error {
	if len(ids) == 0 {
		return fmt.Errorf("EmptyIds")
	}
	return s.txm.RunInTX(func(tx repo.TX) error {
		if err := s.enforce(ctx, tx); err != nil {
			return err
		}
		for _, r := range s.required {
			c, err := r.inst.CountRequire(ctx, tx, repo.RequireProduct, ids)
			if err != nil {
				return err
			}
			if c > 0 {
				return fmt.Errorf("some %s requires", r.name)
			}
		}
		if e := s.prdrepo.DeleteProducts(ctx, tx, ids); e != nil {
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
	ps, e := s.prdrepo.GetProducts(ctx, id)
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
	dbps, e := s.prdrepo.ListProducts(ctx, nil, ToDBProductFilter(filter))
	if e != nil {
		return nil, e
	}
	return ToBizProducts(dbps)
}
