package biz

import (
	"context"
	"fmt"

	"appix/internal/data/repo"

	"github.com/go-kratos/kratos/v2/log"
)

type TeamsUsecase struct {
	teamRepo      repo.TeamsRepo
	hostgroupRepo repo.HostgroupsRepo // hostgroup need team as foreign key
	txm           repo.TxManager
	log           *log.Helper
}

func NewTeamsUsecase(
	teamrepo repo.TeamsRepo,
	hgrepo repo.HostgroupsRepo,
	logger log.Logger,
	txm repo.TxManager) *TeamsUsecase {

	return &TeamsUsecase{
		teamRepo:      teamrepo,
		log:           log.NewHelper(logger),
		txm:           txm,
		hostgroupRepo: hgrepo,
	}
}

func (s *TeamsUsecase) validate(isNew bool, teams []*Team) error {
	for _, t := range teams {
		if err := t.Validate(isNew); err != nil {
			return err
		}
	}
	return nil
}

// CreateTeams is
func (s *TeamsUsecase) CreateTeams(ctx context.Context, teams []*Team) error {
	if err := s.validate(true, teams); err != nil {
		return err
	}

	_teams, e := ToDBTeams(teams)
	if e != nil {
		return e
	}

	return s.teamRepo.CreateTeams(ctx, _teams)
}

// UpdateTeams is
func (s *TeamsUsecase) UpdateTeams(ctx context.Context, teams []*Team) error {
	if err := s.validate(false, teams); err != nil {
		return err
	}
	_teams, e := ToDBTeams(teams)
	if e != nil {
		return e
	}
	return s.teamRepo.UpdateTeams(ctx, _teams)
}

// DeleteTeams is
func (s *TeamsUsecase) DeleteTeams(ctx context.Context, ids []uint32) error {
	if len(ids) == 0 {
		return fmt.Errorf("EmptyIds")
	}

	return s.txm.RunInTX(
		func(tx repo.TX) error {
			// check hostgroups count use team ids
			return nil
		})

	// return s.teamRepo.DeleteTeams(ctx, ids)
}

// GetTeams is
func (s *TeamsUsecase) GetTeams(ctx context.Context, id uint32) (*Team, error) {
	if id <= 0 {
		return nil, fmt.Errorf("EmptyId")
	}
	dbt, e := s.teamRepo.GetTeams(ctx, id)
	if e != nil {
		return nil, e
	}
	return ToBizTeam(dbt)
}

// ListTeams is
func (s *TeamsUsecase) ListTeams(ctx context.Context,
	filter *ListTeamsFilter) ([]*Team, error) {
	if filter != nil {
		if err := filter.Validate(); err != nil {
			return nil, err
		}
	}
	teams, e := s.teamRepo.ListTeams(ctx, nil, ToDBTeamsFilter(filter))

	if e != nil {
		return nil, e
	}
	return ToBizTeams(teams)
}
