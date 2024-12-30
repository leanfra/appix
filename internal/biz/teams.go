package biz

import (
	"context"
	"fmt"

	"appix/internal/data/repo"

	"github.com/go-kratos/kratos/v2/log"
)

type TeamsUsecase struct {
	teamRepo repo.TeamsRepo
	txm      repo.TxManager
	log      *log.Helper
	required []requiredBy
}

func NewTeamsUsecase(
	teamrepo repo.TeamsRepo,
	hgrepo repo.HostgroupsRepo,
	htrepo repo.HostgroupTeamsRepo,
	apprepo repo.ApplicationsRepo,
	logger log.Logger,
	txm repo.TxManager) *TeamsUsecase {

	return &TeamsUsecase{
		teamRepo: teamrepo,
		log:      log.NewHelper(logger),
		txm:      txm,
		required: []requiredBy{
			{inst: hgrepo, name: "hostgroup"},
			{inst: apprepo, name: "app"},
			{inst: htrepo, name: "hostgroup_team"},
		},
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
func (s *TeamsUsecase) DeleteTeams(ctx context.Context,
	ids []uint32) error {
	if len(ids) == 0 {
		return fmt.Errorf("EmptyIds")
	}

	return s.txm.RunInTX(
		func(tx repo.TX) error {

			for _, r := range s.required {
				c, err := r.inst.CountRequire(ctx, tx, repo.RequireTeam, ids)
				if err != nil {
					return err
				}
				if c > 0 {
					return fmt.Errorf("some %s requires", r.name)
				}
			}
			if e := s.teamRepo.DeleteTeams(ctx, tx, ids); e != nil {
				return e
			}
			return nil
		})
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
