package biz

import (
	"appix/internal/data"
	"appix/internal/data/repo"
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
)

type FeaturesUsecase struct {
	ftrepo    repo.FeaturesRepo
	authzrepo repo.AuthzRepo
	log       *log.Helper
	txm       repo.TxManager
	required  []requiredBy
}

func NewFeaturesUsecase(
	repo repo.FeaturesRepo,
	authzrepo repo.AuthzRepo,
	hgftrepo repo.HostgroupFeaturesRepo,
	appftrepo repo.AppFeaturesRepo,
	logger log.Logger,
	txm repo.TxManager) *FeaturesUsecase {
	return &FeaturesUsecase{
		ftrepo:    repo,
		authzrepo: authzrepo,
		log:       log.NewHelper(logger),
		txm:       txm,
		required: []requiredBy{
			{inst: hgftrepo, name: "hostgroup_feature"},
			{inst: appftrepo, name: "app_feature"},
		},
	}
}

func (s *FeaturesUsecase) validate(isNew bool, features []*Feature) error {
	for _, f := range features {
		if err := f.Validate(isNew); err != nil {
			return err
		}
	}
	return nil
}

func (s *FeaturesUsecase) enforce(ctx context.Context, tx repo.TX) error {
	curUser := ctx.Value(data.CtxUserName).(string)
	ires := repo.NewResource4Sv1("features", "", "", "")
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

// CreateFeatures is
func (s *FeaturesUsecase) CreateFeatures(ctx context.Context, features []*Feature) error {
	if err := s.validate(true, features); err != nil {
		return err
	}
	_f, e := ToDBFeatures(features)
	if e != nil {
		return e
	}
	err := s.txm.RunInTX(func(tx repo.TX) error {
		if err := s.enforce(ctx, tx); err != nil {
			return err
		}
		if err := s.ftrepo.CreateFeatures(ctx, tx, _f); err != nil {
			return err
		}
		return nil
	})
	return err
	//return s.repo.CreateFeatures(ctx, _f)
}

// UpdateFeatures is
func (s *FeaturesUsecase) UpdateFeatures(ctx context.Context, features []*Feature) error {
	if err := s.validate(false, features); err != nil {
		return err
	}
	_f, e := ToDBFeatures(features)
	if e != nil {
		return e
	}
	err := s.txm.RunInTX(func(tx repo.TX) error {
		if err := s.enforce(ctx, tx); err != nil {
			return err
		}
		if err := s.ftrepo.UpdateFeatures(ctx, tx, _f); err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteFeatures is
func (s *FeaturesUsecase) DeleteFeatures(ctx context.Context, ids []uint32) error {
	if len(ids) == 0 {
		return fmt.Errorf("EmptyIds")
	}
	return s.txm.RunInTX(func(tx repo.TX) error {
		if err := s.enforce(ctx, tx); err != nil {
			return err
		}
		for _, r := range s.required {
			c, err := r.inst.CountRequire(ctx, nil, repo.RequireFeature, ids)
			if err != nil {
				return err
			}
			if c > 0 {
				return fmt.Errorf("some %s requires", r.name)
			}
		}
		return s.ftrepo.DeleteFeatures(ctx, tx, ids)
	})
}

// GetFeatures is
func (s *FeaturesUsecase) GetFeatures(ctx context.Context, id uint32) (*Feature, error) {
	if id <= 0 {
		return nil, fmt.Errorf("EmptyId")
	}
	_f, e := s.ftrepo.GetFeatures(ctx, id)
	if e != nil {
		return nil, e
	}
	return ToBizFeature(_f)
}

// ListFeatures is
func (s *FeaturesUsecase) ListFeatures(ctx context.Context,
	filter *ListFeaturesFilter) ([]*Feature, error) {

	if filter != nil {
		if err := filter.Validate(); err != nil {
			return nil, err
		}
	}
	_f, e := s.ftrepo.ListFeatures(ctx, nil, ToDBFeaturesFilter(filter))
	if e != nil {
		return nil, e
	}
	return ToBizFeatures(_f)
}
