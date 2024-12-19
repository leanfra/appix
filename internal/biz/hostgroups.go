package biz

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
)

type HostgroupsRepo interface {
	CreateHostgroups(ctx context.Context, hgs []*Hostgroup) error
	UpdateHostgroups(ctx context.Context, hgs []*Hostgroup) error
	DeleteHostgroups(ctx context.Context, ids []uint32) error
	GetHostgroups(ctx context.Context, id uint32) (*Hostgroup, error)
	ListHostgroups(ctx context.Context, filter *ListHostgroupsFilter) ([]*Hostgroup, error)
}

type HostgroupsUsecase struct {
	repo HostgroupsRepo
	log  *log.Helper
}

func NewHostgroupsUsecase(repo HostgroupsRepo, logger log.Logger) *HostgroupsUsecase {
	return &HostgroupsUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (s *HostgroupsUsecase) validate(isNew bool, hgs []*Hostgroup) error {
	for _, hg := range hgs {
		if err := hg.Validate(isNew); err != nil {
			return err
		}
	}
	return nil
}

// CreateHostgroups is
func (s *HostgroupsUsecase) CreateHostgroups(ctx context.Context, hgs []*Hostgroup) error {
	if err := s.validate(true, hgs); err != nil {
		return err
	}
	return s.repo.CreateHostgroups(ctx, hgs)
}

// UpdateHostgroups is
func (s *HostgroupsUsecase) UpdateHostgroups(ctx context.Context, hgs []*Hostgroup) error {
	if err := s.validate(false, hgs); err != nil {
		return err
	}
	return s.repo.UpdateHostgroups(ctx, hgs)
}

// DeleteHostgroups is
func (s *HostgroupsUsecase) DeleteHostgroups(ctx context.Context, ids []uint32) error {
	if len(ids) == 0 {
		return fmt.Errorf("EmptyIds")
	}
	return s.repo.DeleteHostgroups(ctx, ids)
}

// GetHostgroups is
func (s *HostgroupsUsecase) GetHostgroups(ctx context.Context, id uint32) (*Hostgroup, error) {
	if id <= 0 {
		return nil, fmt.Errorf("InvalidId")
	}
	return s.repo.GetHostgroups(ctx, id)
}

// ListHostgroups is
func (s *HostgroupsUsecase) ListHostgroups(
	ctx context.Context, filter *ListHostgroupsFilter) ([]*Hostgroup, error) {

	if filter != nil {
		if err := filter.Validate(); err != nil {
			return nil, err
		}
	}
	return s.repo.ListHostgroups(ctx, filter)
}
