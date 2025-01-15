package biz

import (
	"appix/internal/data"
	"appix/internal/data/repo"
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
)

type EnvsUsecase struct {
	envrepo   repo.EnvsRepo
	authzrepo repo.AuthzRepo
	log       *log.Helper
	txm       repo.TxManager
	required  []requiredBy
}

func NewEnvsUsecase(
	repo repo.EnvsRepo,
	authzrepo repo.AuthzRepo,
	hgrepo repo.HostgroupsRepo,
	logger log.Logger,
	txm repo.TxManager) *EnvsUsecase {
	return &EnvsUsecase{
		envrepo:   repo,
		authzrepo: authzrepo,
		log:       log.NewHelper(logger),
		txm:       txm,
		required: []requiredBy{
			{inst: hgrepo, name: "hostgroup"},
		},
	}
}

func (s *EnvsUsecase) validate(isNew bool, envs []*Env) error {
	for _, e := range envs {
		if err := e.Validate(isNew); err != nil {
			return err
		}
	}
	return nil
}

func (s *EnvsUsecase) enforce(ctx context.Context, tx repo.TX) error {
	curuser := ctx.Value(data.UserName).(string)
	ires := repo.NewResource4Sv1("env", "", "", "")
	can, err := s.authzrepo.Enforce(ctx, tx, &repo.AuthenRequest{
		Sub:      curuser,
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

// CreateEnvs is
func (s *EnvsUsecase) CreateEnvs(ctx context.Context, envs []*Env) error {
	if err := s.validate(true, envs); err != nil {
		return err
	}
	_envs, err := ToDBEnvs(envs)
	if err != nil {
		return err
	}
	err = s.txm.RunInTX(func(tx repo.TX) error {
		if err := s.enforce(ctx, tx); err != nil {
			return err
		}
		if err := s.envrepo.CreateEnvs(ctx, tx, _envs); err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateEnvs is
func (s *EnvsUsecase) UpdateEnvs(ctx context.Context, envs []*Env) error {
	if err := s.validate(false, envs); err != nil {
		return err
	}
	_envs, e := ToDBEnvs(envs)
	if e != nil {
		return e
	}

	//return s.envrepo.UpdateEnvs(ctx, _envs)
	err := s.txm.RunInTX(func(tx repo.TX) error {
		if err := s.enforce(ctx, tx); err != nil {
			return err
		}
		if err := s.envrepo.UpdateEnvs(ctx, tx, _envs); err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteEnvs is
func (s *EnvsUsecase) DeleteEnvs(ctx context.Context, ids []uint32) error {
	if len(ids) == 0 {
		return fmt.Errorf("EmptyIds")
	}
	return s.txm.RunInTX(func(tx repo.TX) error {
		if err := s.enforce(ctx, tx); err != nil {
			return err
		}
		for _, r := range s.required {
			c, err := r.inst.CountRequire(ctx, nil, repo.RequireEnv, ids)
			if err != nil {
				return err
			}
			if c > 0 {
				return fmt.Errorf("some %s requires", r.name)
			}
		}
		return s.envrepo.DeleteEnvs(ctx, tx, ids)
	})
}

// GetEnvs is
func (s *EnvsUsecase) GetEnvs(ctx context.Context, id uint32) (*Env, error) {
	if id <= 0 {
		return nil, fmt.Errorf("InvalidId")
	}
	_envs, err := s.envrepo.GetEnvs(ctx, id)
	if err != nil {
		return nil, err
	}
	return ToBizEnv(_envs)
}

// ListEnvs is
func (s *EnvsUsecase) ListEnvs(ctx context.Context, filter *ListEnvsFilter) ([]*Env, error) {

	if filter != nil {
		if err := filter.Validate(); err != nil {
			return nil, err
		}
	}

	_envs, err := s.envrepo.ListEnvs(ctx, nil, ToDBEnvsFilter(filter))
	if err != nil {
		return nil, err
	}
	return ToBizEnvs(_envs)
}
