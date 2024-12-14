package biz

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
)

type EnvsRepo interface {
	CreateEnvs(ctx context.Context, envs []Env) error
	UpdateEnvs(ctx context.Context, envs []Env) error
	DeleteEnvs(ctx context.Context, ids []int64) error
	GetEnvs(ctx context.Context, id int64) (*Env, error)
	ListEnvs(ctx context.Context, filter *ListEnvsFilter) ([]Env, error)
}

type EnvsUsecase struct {
	repo EnvsRepo
	log  *log.Helper
}

func NewEnvsUsecase(repo EnvsRepo, logger log.Logger) *EnvsUsecase {
	return &EnvsUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (s *EnvsUsecase) validate(isNew bool, envs []Env) error {
	for _, e := range envs {
		if err := e.Validate(isNew); err != nil {
			return err
		}
	}
	return nil
}

// CreateEnvs is
func (s *EnvsUsecase) CreateEnvs(ctx context.Context, envs []Env) error {
	if err := s.validate(true, envs); err != nil {
		return err
	}
	return s.repo.CreateEnvs(ctx, envs)
}

// UpdateEnvs is
func (s *EnvsUsecase) UpdateEnvs(ctx context.Context, envs []Env) error {
	if err := s.validate(false, envs); err != nil {
		return err
	}
	return s.repo.UpdateEnvs(ctx, envs)
}

// DeleteEnvs is
func (s *EnvsUsecase) DeleteEnvs(ctx context.Context, ids []int64) error {
	if len(ids) == 0 {
		return fmt.Errorf("EmptyIds")
	}
	return s.repo.DeleteEnvs(ctx, ids)
}

// GetEnvs is
func (s *EnvsUsecase) GetEnvs(ctx context.Context, id int64) (*Env, error) {
	if id <= 0 {
		return nil, fmt.Errorf("InvalidId")
	}
	return s.repo.GetEnvs(ctx, id)
}

// ListEnvs is
func (s *EnvsUsecase) ListEnvs(ctx context.Context, filter *ListEnvsFilter) ([]Env, error) {

	if filter != nil {
		if err := filter.Validate(); err != nil {
			return nil, err
		}
	}
	return s.repo.ListEnvs(ctx, filter)
}
